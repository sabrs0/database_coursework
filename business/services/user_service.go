package services

import (
	chk "db_course/business/checker"
	ents "db_course/business/entities"
	repos "db_course/dataAccess/repositories"
	"db_course/my_errors"
	"fmt"
	"strconv"
)

type IUserService interface {
	Add(UPars chk.UserMainParams) error
	Update(id_ string, UPars chk.UserMainParams) error
	Delete(id_ string) error
	GetAll() ([]ents.User, error)
	GetById(id_ string) (ents.User, error)
	GetByLogin(id_ string) (ents.User, error)
	ExistsById(id uint64) bool
	ExistsByLogin(name string) bool
	Donate(id_ string, DP chk.UserDonateParams) error
	ReplenishBalance(id_ string, sum float64) error
}

type UserService struct {
	UR repos.UserRepository
}

func NewUserService(urepo repos.UserRepository) UserService {
	return UserService{UR: urepo}
}

func (US *UserService) ExistsById(id uint64) bool {
	_, result := US.UR.SelectById(id)
	return result
}

func (US *UserService) ExistsByLogin(name string) bool {
	_, result := US.UR.SelectByLogin(name)
	return result
}

func (US *UserService) Add(UPars chk.UserMainParams) error {
	if US.ExistsByLogin(UPars.Login) {
		return fmt.Errorf(my_errors.ErrAlreadyExists)
	} else {
		var U ents.User = ents.NewUser()
		var err_ error = UPars.CheckParams()
		if err_ == nil {
			U.SetLogin(UPars.Login)
			U.SetPassword(UPars.Password)
			US.UR.Insert(U)
		} else {
			return err_
		}
	}
	return nil
}

func (US *UserService) Update(id_ string, UPars chk.UserMainParams) error {
	if US.ExistsByLogin(UPars.Login) {
		return fmt.Errorf(my_errors.ErrAlreadyExists)
	} else {
		var errGet error
		var U ents.User
		U, errGet = US.GetById(id_)
		if errGet != nil {
			return errGet
		} else {
			var err_ error = UPars.CheckParams()
			if err_ == nil {
				U.SetLogin(UPars.Login)
				U.SetPassword(UPars.Password)
				US.UR.Update(U)
			} else {
				return err_
			}
		}
	}
	return nil
}

func (US *UserService) Delete(id_ string) error {
	var errGet error
	var U ents.User
	U, errGet = US.GetById(id_)
	if errGet != nil {
		return errGet
	} else {
		return US.UR.Delete(U)
	}
}
func (US *UserService) GetAll() ([]ents.User, error) {
	Users, err := US.UR.Select()
	if !err {
		return nil, fmt.Errorf("error in get all user service")
	} else {
		return Users, nil
	}
}
func (US *UserService) GetById(id_ string) (ents.User, error) {
	sid, err := strconv.Atoi(id_)
	id := uint64(sid)
	var U ents.User
	if err != nil {
		return U, fmt.Errorf("error incorrect id")
	} else {
		if !US.ExistsById(id) {
			return U, fmt.Errorf(my_errors.ErrNotExists)
		} else {
			var err_ bool
			U, err_ = US.UR.SelectById(id)
			if !err_ {
				return U, fmt.Errorf("error while selecting user by id service")
			}
		}
	}
	return U, nil
}

func (US *UserService) GetByLogin(name_ string) (ents.User, error) {
	var U ents.User
	if !US.ExistsByLogin(name_) {
		return U, fmt.Errorf(my_errors.ErrNotExists)
	} else {
		var err_ bool
		U, err_ = US.UR.SelectByLogin(name_)
		if !err_ {
			return U, fmt.Errorf("error while selecting user by id service")
		}
	}
	return U, nil
}
func (US *UserService) Donate(U *ents.User, DP chk.UserDonateParams) error {

	U.Balance -= DP.Sum_of_money
	U.CharitySum += DP.Sum_of_money
	if DP.IsClosedFoundrising {
		U.ClosedFingAmount += 1
	}
	US.UR.Update(*U)
	return nil
}

func (US *UserService) ReplenishBalance(id_ string, sum float64) error {
	var U ents.User
	var err error
	U, err = US.GetById(id_)
	if err != nil {
		return fmt.Errorf("error in donate user service")
	} else {
		U.Balance += sum
	}
	return nil
}
