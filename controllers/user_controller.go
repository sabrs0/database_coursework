package controllers

import (
	"db_course/business/checker"
	ents "db_course/business/entities"
	servs "db_course/business/services"
	"fmt"
	"strconv"
)

type UserController struct {
	US  servs.UserService
	FS  servs.FoundationService
	FgS servs.FoundrisingService
	TS  servs.TransactionService
}

func (UC *UserController) GetAll() ([]ents.User, error) {
	return UC.US.GetAll()
}
func (UC *UserController) GetByID(id_ string) (ents.User, error) {
	return UC.US.GetById(id_)
}
func (UC *UserController) GetByLogin(login string) (ents.User, error) {
	return UC.US.GetByLogin(login)
}

func (UC *UserController) Add(login string, password string) error {
	UP := checker.NewUserMainParams(login, password)
	return UC.US.Add(UP)

}

func (UC *UserController) Delete(id string) error {
	err := UC.US.Delete(id)
	if err == nil {
		transactions1, err := UC.TS.GetFromId(ents.FROM_USER, id, UC.FS, UC.US)
		if err == nil {
			for i := range transactions1 {
				err = UC.TS.Delete(strconv.FormatUint(transactions1[i].Id, 10))
				if err != nil {
					return err
				}
			}
		}
	}
	return err

}
func (UC *UserController) Update(id string, login string, password string) error {
	var User ents.User
	User, _ = UC.GetByID(id)
	if login == "" {
		login = User.Login
	}
	if password == "" {
		password = User.Password
	}
	UP := checker.NewUserMainParams(login, password)
	return UC.US.Update(id, UP)

}
func (UC *UserController) DonateToFoundation(sum string, comm string, to_id string, U *ents.User) error {
	var err error
	var reqSum float64
	reqSum, err = strconv.ParseFloat(sum, 64)
	var UDP checker.UserDonateParams
	if err == nil {
		err = checkMoneyFormat(sum)
		if err == nil {
			if U.Balance < reqSum {
				return fmt.Errorf("недостаточно средств ")
			} else {
				UDP = checker.NewUserDonateParams(reqSum, false)
				err = UC.US.Donate(U, UDP)
				if err == nil {
					err = UC.FS.AcceptDonate(to_id, reqSum)
					if err == nil {
						sid, err := strconv.Atoi(to_id)
						found_id := uint64(sid)
						if err == nil {
							TP := checker.NewTransactionMainParams(ents.FROM_USER, U.Id, ents.TO_FOUNDATION, reqSum,
								comm, found_id)
							err = UC.TS.Add(TP)
							return err
						}
					}
				}
			}
		}
	}
	return err
}
func isNewPhilantrop(TS servs.TransactionService, FgS servs.FoundrisingService, foundrisingId string, userId uint64) (bool, error) {
	IDs, err := TS.GetFoundrisingPhilantropIds(foundrisingId, FgS)
	var isNew bool
	if err == nil {
		for _, id := range IDs {
			if id == userId {
				isNew = true
			}
		}
	}
	return isNew, err
}
func (UC *UserController) DonateToFoundrising(sum string, comm string, to_id string, U *ents.User) error {
	var err error
	var reqSum float64
	reqSum, err = strconv.ParseFloat(sum, 64)
	var UDP checker.UserDonateParams
	if err == nil {
		err = checkMoneyFormat(sum)
		if err == nil {
			if U.Balance < reqSum {
				return fmt.Errorf("недостаточно средств ")
			} else {
				foundrising, err := UC.FgS.GetById(to_id)
				if err == nil {
					if !foundrising.Closing_date.Valid {
						UDP = checker.NewUserDonateParams(reqSum, false)
						err = UC.US.Donate(U, UDP)
						if err == nil {
							isNewPh, err := isNewPhilantrop(UC.TS, UC.FgS, to_id, U.Id)
							if err == nil {
								var remainder float64
								remainder, err = UC.FgS.AcceptDonate(to_id, reqSum, isNewPh)
								if err == nil {
									sid, err := strconv.Atoi(to_id)
									foundrising_id := uint64(sid)
									if err == nil {
										TP := checker.NewTransactionMainParams(ents.FROM_USER, U.Id, ents.TO_FOUNDRISING, reqSum,
											comm, foundrising_id)
										err = UC.TS.Add(TP)
										if err == nil {
											if remainder > 0.0 {
												foundrising, _ := UC.FgS.GetById(to_id)
												found_id := foundrising.Found_id
												TP := checker.NewTransactionMainParams(ents.FROM_USER, U.Id, ents.TO_FOUNDATION, remainder,
													"returning the remain", found_id)
												err = UC.TS.Add(TP)
												return err
											} else if remainder <= 1e-9 {
												U.ClosedFingAmount += 1
												UC.US.UR.Update(*U)
											}
											return nil
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return err
}

func (UC *UserController) ReplenishBalance(sum string, U *ents.User) error {
	var err error
	var reqSum float64
	reqSum, err = strconv.ParseFloat(sum, 64)
	if err == nil {
		err = checkMoneyFormat(sum)
		if err == nil {
			if reqSum > 50000.00 {
				return fmt.Errorf("введенная сумма превышается 50 000")
			} else {
				return UC.US.ReplenishBalance(U, reqSum)
			}
		}
	}
	return err
}
