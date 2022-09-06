package services

import (
	chk "db_course/business/checker"
	ents "db_course/business/entities"
	repos "db_course/dataAccess/repositories"
	"db_course/my_errors"
	"fmt"
	"strconv"
)

type IFoundationService interface {
	Add(FPars chk.FoundationMainParams) error
	Update(id_ string, FPars chk.FoundationMainParams) error
	Delete(id_ string) error
	GetAll() ([]ents.Foundation, error)
	GetById(id_ string) (ents.Foundation, error)
	GetBylogin(id_ string) (ents.Foundation, error)
	ExistsById(id uint64) bool
	ExistsBylogin(name string) bool
	Donate(id_ string, DP chk.FoundationDonateParams) error
	AcceptDonate(id_ string, sum float64) error
}

type FoundationService struct {
	FR repos.FoundationRepository
}

func NewFoundationService(frepo repos.FoundationRepository) FoundationService {
	return FoundationService{FR: frepo}
}

func (FS *FoundationService) ExistsById(id uint64) bool {
	_, result := FS.FR.SelectById(id)
	return result
}
func (FS *FoundationService) ExistsByLogin(name string) bool {
	_, result := FS.FR.SelectByLogin(name)
	return result
}
func (FS *FoundationService) ExistsByName(name string) bool {
	_, result := FS.FR.SelectByName(name)
	return result
}

func (FS *FoundationService) ExistsByCountry(country string) bool {
	_, result := FS.FR.SelectByCountry(country)
	return result
}

func (FS *FoundationService) Add(FPars chk.FoundationMainParams) error {
	if FS.ExistsByLogin(FPars.Login) {
		return fmt.Errorf(my_errors.ErrAlreadyExists)
	} else {
		if chk.CheckCountry(FPars.Country) != nil {
			return fmt.Errorf(my_errors.ErrCountry)
		} else {
			var F ents.Foundation = ents.NewFoundation()
			var err_ error = FPars.CheckParams()
			if err_ == nil {
				F.SetLogin(FPars.Login)
				F.SetPassword(FPars.Password)
				F.SetCountry(FPars.Country)
				F.SetName(FPars.Name)
				FS.FR.Insert(F)
			} else {
				return err_
			}
		}
	}
	return nil
}

func (FS *FoundationService) Update(id_ string, FPars chk.FoundationMainParams) error {
	if FS.ExistsByLogin(FPars.Login) {
		return fmt.Errorf(my_errors.ErrAlreadyExists)
	} else {
		var errGet error
		var F ents.Foundation
		F, errGet = FS.GetById(id_)
		if errGet != nil {
			return errGet
		} else {
			var err_ error = FPars.CheckParams()
			if err_ == nil {
				F.SetLogin(FPars.Login)
				F.SetPassword(FPars.Password)
				F.SetCountry(FPars.Country)
				F.SetName(FPars.Name)
				FS.FR.Update(F)
			} else {
				return err_
			}
		}
	}
	return nil
}

func (FS *FoundationService) Delete(id_ string) error {
	var errGet error
	var F ents.Foundation
	F, errGet = FS.GetById(id_)
	if errGet != nil {
		return errGet
	} else {
		return FS.FR.Delete(F)
	}
}
func (FS *FoundationService) GetAll() ([]ents.Foundation, error) {
	Foundations, err := FS.FR.Select()
	if !err {
		return nil, fmt.Errorf("не удалось получить список всех фондов")
	} else {
		return Foundations, nil
	}
}

func (FS *FoundationService) GetById(id_ string) (ents.Foundation, error) {
	sid, err := strconv.Atoi(id_)
	id := uint64(sid)
	var F ents.Foundation
	if err != nil {
		return F, fmt.Errorf("некорректный id")
	} else {
		if !FS.ExistsById(id) {
			return F, fmt.Errorf(my_errors.ErrNotExists)
		} else {
			var err_ bool
			F, err_ = FS.FR.SelectById(id)
			if !err_ {
				return F, fmt.Errorf("не удалось получить фонд по id")
			}
		}
	}
	return F, nil
}
func (FS *FoundationService) GetByLogin(name_ string) (ents.Foundation, error) {
	var F ents.Foundation
	if !FS.ExistsByLogin(name_) {
		return F, fmt.Errorf(my_errors.ErrNotExists)
	} else {
		var err_ bool
		F, err_ = FS.FR.SelectByLogin(name_)
		if !err_ {
			return F, fmt.Errorf("не удалось получить фонд по логину")
		}
	}
	return F, nil
}
func (FS *FoundationService) GetByName(name_ string) (ents.Foundation, error) {
	var F ents.Foundation
	if !FS.ExistsByName(name_) {
		return F, fmt.Errorf(my_errors.ErrNotExists)
	} else {
		var err_ bool
		F, err_ = FS.FR.SelectByName(name_)
		if !err_ {
			return F, fmt.Errorf("не удалось получить фонд по названию")
		}
	}
	return F, nil
}

func (FS *FoundationService) GetByCountry(country string) ([]ents.Foundation, error) {
	var F []ents.Foundation
	if chk.CheckCountry(country) != nil {
		return F, fmt.Errorf(my_errors.ErrCountry)
	} else {
		var err_ bool
		F, err_ = FS.FR.SelectByCountry(country)
		if !err_ {
			return F, fmt.Errorf("не удалось получить фонд по стране")
		}
	}
	return F, nil
}

func (FS *FoundationService) Donate(F *ents.Foundation, DP chk.FoundationDonateParams) error {

	F.Fund_balance -= DP.Sum_of_money
	F.Outcome_history += DP.Sum_of_money
	if DP.IsClosedFoundrising {
		F.ClosedFoundrisingAmount += 1
		F.CurFoudrisingAmount -= 1
	}
	err := FS.FR.Update(*F)
	return err
}

func (FS *FoundationService) AcceptDonate(id_ string, sum float64) error {

	var F ents.Foundation
	var err error
	F, err = FS.GetById(id_)
	if err != nil {
		return fmt.Errorf("фонду не удалось получить средства")
	} else {
		F.Fund_balance += sum
		F.Income_history += sum
	}
	FS.FR.Update(F)
	return nil
}
func (US *FoundationService) ReplenishBalance(U *ents.Foundation, sum float64) error {
	var err error
	if err != nil {
		return fmt.Errorf("фонду не удалось пополнить баланс")
	} else {
		U.Fund_balance += sum
	}
	US.FR.Update(*U)
	return nil
}
