package validation

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateCoolVideoTitle(fl validator.FieldLevel) bool {
	return strings.Contains(fl.Field().String(), "Cool")
}