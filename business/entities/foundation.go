package entities

var Countries [9]string = [9]string{"США", "Россия", "Великобритания", "Канада", "Франция", "Германия", "Китай", "Италия", "Испания"}

type Foundation struct {
	Id                      uint64 `gorm:"primaryKey;not null"`
	Name                    string `gorm:"not null"`
	Password                string `gorm:"not null"`
	CurFoudrisingAmount     uint32
	ClosedFoundrisingAmount uint32
	Fund_balance            float64
	Income_history          float64
	Outcome_history         float64
	Volunteer_amount        uint32
	Country                 string
	Login                   string `gorm:"not null"`
}

func NewFoundation() Foundation {
	return Foundation{}
}

func NewFoundationPtr() *Foundation {
	return &Foundation{}
}

func (F *Foundation) SetLogin(newName string) {
	F.Login = newName
}

func (F *Foundation) SetPassword(newPass string) {
	F.Password = newPass
}

func (F *Foundation) SetName(newName string) {
	F.Name = newName
}

func (F *Foundation) SetCountry(newCntry string) {
	F.Country = newCntry
}
