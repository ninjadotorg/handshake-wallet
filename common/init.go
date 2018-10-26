package common

import (
	"github.com/go-playground/validator"
	"github.com/shopspring/decimal"
)

var DataValidator = NewValidator()

func NewValidator() *validator.Validate {
	return validator.New()
}

func Float64ToDecimal(value float64) decimal.Decimal {
	return decimal.NewFromFloat(value)
}
