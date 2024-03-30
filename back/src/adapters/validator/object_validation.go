package validator

import (
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
)

var ValidationErrorFieldIsRequired = domain.Translatable{Key: "ValidationFieldIsRequired"}

func (validator *Validator) RequiredValidator(field FieldToValidate) []domain.Translatable {
	if field.FieldValue.String() == "" {
		return []domain.Translatable{ValidationErrorFieldIsRequired}
	}
	return []domain.Translatable{}
}
