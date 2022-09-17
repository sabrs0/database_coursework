package entities

type User struct {
	Id               uint64 `gorm:"primaryKey;not null"`
	Login            string `gorm:"not null"`
	Password         string `gorm:"not null"`
	Balance          float64
	CharitySum       float64
	ClosedFingAmount uint64
}

func NewUser() User {
	return User{}
}

func NewUserPtr() *User {
	return &User{}
}

func (U *User) SetLogin(newName string) {
	U.Login = newName
}

func (U *User) SetPassword(newPass string) {
	U.Password = newPass
}
