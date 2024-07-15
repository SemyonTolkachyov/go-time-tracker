package output

import "fmt"

type User struct {
	PassportNumber string `json:"passport_number"`
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	Patronymic     string `json:"patronymic"`
	Address        string `json:"address"`
}

type Users struct {
	Users []User `json:"users"`
}

func (u User) String() string {
	return fmt.Sprintf("{PassportNumber: %s, Name: %s, Surname: %s, Patronymic: %s, Address: %s}", u.PassportNumber, u.Name, u.Surname, u.Patronymic, u.Address)
}

func (u Users) String() string {
	return fmt.Sprintf("{Users: %+v}", u.Users)
}
