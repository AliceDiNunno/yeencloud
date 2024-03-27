package validator

import (
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"net/mail"
)

func (validator *Validator) EmailValidator(field FieldToValidate) []domain.ValidationFieldError {
	email := field.FieldValue.String()

	_, err := mail.ParseAddress(email)

	if err != nil {
		return []domain.ValidationFieldError{
			"Email is not valid",
		}
	}
	return []domain.ValidationFieldError{}
}
