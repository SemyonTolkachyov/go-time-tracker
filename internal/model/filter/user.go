package filter

import (
	"fmt"
	"go-time-tracker/internal/utils"
)

type User struct {
	Id             *int
	CreatedAt      *TimeFilter
	UpdatedAt      *TimeFilter
	PassportNumber *string
	Name           *string
	Surname        *string
	Patronymic     *string
	Address        *string
}

func (u User) String() string {
	def := "<nil>"
	return fmt.Sprintf("{Id: %d, CreatedAt: %s, UpdatedAt: %s, PassportNumber: %s, Name: %s, Surname: %s, Patronymic: %s, Addres: %s}", u.Id, u.CreatedAt, u.UpdatedAt, utils.GetStrValOrDef(u.PassportNumber, &def), utils.GetStrValOrDef(u.Name, &def), utils.GetStrValOrDef(u.Surname, &def), utils.GetStrValOrDef(u.Patronymic, &def), utils.GetStrValOrDef(u.Address, &def))
}
