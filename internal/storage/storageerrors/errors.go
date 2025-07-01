package storageerrors

import "errors"

var (
	ErrBookAlreadyExist = errors.New("book alredy exist")
	ErrEmptyStorage     = errors.New("book storage is empty")
	ErrBookNoFound      = errors.New("book not found")

	ErrUserAlredyExist = errors.New("user alredy exist")
	ErrInvalidPassword = errors.New("invalid password")
	ErrUserNoExist     = errors.New("user no exist")
)
