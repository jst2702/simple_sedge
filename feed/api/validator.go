package api

import "github.com/go-playground/validator/v10"

var validPassword validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if password, ok := fieldLevel.Field().Interface().(string); ok {
		if len(password) < 6 {
			return false
		} else {
			return true
		}
	}
	// otherwise, the field is not a string.
	return false
}
