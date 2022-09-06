package repositories

import (
	ents "db_course/business/entities"
	"fmt"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Insert(U ents.User) error
	Update(U ents.User) error
	Delete(U ents.User) error
	Select() ([]ents.User, bool)
	SelectById(id uint64) (ents.User, bool)
	SelectByName(name string) (ents.User, bool)
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db_ *gorm.DB) *UserRepository {
	return &UserRepository{DB: db_}
}

func (UR *UserRepository) Insert(U ents.User) error {
	var LastUsr ents.User
	UR.DB.Table("user_tab").Last(&LastUsr)
	U.Id = LastUsr.Id + 1
	result := UR.DB.Table("user_tab").Create(&U)
	if result.Error != nil {
		return fmt.Errorf("error in insert user repo")
	} else {
		return nil
	}
}

func (UR *UserRepository) Delete(U ents.User) error {
	result := UR.DB.Table("user_tab").Delete(&U)
	if result.Error != nil {
		return fmt.Errorf("error in Delete user repo")
	} else {
		return nil
	}
}

func (UR *UserRepository) Update(U ents.User) error {
	result := UR.DB.Table("user_tab").Save(&U)
	if result.Error != nil {
		return fmt.Errorf("error in Update user repo")
	} else {
		return nil
	}
}

func (UR *UserRepository) Select() ([]ents.User, bool) {
	var user_tab []ents.User
	result := UR.DB.Table("user_tab").Order("ID").Find(&user_tab)
	return user_tab, (result.Error == nil)
}
func (UR *UserRepository) SelectById(id uint64) (ents.User, bool) {
	/*var User []ents.User
	result := UR.DB.Table("user_tab").Find(&User, []uint64{id})*/
	var Users []ents.User
	var U ents.User
	result := UR.DB.Table("user_tab").Where("id = ?", id).Find(&Users)
	if len(Users) != 0 {
		U = Users[0]
	}
	return U, (result.Error == nil && len(Users) != 0)
}

func (UR *UserRepository) SelectByLogin(name string) (ents.User, bool) {
	var Users []ents.User
	var U ents.User
	result := UR.DB.Table("user_tab").Where("login = ?", name).Find(&Users)
	if len(Users) != 0 {
		U = Users[0]
	}
	return U, (result.Error == nil && len(Users) != 0)
}
