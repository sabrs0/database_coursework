package controllers

import (
	"db_course/business/checker"
	ents "db_course/business/entities"
	servs "db_course/business/services"
	"db_course/my_errors"
	"fmt"
	"strconv"
	"time"
)

type FoundrisingController struct {
	FS   servs.FoundrisingService
	FndS servs.FoundationService
}

func checkMoneyFormat(money string) error {
	for i, c := range money {
		if c == '.' {
			if i < len(money)-1 && len(money)-1-i < 3 && i > 0 {
				for j := i + 1; j < len(money); j++ {
					if money[j] < '0' || money[j] > '9' {
						return fmt.Errorf(my_errors.ErrMoney)
					}
				}
			} else {
				return fmt.Errorf(my_errors.ErrMoney)
			}
		}
	}
	return nil
}

func (UC *FoundrisingController) GetAll() ([]ents.Foundrising, error) {
	return UC.FS.GetAll()
}
func (UC *FoundrisingController) GetByID(id_ string) (ents.Foundrising, error) {
	return UC.FS.GetById(id_)
}
func (UC *FoundrisingController) GetByCreDate(date string) ([]ents.Foundrising, error) {
	return UC.FS.GetByCreateDate(date)
}
func (UC *FoundrisingController) GetByCloDate(date string) ([]ents.Foundrising, error) {
	return UC.FS.GetByCloseDate(date)
}
func (UC *FoundrisingController) GetByFoundId(id string) ([]ents.Foundrising, error) {
	return UC.FS.GetByFoundId(id)
}
func (UC *FoundrisingController) GetByFoundIdActive(id string) ([]ents.Foundrising, error) {
	return UC.FS.GetByFoundIdActive(id)
}

func (UC *FoundrisingController) Add(found_id_ string, descr string, reqSum_ string) error {
	sid, err := strconv.Atoi(found_id_)
	found_id := uint64(sid)
	var FP checker.FoundrisingMainParams
	if err == nil {
		reqSum, err := strconv.ParseFloat(reqSum_, 64)
		if err == nil {
			err = checkMoneyFormat(reqSum_)
			if err == nil {
				create_date := time.Now().Format(ents.DateFormat)
				FP = checker.NewFoundrisingMainParams(found_id, descr, reqSum, create_date)
				if !UC.FndS.ExistsById(found_id) {
					return fmt.Errorf("фонда с таким ID не существует")
				}
			} else {
				return fmt.Errorf(my_errors.ErrMoney)
			}
		} else {
			return fmt.Errorf(my_errors.ErrMoney)
		}
	} else {
		return fmt.Errorf("incorrect found_id ")
	}
	return UC.FS.Add(FP)

}

func (UC *FoundrisingController) Delete(id string) error {
	return UC.FS.Delete(id)

}
func (UC *FoundrisingController) Update(id string, descr string, reqSum_ string) error {
	var Foundrising ents.Foundrising
	Foundrising, _ = UC.GetByID(id)
	var FP checker.FoundrisingMainParams
	if descr == "" {
		descr = Foundrising.Description
	}
	if reqSum_ == "" {
		reqSum_ = strconv.FormatFloat(Foundrising.Required_sum, 'f', 2, 64)

	}
	reqSumfloat, err := strconv.ParseFloat(reqSum_, 64)
	if err == nil {
		err = checkMoneyFormat(reqSum_)
		if err == nil {
			if reqSumfloat < Foundrising.Required_sum {
				return fmt.Errorf("новая сумма меньше той, что была прежде")
			} else {
				FP = checker.NewFoundrisingMainParams(Foundrising.Found_id, descr, reqSumfloat, Foundrising.Creation_date)
			}
		} else {
			return fmt.Errorf(my_errors.ErrMoney)
		}
	} else {
		return fmt.Errorf(my_errors.ErrMoney)
	}
	return UC.FS.Update(id, FP)

}
