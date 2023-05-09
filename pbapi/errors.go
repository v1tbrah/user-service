package pbapi

import "errors"

var (
	ErrEmptyName    = errors.New("empty name")
	ErrEmptySurname = errors.New("empty surname")
)

var (
	ErrUserNotFoundByID     = errors.New("user not found by id")
	ErrInterestNotFoundByID = errors.New("interest not found by id")
	ErrCityNotFoundByID     = errors.New("city not found by id")
)
