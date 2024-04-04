package usecases

import (
	"github.com/AliceDiNunno/yeencloud/src/adapters/validator"
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
)

func (self UCs) UniqueMailValidator(field validator.FieldToValidate) []domain.Translatable {
	email := field.FieldValue.String()
	_, err := self.i.Persistence.FindUserByEmail(email)
	// If there is no error, it means the user exists so it is not unique therefore we return that there is an error.

	if err == nil {
		return []domain.Translatable{domain.ValidationErrorUserAlreadyExists}
	}

	return []domain.Translatable{}
}
