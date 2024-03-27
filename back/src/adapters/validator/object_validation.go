package validator

import (
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
)

func (validator *Validator) RequiredValidator(field FieldToValidate) []domain.ValidationFieldError {
	if field.FieldValue.String() == "" {
		return []domain.ValidationFieldError{
			"Field is required",
		}
	}
	return []domain.ValidationFieldError{}
}
