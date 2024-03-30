package order

import (
	"errors"
)

var (
	ErrOrderNotFound      = errors.New("order not found")
	ErrOrderEmpty         = errors.New("order still empty")
	ErrNullRecordAffected = errors.New("null record affected")
	ErrInvalidID          = errors.New("invalid id")
)
