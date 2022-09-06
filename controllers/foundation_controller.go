package controllers

import (
	"db_course/business/checker"
	ents "db_course/business/entities"
	servs "db_course/business/services"
	repos "db_course/dataAccess/repositories"
	"fmt"
	"strconv"
)

type FoundationController struct {
	FS  servs.FoundationService
	TS  servs.TransactionService
	FgS servs.FoundrisingService
}

func (UC *FoundationController) GetAll() ([]ents.Foundation, error) {
	return UC.FS.GetAll()
}
func (UC *FoundationController) GetByID(id_ string) (ents.Foundation, error) {
	return UC.FS.GetById(id_)
}
func (UC *FoundationController) GetByName(name string) (ents.Foundation, error) {
	return UC.FS.GetByName(name)
}
func (UC *FoundationController) GetByLogin(login string) (ents.Foundation, error) {
	return UC.FS.GetByLogin(login)
}

func (UC *FoundationController) GetByCountry(country string) ([]ents.Foundation, error) {
	return UC.FS.GetByCountry(country)
}

func (UC *FoundationController) Add(login string, password string, name string, country string) error {
	UP := checker.NewFoundationMainParams(login, password, name, country)
	foundations_by_country, err := UC.FS.GetByCountry(country)
	if err == nil {
		for _, f := range foundations_by_country {
			if f.Name == name {
				return fmt.Errorf("в данной стране уже есть фонд с таким названием")
			}
		}
	}
	return UC.FS.Add(UP)

}

func (UC *FoundationController) Delete(id string) error {
	err := UC.FS.Delete(id)
	if err == nil {
		UR := repos.NewUserRepository(UC.FS.FR.DB)
		US := servs.NewUserService(*UR)
		transactions1, err := UC.TS.GetFromId(ents.FROM_FOUNDATION, id, UC.FS, US)
		if err == nil {
			for i := range transactions1 {
				err = UC.TS.Delete(strconv.FormatUint(transactions1[i].Id, 10))
				if err != nil {
					return err
				}
			}
			transactions2, err := UC.TS.GetToId(ents.TO_FOUNDATION, id, UC.FS, UC.FgS)
			if err == nil {
				for i := range transactions2 {
					err = UC.TS.Delete(strconv.FormatUint(transactions2[i].Id, 10))
					if err != nil {
						return err
					}
				}
			}
			foundrisings, err := UC.FgS.GetByFoundId(id)
			if err == nil {
				for i := range foundrisings {
					err = UC.FgS.Delete(strconv.FormatUint(foundrisings[i].Id, 10))
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return err

}
func (UC *FoundationController) Update(id string, login string, password string, name string, country string) error {
	var foundation ents.Foundation
	foundation, _ = UC.GetByID(id)
	if login == "" {
		login = foundation.Login
	}
	if password == "" {
		password = foundation.Password
	}
	if name == "" {
		name = foundation.Name
	}
	if country == "" {
		country = foundation.Country
	}
	UP := checker.NewFoundationMainParams(login, password, name, country)
	return UC.FS.Update(id, UP)

}
func (UC *FoundationController) AcceptDonate(id string, sum float64) error {
	return UC.FS.AcceptDonate(id, sum)
}
func (UC *FoundationController) DonateToFoundrising(sum string, comm string, to_id string, U *ents.Foundation) error {
	var err error
	var reqSum float64
	reqSum, err = strconv.ParseFloat(sum, 64)
	var UDP checker.FoundationDonateParams
	if err == nil {
		err = checkMoneyFormat(sum)
		if err == nil {
			if U.Fund_balance < reqSum {
				return fmt.Errorf("недостаточно средств ")
			} else {
				foundrising, err := UC.FgS.GetById(to_id)
				if err == nil {
					if !foundrising.Closing_date.Valid {
						UDP = checker.NewFoundationDonateParams(reqSum, false)
						err = UC.FS.Donate(U, UDP)
						if err == nil {
							if err == nil {
								var remainder float64
								remainder, err = UC.FgS.AcceptDonate(to_id, reqSum, false)
								if err == nil {
									sid, err := strconv.Atoi(to_id)
									foundrising_id := uint64(sid)
									if err == nil {
										TP := checker.NewTransactionMainParams(ents.FROM_FOUNDATION, U.Id, ents.TO_FOUNDRISING, reqSum,
											comm, foundrising_id)
										err = UC.TS.Add(TP)
										if err == nil {
											if remainder > 0.0 {
												foundrising, _ := UC.FgS.GetById(to_id)
												found_id := foundrising.Found_id
												TP := checker.NewTransactionMainParams(ents.FROM_FOUNDATION, U.Id, ents.TO_FOUNDATION, remainder,
													"returning the remain", found_id)
												err = UC.TS.Add(TP)
												return err
											} else if remainder <= 1e-9 {
												U.ClosedFoundrisingAmount += 1
												if U.CurFoudrisingAmount > 0 {
													U.CurFoudrisingAmount -= 1
												}
												UC.FS.FR.Update(*U)
											}
											return nil
										}
									}
								}
							}
						}
					} else {
						return fmt.Errorf("данный сбор уже закрыт")
					}
				}
			}
		}
	}
	return err
}

func (UC *FoundationController) ReplenishBalance(sum string, U *ents.Foundation) error {
	var err error
	var reqSum float64
	reqSum, err = strconv.ParseFloat(sum, 64)
	if err == nil {
		err = checkMoneyFormat(sum)
		if err == nil {
			if reqSum > 50000.00 {
				return fmt.Errorf("введенная сумма превышается 50 000")
			} else {
				return UC.FS.ReplenishBalance(U, reqSum)
			}
		}
	}
	return err
}
