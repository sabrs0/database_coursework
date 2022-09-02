package checker

import "fmt"

type UserDonateParams struct {
	Sum_of_money        float64
	IsClosedFoundrising bool
}
type UserMainParams struct {
	Login    string
	Password string
}

func CheckLogin(name string) bool {
	for _, c := range name {
		if (c < 'a' || c > 'z') && !(c >= 'A' && c <= 'Z') && !(c >= '0' && c <= '9') {
			return true
		}
	}
	return false
}
func (UP *UserMainParams) CheckParams() error {
	if CheckLogin(UP.Login) {
		return fmt.Errorf("error incorrect username")
	}
	return nil
}

func NewUserMainParams(login string, passw string) UserMainParams {
	return UserMainParams{Login: login, Password: passw}
}

func NewUserDonateParams(sum float64, isClosed bool) UserDonateParams {
	return UserDonateParams{Sum_of_money: sum, IsClosedFoundrising: isClosed}
}
