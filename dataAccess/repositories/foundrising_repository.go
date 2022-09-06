package repositories

import (
	ents "db_course/business/entities"

	"fmt"

	"gorm.io/gorm"
)

type IFoundrisingRepository interface {
	Insert(F ents.Foundrising) error
	Update(F ents.Foundrising) error
	Delete(F ents.Foundrising) error
	Select() ([]ents.Foundrising, bool)
	SelectById(id uint64) (ents.Foundrising, bool)
	SelectByFoundId(id uint64) ([]ents.Foundrising, bool)
	SelectByCreateDate(date string) ([]ents.Foundrising, bool)
	SelectByCloseDate(date string) ([]ents.Foundrising, bool)
}

type FoundrisingRepository struct {
	DB *gorm.DB
}

func NewFoundrisingRepository(db_ *gorm.DB) *FoundrisingRepository {
	return &FoundrisingRepository{DB: db_}
}

func (FR *FoundrisingRepository) Insert(F ents.Foundrising) error {
	var LastFnd ents.Foundrising
	FR.DB.Table("foundrising_tab").Last(&LastFnd)
	F.Id = LastFnd.Id + 1
	result := FR.DB.Table("foundrising_tab").Create(&F)
	if result.Error != nil {
		return fmt.Errorf("error in insert foundrising repo")
	} else {
		return nil
	}
}

func (FR *FoundrisingRepository) Delete(F ents.Foundrising) error {
	result := FR.DB.Table("foundrising_tab").Delete(&F)
	if result.Error != nil {
		return fmt.Errorf("error in Delete foundrising repo")
	} else {
		return nil
	}
}

func (FR *FoundrisingRepository) Update(F ents.Foundrising) error {
	result := FR.DB.Table("foundrising_tab").Save(&F)
	if result.Error != nil {
		return fmt.Errorf("error in Update foundrising repo")
	} else {
		return nil
	}
}

func (FR *FoundrisingRepository) Select() ([]ents.Foundrising, bool) {
	var foundrisings []ents.Foundrising
	result := FR.DB.Table("foundrising_tab").Order("ID").Find(&foundrisings)
	return foundrisings, (result.Error == nil)
}
func (FR *FoundrisingRepository) SelectById(id uint64) (ents.Foundrising, bool) {
	var Foundrisings []ents.Foundrising
	var f ents.Foundrising
	result := FR.DB.Table("foundrising_tab").Where("id = ?", id).Find(&Foundrisings)
	if len(Foundrisings) != 0 {
		f = Foundrisings[0]
	}
	return f, (result.Error == nil && len(Foundrisings) != 0)
}

func (FR *FoundrisingRepository) SelectByFoundId(id uint64) ([]ents.Foundrising, bool) {
	var Foundrising []ents.Foundrising
	result := FR.DB.Table("foundrising_tab").Where("found_id = ?", id).Order("ID").Find(&Foundrising)
	fmt.Println(Foundrising)
	return Foundrising, (result.Error == nil && len(Foundrising) != 0)
}
func (FR *FoundrisingRepository) SelectByFoundIdActive(id uint64) ([]ents.Foundrising, bool) {
	var Foundrising []ents.Foundrising
	result := FR.DB.Table("foundrising_tab").Where("found_id = ? AND closing_date = ", id, "").Order("ID").Find(&Foundrising)
	fmt.Println(Foundrising)
	return Foundrising, (result.Error == nil && len(Foundrising) != 0)
}

func (FR *FoundrisingRepository) SelectByCreateDate(date string) ([]ents.Foundrising, bool) {
	var Foundrising []ents.Foundrising
	result := FR.DB.Table("foundrising_tab").Where("creation_date = ?", date).Order("ID").Find(&Foundrising)
	return Foundrising, (result.Error == nil && len(Foundrising) != 0)
}

func (FR *FoundrisingRepository) SelectByCloseDate(date string) ([]ents.Foundrising, bool) {
	var Foundrising []ents.Foundrising
	result := FR.DB.Table("foundrising_tab").Where("closing_date = ?", date).Order("ID").Find(&Foundrising)
	return Foundrising, (result.Error == nil && len(Foundrising) != 0)
}

func (FR *FoundrisingRepository) SelectByIdAndFoundId(id_ uint64, found_id_ uint64) (ents.Foundrising, bool) {
	var Foundrising []ents.Foundrising
	result := FR.DB.Table("foundrising_tab").Where("id = ? AND found_id = ?", id_, found_id_).Order("ID").Find(&Foundrising)
	var F ents.Foundrising
	if len(Foundrising) > 0 {
		F = Foundrising[0]
	}
	return F, (result.Error == nil && len(Foundrising) != 0)
}
