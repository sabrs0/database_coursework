package services

import (
	chk "db_course/business/checker"
	ents "db_course/business/entities"
	repos "db_course/dataAccess/repositories"
	"db_course/my_errors"
	"fmt"
	"math"
	"strconv"
	"time"
)

type IFoundrisingService interface {
	Add(FPars chk.FoundrisingMainParams) error
	Update(id_ string, FPars chk.FoundrisingMainParams) error
	Delete(id_ string) error
	GetAll() ([]ents.Foundrising, error)
	GetById(id_ string) (ents.Foundrising, error)
	GetByCreateDate(date string) ([]ents.Foundrising, error)
	GetByCloseDate(date string) ([]ents.Foundrising, error)
	ExistsById(id uint64) bool
	AcceptDonate(id_ string, sum float64) (float64, error)
}

type FoundrisingService struct {
	FR repos.FoundrisingRepository
}

func NewFoundrisingService(frepo repos.FoundrisingRepository) FoundrisingService {
	return FoundrisingService{FR: frepo}
}

func (FS *FoundrisingService) ExistsById(id uint64) bool {
	_, result := FS.FR.SelectById(id)
	return result
}

func (FS *FoundrisingService) Add(FPars chk.FoundrisingMainParams) error {

	var U ents.Foundrising = ents.NewFoundrising()
	U.SetReqSum(FPars.ReqSum)
	U.SetCreateDate(FPars.CreateDate)
	U.SetDescription(FPars.Descr)
	U.SetFoundId(FPars.Found_id)
	err := FS.FR.Insert(U)

	return err
}

func (FS *FoundrisingService) Update(id_ string, FPars chk.FoundrisingMainParams) error {

	var errGet error
	var U ents.Foundrising
	U, errGet = FS.GetById(id_)
	if errGet != nil {
		return errGet
	} else {
		U.SetDescription(FPars.Descr)
		U.SetReqSum(FPars.ReqSum)
		FS.FR.Update(U)
	}
	return nil
}

func (FS *FoundrisingService) Delete(id_ string) error {
	var errGet error
	var U ents.Foundrising
	U, errGet = FS.GetById(id_)
	if errGet != nil {
		return errGet
	} else {
		FS.FR.Delete(U)
	}
	return nil
}
func (FS *FoundrisingService) GetAll() ([]ents.Foundrising, error) {
	Foundrisings, err := FS.FR.Select()
	fmt.Println(Foundrisings)

	if !err {
		return nil, fmt.Errorf("не удалось получить список всех сборов")
	} else {
		return Foundrisings, nil
	}
}
func (FS *FoundrisingService) GetById(id_ string) (ents.Foundrising, error) {
	sid, err := strconv.Atoi(id_)
	id := uint64(sid)
	var U ents.Foundrising
	if err != nil {
		return U, fmt.Errorf("некорректный id")
	} else {
		if !FS.ExistsById(id) {
			return U, fmt.Errorf("несуществующий id")
		} else {
			var err_ bool
			U, err_ = FS.FR.SelectById(id)
			if !err_ {
				return U, fmt.Errorf("не удалось получить сбор по id")
			}
		}
	}
	return U, nil
}

func (FS *FoundrisingService) GetByFoundId(id_ string) ([]ents.Foundrising, error) {
	sid, err := strconv.Atoi(id_)
	id := uint64(sid)
	var U []ents.Foundrising
	if err != nil {
		return U, fmt.Errorf("некорректный id")
	} else {
		var err_ bool
		U, err_ = FS.FR.SelectByFoundId(id)
		if !err_ {
			return U, fmt.Errorf("не удалось получить сбор по id фонда")
		}
	}
	return U, nil
}

func (FS *FoundrisingService) GetByIdAndFoundId(id_ string, found_id_ string) (ents.Foundrising, error) {
	sid, err1 := strconv.Atoi(id_)
	id := uint64(sid)

	s_found_id, err := strconv.Atoi(found_id_)
	found_id := uint64(s_found_id)
	var U ents.Foundrising
	if err != nil || err1 != nil {
		return U, fmt.Errorf("некорректный id")
	} else {
		var err_ bool
		U, err_ = FS.FR.SelectByIdAndFoundId(id, found_id)
		if !err_ {
			return U, fmt.Errorf("не удалось получить фонд данного сбора")
		}
	}
	return U, nil
}

func (FS *FoundrisingService) GetByFoundIdActive(id_ string) ([]ents.Foundrising, error) {
	sid, err := strconv.Atoi(id_)
	id := uint64(sid)
	var U []ents.Foundrising
	if err != nil {
		return U, fmt.Errorf("некорректный id")
	} else {
		var err_ bool
		U, err_ = FS.FR.SelectByFoundIdActive(id)
		if !err_ {
			return U, fmt.Errorf(my_errors.ErrNotExistsFoundrising)
		}
	}
	return U, nil
}

func (FS *FoundrisingService) GetByCreateDate(date string) ([]ents.Foundrising, error) {
	var U []ents.Foundrising
	var err_ bool
	U, err_ = FS.FR.SelectByCreateDate(date)
	if !err_ {
		return U, fmt.Errorf("не удалось получить сбор по id по дате создания")
	}
	return U, nil
}

func (FS *FoundrisingService) GetByCloseDate(date string) ([]ents.Foundrising, error) {
	var U []ents.Foundrising
	var err_ bool
	U, err_ = FS.FR.SelectByCloseDate(date)
	if !err_ {
		return U, fmt.Errorf("не удалось получить сбор по дате закрытия")
	}
	return U, nil
}

func (FS *FoundrisingService) AcceptDonate(id_ string, sum float64, isNewPhil bool) (float64, error) {

	var remainder float64 = -1.0
	var F ents.Foundrising
	var err error
	F, err = FS.GetById(id_)
	if err != nil {
		return remainder, err
	} else {
		if isNewPhil {
			F.Philantrops_amount += 1
		}
		if F.Current_sum+sum > F.Required_sum {
			remainder = F.Current_sum + sum - F.Required_sum
			F.Closing_date.String = time.Now().Format(ents.DateFormat)
			F.Closing_date.Valid = true
			F.Current_sum = F.Required_sum
		} else if math.Abs(F.Current_sum+sum-F.Required_sum) <= 1e-9 {
			F.Closing_date.String = time.Now().Format(ents.DateFormat)
			F.Closing_date.Valid = true
			F.Current_sum = F.Required_sum
			remainder = 0.00
		} else {
			F.Current_sum += sum
		}
	}
	FS.FR.Update(F)
	return remainder, nil
}
