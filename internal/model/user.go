package model

import "fmt"

type UserInfo struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Address    string `json:"address"`
}

type UserData struct {
	Name           string
	Surname        string
	Patronymic     string
	Address        string
	PassportNumber string
}

func (u UserData) String() string {
	return fmt.Sprintf("{PassportNumber: %s, Name: %s, Surname: %s, Patronymic: %s, Address: %s}", u.PassportNumber, u.Name, u.Surname, u.Patronymic, u.Address)
}

func (u UserInfo) String() string {
	return fmt.Sprintf("{Name: %s, Surname: %s, Patronymic: %s, Address: %s}", u.Name, u.Surname, u.Patronymic, u.Address)
}
