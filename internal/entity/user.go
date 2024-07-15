package entity

import (
	"fmt"
	"time"
)

type User struct {
	Id             int        `db:"id"`
	CreatedAt      time.Time  `db:"created_at"`
	UpdatedAt      *time.Time `db:"updated_at"`
	DeletedAt      *time.Time `db:"deleted_at"`
	PassportNumber string     `db:"passport_number"`
	Name           string     `db:"name"`
	Surname        string     `db:"surname"`
	Patronymic     string     `db:"patronymic"`
	Address        string     `db:"address"`
}

func (u User) String() string {
	return fmt.Sprintf("{Id: %d, CreatedAt: %s, UpdatedAt: %s, DeletedAt: %s, PasswordNumber: %s, Name: %s, Surname: %s, Patronymic: %s, Address: %s}", u.Id, u.CreatedAt, u.UpdatedAt, u.DeletedAt, u.PassportNumber, u.Name, u.Surname, u.Patronymic, u.Address)
}
