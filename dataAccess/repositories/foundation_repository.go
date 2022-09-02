package repositories

import (
	ents "db_course/business/entities"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type IFoundationRepository interface {
	Insert(F ents.Foundation) error
	Update(F ents.Foundation) error
	Delete(F ents.Foundation) error
	Select() ([]ents.Foundation, bool)
	SelectById(id uint64) (ents.Foundation, bool)
	SelectByLogin(name string) (ents.Foundation, bool)
}

type FoundationRepository struct {
	DB *gorm.DB
}

func NewFoundationRepository(db_ *gorm.DB) *FoundationRepository {
	return &FoundationRepository{DB: db_}
}

func (FR *FoundationRepository) Insert(F ents.Foundation) error {
	var LastUsr ents.Foundation
	FR.DB.Table("foundation_tab").Last(&LastUsr)
	F.Id = LastUsr.Id + 1
	result := FR.DB.Table("foundation_tab").Create(&F)
	if result.Error != nil {
		return fmt.Errorf("error in insert Foundation repo")
	} else {
		return nil
	}
}

func (FR *FoundationRepository) Delete(F ents.Foundation) error {
	result := FR.DB.Table("foundation_tab").Delete(&F)
	if result.Error != nil {
		return fmt.Errorf("error in Delete Foundation repo")
	} else {
		return nil
	}
}

func (FR *FoundationRepository) Update(F ents.Foundation) error {
	result := FR.DB.Table("foundation_tab").Save(&F)
	if result.Error != nil {
		return fmt.Errorf("error in Update Foundation repo")
	} else {
		return nil
	}
}

func (FR *FoundationRepository) Select() ([]ents.Foundation, bool) {
	var foundation_tab []ents.Foundation
	result := FR.DB.Table("foundation_tab").Order("ID").Find(&foundation_tab)
	return foundation_tab, (result.Error == nil)
}
func (FR *FoundationRepository) SelectById(id uint64) (ents.Foundation, bool) {
	var Foundation ents.Foundation
	result := FR.DB.Table("foundation_tab").Where("id = ?", id).First(&Foundation)
	return Foundation, (result.Error == nil && !errors.Is(result.Error, gorm.ErrRecordNotFound))
}

func (FR *FoundationRepository) SelectByLogin(name string) (ents.Foundation, bool) {
	var Foundation ents.Foundation
	result := FR.DB.Table("foundation_tab").Where("login = ?", name).First(&Foundation)
	return Foundation, (result.Error == nil && !errors.Is(result.Error, gorm.ErrRecordNotFound))
}
func (FR *FoundationRepository) SelectByName(name string) (ents.Foundation, bool) {
	var Foundation ents.Foundation
	result := FR.DB.Table("foundation_tab").Where("name = ?", name).Find(&Foundation)
	return Foundation, (result.Error == nil && Foundation.Id != 0)
}
func (FR *FoundationRepository) SelectByCountry(country string) ([]ents.Foundation, bool) {
	var Foundation []ents.Foundation
	result := FR.DB.Table("foundation_tab").Where("country = ?", country).Order("ID").Find(&Foundation)
	return Foundation, (result.Error == nil && len(Foundation) != 0)
}
