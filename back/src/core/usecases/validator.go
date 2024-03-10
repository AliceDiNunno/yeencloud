package usecases

import (
	"github.com/go-playground/validator/v10"
)

func (i interactor) PasswordValidator() validator.Func {
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

func (i interactor) UniqueMailValidator() validator.Func {
	return func(fl validator.FieldLevel) bool {
		email := fl.Field().String()
		_, err := i.userRepo.FindUserByEmail(email)
		// if there is no error, it means the user exists so it is not unique therefore we return that there is an error
		if err == nil {
			return false
		}
		return true
	}
}
