package entities

type Transaction struct {
	Id                uint64 `gorm:"primaryKey;not null"`
	From_essence_type bool
	From_id           uint64 `gorm:"not null"`
	To_essence_type   bool
	Sum_of_money      float64
	Comment           string
	To_id             uint64 `gorm:"not null"`
}

func NewTransaction() Transaction {
	return Transaction{}
}

func NewTransactionPtr() *Transaction {
	return &Transaction{}
}

func (T *Transaction) SetFromType(type_ bool) {
	T.From_essence_type = type_
}

func (T *Transaction) SetToType(type_ bool) {
	T.To_essence_type = type_
}

func (T *Transaction) SetFromId(id_ uint64) {
	T.From_id = id_
}

func (T *Transaction) SetToId(id_ uint64) {
	T.To_id = id_
}

func (T *Transaction) SetSumOfMoney(sum float64) {
	T.Sum_of_money = sum
}

func (T *Transaction) SetComment(newName string) {
	T.Comment = newName
}

const (
	FROM_USER       = false
	FROM_FOUNDATION = true
	TO_FOUNDATION   = false
	TO_FOUNDRISING  = true
)
