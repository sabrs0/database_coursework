package controllers

import (
	ents "db_course/business/entities"
	servs "db_course/business/services"
	repos "db_course/dataAccess/repositories"

	"strconv"

	"gorm.io/gorm"
)

type TransactionController struct {
	TS  servs.TransactionService
	FS  servs.FoundationService
	FgS servs.FoundrisingService
	US  servs.UserService
}

func NewTransactionController(db *gorm.DB, fs servs.FoundationService, fgs servs.FoundrisingService,
	us servs.UserService) TransactionController {
	FR := repos.NewTransactionRepository(db)
	tS := servs.NewTransactionService(*FR)
	return TransactionController{TS: tS, FS: fs, FgS: fgs, US: us}
}

func (TC *TransactionController) GetAll() ([]ents.Transaction, error) {
	return TC.TS.GetAll()
}
func (TC *TransactionController) GetByID(id_ string) (ents.Transaction, error) {
	return TC.TS.GetById(id_)
}
func (TC *TransactionController) GetFromId(type_ string, id_ string) ([]ents.Transaction, error) {
	booltype, err := strconv.ParseBool(type_)
	var Transactions []ents.Transaction
	if err == nil {
		Transactions, err = TC.TS.GetFromId(booltype, id_, TC.FS, TC.US)
	}
	return Transactions, err
}
func (TC *TransactionController) GetToId(type_ string, id_ string) ([]ents.Transaction, error) {
	booltype, err := strconv.ParseBool(type_)
	var Transactions []ents.Transaction
	if err == nil {
		Transactions, err = TC.TS.GetToId(booltype, id_, TC.FS, TC.FgS)
	}
	return Transactions, err
}

func (TC *TransactionController) Delete(id string) error {
	return TC.TS.Delete(id)

}

func (TC *TransactionController) GetFoundrisingPhilantropIds(id_ string) ([]uint64, error) {
	return TC.TS.GetFoundrisingPhilantropIds(id_, TC.FgS)
}
