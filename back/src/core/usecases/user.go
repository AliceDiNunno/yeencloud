package usecases

import (
	"back/src/core/domain"
	"back/src/core/domain/requests"
	"github.com/go-playground/validator/v10"
)

func (i interactor) CreateUser(user requests.NewUser) (domain.User, error) {
	return domain.User{}, nil
}

// TODO: add better validation system that allows for custom error messages
func (i *interactor) PasswordValidator() validator.Func {
	return func(fl validator.FieldLevel) bool {
		value := fl.Field().String()

		if len(value) < 8 {
			return false
		}
		if len(value) > 64 {
			return false
		}

		//check for at least one number
		hasNumber := false
		for _, char := range value {
			if char >= '0' && char <= '9' {
				hasNumber = true
				break
			}
		}
		if !hasNumber {
			return false
		}

		//check for at least one uppercase letter
		hasUppercase := false
		for _, char := range value {
			if char >= 'A' && char <= 'Z' {
				hasUppercase = true
				break
			}
		}

		if !hasUppercase {
			return false
		}

		//check for at least one lowercase letter
		hasLowercase := false
		for _, char := range value {
			if char >= 'a' && char <= 'z' {
				hasLowercase = true
				break
			}
		}

		if !hasLowercase {
			return false
		}

		//check for at least one special character
		hasSpecial := false
		allowedSpecialChars := []rune{'!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '+', '=', '{', '}', '[', ']', '|', '\\', ':', ';', '"', '\'', '<', '>', ',', '.', '?', '/', '`', '~'}
		for _, char := range value {
			for _, specialChar := range allowedSpecialChars {
				if char == specialChar {
					hasSpecial = true
					break
				}
			}
		}

		if !hasSpecial {
			return false
		}

		return true
	}
}
