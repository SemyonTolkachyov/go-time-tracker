package input

import (
	"fmt"
	"go-time-tracker/internal/apperror"
	"go-time-tracker/internal/utils"
)

type NewUser struct {
	PassportNumber string `json:"passportNumber" binding:"required"`
}

type UpdateUser struct {
	Id             int     `json:"id" binding:"required"`
	PassportNumber *string `json:"passport_number"`
	Name           *string `json:"name"`
	Surname        *string `json:"surname"`
	Patronymic     *string `json:"patronymic"`
	Address        *string `json:"address"`
}

func (n NewUser) String() string {
	return fmt.Sprintf("{PassportNumber: %s}", n.PassportNumber)
}

func (u UpdateUser) String() string {
	def := "<nil>"
	return fmt.Sprintf("{Id: %d, PassportNumber: %s, Name: %s, Surname: %s, Patronymic: %s, Address: %s}", u.Id, utils.GetStrValOrDef(u.PassportNumber, &def), utils.GetStrValOrDef(u.Name, &def), utils.GetStrValOrDef(u.Surname, &def), utils.GetStrValOrDef(u.Patronymic, &def), utils.GetStrValOrDef(u.Address, &def))
}

func (u UpdateUser) Validate() error {
	if u.PassportNumber == nil && u.Name == nil && u.Surname == nil && u.Patronymic == nil && u.Address == nil {
		return apperror.NewValidationError("update structure has no values")
	}
	return nil
}
