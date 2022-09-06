package services

import (
	chk "db_course/business/checker"
	ents "db_course/business/entities"
	repos "db_course/dataAccess/repositories"
	"fmt"
	"strconv"
)

type ITransactionService interface {
	Add(TPars chk.TransactionMainParams) error
	Delete(id_ string) error
	GetAll() ([]ents.Transaction, error)
	GetById(id_ string) (ents.Transaction, error)

	GetFromId(type_ bool, id_ string, FndS FoundationService,
		FndgS FoundrisingService,
		US UserService) ([]ents.Transaction, error)

	GetToId(type_ bool, id_ string, FndS FoundationService,
		FndgS FoundrisingService,
		US UserService) ([]ents.Transaction, error)

	ExistsById(id uint64) bool
}

type TransactionService struct {
	TR repos.TransactionRepository
}

func NewTransactionService(frepo repos.TransactionRepository) TransactionService {
	return TransactionService{TR: frepo}
}

func (FS *TransactionService) ExistsById(id uint64) bool {
	_, result := FS.TR.SelectById(id)
	return result
}

func (FS *TransactionService) Add(TPars chk.TransactionMainParams) error {

	var U ents.Transaction = ents.NewTransaction()
	U.SetFromId(TPars.From_id)
	U.SetFromType(TPars.From_essence_type)
	U.SetComment(TPars.Comment)
	U.SetSumOfMoney(TPars.Sum_of_money)
	U.SetToId(TPars.To_id)
	U.SetToType(TPars.To_essence_type)

	err := FS.TR.Insert(U)

	return err
}
func (FS *TransactionService) Delete(id_ string) error {
	var errGet error
	var U ents.Transaction
	U, errGet = FS.GetById(id_)
	if errGet != nil {
		return errGet
	} else {
		FS.TR.Delete(U)
	}
	return nil
}
func (FS *TransactionService) GetAll() ([]ents.Transaction, error) {
	Transactions, err := FS.TR.Select()
	if !err {
		return nil, fmt.Errorf("не удалось получить список всех транзакций")
	} else {
		return Transactions, nil
	}
}
func (FS *TransactionService) GetById(id_ string) (ents.Transaction, error) {
	sid, err := strconv.Atoi(id_)
	id := uint64(sid)
	var U ents.Transaction
	if err != nil {
		return U, fmt.Errorf("некорректный id")
	} else {
		if !FS.ExistsById(id) {
			return U, fmt.Errorf("несуществующий id")
		} else {
			var err_ bool
			U, err_ = FS.TR.SelectById(id)
			if !err_ {
				return U, fmt.Errorf("не удалось получить транзакцию по id")
			}
		}
	}
	return U, nil
}

// в списке аргументов полный кринжжжж
func (TS *TransactionService) GetFromId(type_ bool, id_ string, FndS FoundationService,
	US UserService) ([]ents.Transaction, error) {
	sid, err := strconv.Atoi(id_)
	id := uint64(sid)
	var U []ents.Transaction
	if err != nil {
		return U, fmt.Errorf("некорректный id")
	} else {
		if type_ == ents.FROM_USER {
			if !US.ExistsById(id) {
				return U, fmt.Errorf("несуществующий id")
			} else {
				var err_ bool
				U, err_ = TS.TR.SelectFrom(type_, id)
				if !err_ {
					return U, fmt.Errorf("не удалось получить транзакцию по id отправителя")
				}
			}
		} else {
			if !FndS.ExistsById(id) {
				return U, fmt.Errorf("несуществующий id")
			} else {
				var err_ bool
				U, err_ = TS.TR.SelectFrom(type_, id)
				if !err_ {
					return U, fmt.Errorf("не удалось получить транзакцию по id отправителя")
				}
			}
		}
		return U, nil
	}
}

func (TS *TransactionService) GetToId(type_ bool, id_ string, FndS FoundationService,
	FndgS FoundrisingService) ([]ents.Transaction, error) {
	sid, err := strconv.Atoi(id_)
	id := uint64(sid)
	var U []ents.Transaction
	if err != nil {
		return U, fmt.Errorf("некорректный id")
	} else {
		if type_ == ents.TO_FOUNDATION {
			if !FndS.ExistsById(id) {
				return U, fmt.Errorf("несуществующий id")
			} else {
				var err_ bool
				U, err_ = TS.TR.SelectTo(type_, id)
				if !err_ {
					return U, fmt.Errorf("не удалось получить транзакцию по id получателя")
				}
			}
		} else {
			if !FndgS.ExistsById(id) {
				return U, fmt.Errorf("несуществующий id")
			} else {
				var err_ bool
				U, err_ = TS.TR.SelectTo(type_, id)
				if !err_ {
					return U, fmt.Errorf("не удалось получить транзакцию по id получателя")
				}
			}
		}
		return U, nil
	}
}

func (TS *TransactionService) GetFoundrisingPhilantropIds(id_ string, FndgS FoundrisingService) ([]uint64, error) {
	sid, err := strconv.Atoi(id_)
	id := uint64(sid)
	var IDs []uint64
	if err != nil {
		return IDs, fmt.Errorf("некорректный id")
	} else {
		if !FndgS.ExistsById(id) {
			return IDs, fmt.Errorf("несуществующий id")
		} else {
			var err_ bool
			IDs, err_ = TS.TR.SelectFoundrisingPhilantropIds(id)
			if !err_ {
				return IDs, fmt.Errorf("не удалось получить id филантропов ")
			}
		}
		return IDs, nil
	}
}
