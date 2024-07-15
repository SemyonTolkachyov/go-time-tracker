package apperror

import "errors"

var (
	ValidationErr      = AppError{Code: Validation, Message: "validation error"}
	ParseErr           = AppError{Code: Parse, Message: "parse error"}
	UnknownOperatorErr = AppError{Code: UnknownOperator, Message: "unknown operator error"}
	ExistsErr          = AppError{Code: Exists, Message: "already exists"}
	NotFoundErr        = AppError{Code: NotFound, Message: "not found"}
	BadRequestErr      = AppError{Code: BadRequest, Message: "bad request"}
)

type AppError struct {
	Code    string
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

func (e AppError) Is(tgt error) bool {
	var target AppError
	ok := errors.As(tgt, &target)
	if !ok {
		return false
	}
	return e.Code == target.Code
}

func New(code string, message string) error {
	return &AppError{code, message}
}

func NewValidationError(message string) error {
	return New(Validation, message)
}

func NewUnknownOperatorError(message string) error {
	return New(UnknownOperator, message)
}

func NewParseError(message string) error {
	return New(Parse, message)
}

func NewExistsError(message string) error {
	return New(Exists, message)
}

func NewNotFoundError(message string) error {
	return New(NotFound, message)
}

func NewBadRequestError(message string) error {
	return New(BadRequest, message)
}
