package storage

import "errors"

var (
	ErrCityAlreadyExists     = errors.New("city already exists")
	ErrInterestAlreadyExists = errors.New("interest already exists")
)
