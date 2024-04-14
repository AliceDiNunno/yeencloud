package validator

import (
	"net/mail"

	"github.com/AliceDiNunno/yeencloud/internal/core/domain"
)

var ValidationErrorMailIsInvalid = domain.Translatable{Key: "ValidationMailIsInvalid"}

func (validator *Validator) EmailValidator(field FieldToValidate) []domain.Translatable {
	email := field.FieldValue.String()

	_, err := mail.ParseAddress(email)

	if err != nil {
		return []domain.Translatable{ValidationErrorMailIsInvalid}
	}
	return []domain.Translatable{}
}
