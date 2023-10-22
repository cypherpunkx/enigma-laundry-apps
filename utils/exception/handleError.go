package exception

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func FieldErrors(err error) []map[string]interface{} {
	var errorStack []map[string]interface{}

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, ve := range validationErrors {
			errorStack = append(errorStack, processValidationError(ve))
		}
	}

	return errorStack
}

func processValidationError(ve validator.FieldError) map[string]interface{} {
	errorData := make(map[string]interface{})
	errorData["field"] = ve.Field()
	errorData["tag"] = ve.Tag()

	switch ve.Tag() {
	case "alpha":
		errorData["required"] = fmt.Sprintf("Field %s Harus %s", ve.Field(), ve.Tag())
	case "numeric":
		errorData["required"] = fmt.Sprintf("Field %s Harus %s", ve.Field(), ve.Tag())
	case "max":
		errorData["required"] = fmt.Sprintf("Field %s %s %s karakter", ve.Field(), ve.Tag(), ve.Param())
	}

	return errorData
}
