package controllers

import (
	"db_course/business/checker"
	ents "db_course/business/entities"
	servs "db_course/business/services"
	"fmt"
)

type FoundationController struct {
	FS servs.FoundationService
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
	return UC.FS.Delete(id)

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
