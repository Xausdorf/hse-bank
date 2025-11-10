package factory

import "errors"

var (
	ErrEmptyName      = errors.New("name cannot be empty")
	ErrNegativeAmount = errors.New("amount cannot be negative")
)
