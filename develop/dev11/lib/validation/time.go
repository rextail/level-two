package validation

import (
	"github.com/go-playground/validator/v10"
	"time"
)

// ValidateDate проверяет, возможно ли преобразовать строку в дату.
func ValidateDate(fl validator.FieldLevel) bool {
	const op = `lib.validation.time.ParseDate`

	t := fl.Field().String()

	_, err := time.Parse(time.DateOnly, t)

	return err == nil
}
