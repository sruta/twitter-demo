package pkg

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

func ValidateStruct(s any) Error {
	err := validator.New().Struct(s)
	if err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			var errs []string
			for _, fieldErr := range validationErrors {
				errs = append(errs, fmt.Sprintf("field '%s' failed for the '%s' constraint", fieldErr.Field(), fieldErr.Tag()))
			}

			return NewGenericError(strings.Join(errs, ", "), err)
		}

		return NewInvalidBodyGenericError(err)
	}

	return nil
}
