package usecases

import (
	"github.com/AliceDiNunno/yeencloud/src/adapters/validator"
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
)

func (self UCs) UniqueMailValidator(field validator.FieldToValidate) []domain.ValidationFieldError {
	email := field.FieldValue.String()
	_, err := self.i.Persistence.User.FindUserByEmail(email)
	// If there is no error, it means the user exists so it is not unique therefore we return that there is an error.

	if err == nil {
		return []domain.ValidationFieldError{
			"A user with this email already exists",
		}
	}

	return []domain.ValidationFieldError{}
}
