package storageerrors

import "errors"

var (
	ErrBookAlreadyExist = errors.New("book already exist")
	ErrEmptyStorage     = errors.New("book storage empty")
	ErrBookNoFound      = errors.New("book not found")
)
