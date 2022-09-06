package repositories

import (
	ents "db_course/business/entities"
	"fmt"

	"gorm.io/gorm"
)

type ITransactionRepository interface {
	Insert(T ents.Transaction) error
	Delete(T ents.Transaction) error
	Select() ([]ents.Transaction, bool)
	SelectFrom(type_ bool, id uint64) ([]ents.Transaction, bool)
	SelectTo(type_ bool, id uint64) ([]ents.Transaction, bool)
	SelectById(id uint64) (ents.Transaction, bool)
}

type TransactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(db_ *gorm.DB) *TransactionRepository {
	return &TransactionRepository{DB: db_}
}

func (TR *TransactionRepository) Insert(T ents.Transaction) error {
	var LastTrn ents.Transaction
	TR.DB.Table("transaction_tab").Last(&LastTrn)
	T.Id = LastTrn.Id + 1
	result := TR.DB.Table("transaction_tab").Create(&T)
	if result.Error != nil {
		return fmt.Errorf("error in insert transaction repo")
	} else {
		return nil
	}
}

func (TR *TransactionRepository) Delete(T ents.Transaction) error {
	result := TR.DB.Table("transaction_tab").Delete(&T)
	if result.Error != nil {
		return fmt.Errorf("error in Delete transaction repo")
	} else {
		return nil
	}
}

func (TR *TransactionRepository) Select() ([]ents.Transaction, bool) {
	var transactions []ents.Transaction
	result := TR.DB.Table("transaction_tab").Order("ID").Find(&transactions)
	return transactions, (result.Error == nil)
}
func (TR *TransactionRepository) SelectById(id uint64) (ents.Transaction, bool) {
	var transactions []ents.Transaction
	var t ents.Transaction
	result := TR.DB.Table("transaction_tab").Where("id = ?", id).Find(&transactions)
	if len(transactions) != 0 {
		t = transactions[0]
	}
	return t, (result.Error == nil && len(transactions) != 0)
}

func (TR *TransactionRepository) SelectFrom(type_ bool, id uint64) ([]ents.Transaction, bool) {
	var Transactions []ents.Transaction
	result := TR.DB.Table("transaction_tab").Where("from_essence_type = ? AND from_id = ?", type_, id).Order("ID").Find(&Transactions)
	return Transactions, (result.Error == nil && len(Transactions) != 0)
}

func (TR *TransactionRepository) SelectTo(type_ bool, id uint64) ([]ents.Transaction, bool) {
	var Transactions []ents.Transaction
	result := TR.DB.Table("transaction_tab").Where("to_essence_type = ? AND to_id = ?", type_, id).Order("ID").Find(&Transactions)
	return Transactions, (result.Error == nil && len(Transactions) != 0)
}
func (TR *TransactionRepository) SelectFoundrisingPhilantropIds(foundrising_id uint64) ([]uint64, bool) {
	var Transactions []ents.Transaction
	result := TR.DB.Table("transaction_tab").Where("from_essence_type = ?"+
		" AND to_essence_type =  ? AND to_id = ?",
		ents.FROM_USER, ents.TO_FOUNDRISING, foundrising_id).Order("ID").Find(&Transactions)
	IDs := make([]uint64, len(Transactions))
	for i := range IDs {
		IDs[i] = Transactions[i].From_id
	}
	return IDs, (result.Error == nil /*&& len(Transactions) != 0*/)
}
