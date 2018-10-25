package common

import (
	"github.com/go-playground/validator"
)

var DataValidator = NewValidator()

func NewValidator() *validator.Validate {
	return validator.New()
}
