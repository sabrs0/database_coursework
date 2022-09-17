package entities

import "database/sql"

const DateFormat string = "2006-01-02"

type Foundrising struct {
	Id                 uint64 `gorm:"primaryKey;not null"`
	Found_id           uint64 `gorm:"not null"`
	Description        string
	Required_sum       float64
	Current_sum        float64
	Philantrops_amount uint64
	Creation_date      string
	Closing_date       sql.NullString
}

func NewFoundrising() Foundrising {
	return Foundrising{}
}

func NewFoundrisingPtr() *Foundrising {
	return &Foundrising{}
}

func (F *Foundrising) SetDescription(n string) {
	F.Description = n
}

func (F *Foundrising) SetReqSum(sum float64) {
	F.Required_sum = sum
}
func (F *Foundrising) SetCreateDate(newName string) {
	F.Creation_date = newName
}

func (F *Foundrising) SetFoundId(id uint64) {
	F.Found_id = id
}
