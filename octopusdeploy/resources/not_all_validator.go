package resources

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// NotAll is the validation function for validating if the current field has a
// value of "all".
func NotAll(fl validator.FieldLevel) bool {
	field := fl.Field()

	if field.Kind() == reflect.String {
		return strings.ToLower(field.String()) != "all"
	}

	return true
}
