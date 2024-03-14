package usecases

import (
	"github.com/go-playground/validator/v10"
)

func stringHasNumber(value string) bool {
	for _, char := range value {
		if char >= '0' && char <= '9' {
			return true
		}
	}
	return false
}

func stringHasUppercaseLetter(value string) bool {
	for _, char := range value {
		if char >= 'A' && char <= 'Z' {
			return true
		}
	}
	return false
}

func stringHasLowercaseLetter(value string) bool {
	for _, char := range value {
		if char >= 'a' && char <= 'z' {
			return true
		}
	}
	return false
}

func stringHasSpecialCharacter(value string) bool {
	allowedSpecialChars := []rune{'!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '+', '=', '{', '}', '[', ']', '|', '\\', ':', ';', '"', '\'', '<', '>', ',', '.', '?', '/', '`', '~'}
	for _, char := range value {
		for _, specialChar := range allowedSpecialChars {
			if char == specialChar {
				return true
			}
		}
	}
	return false
}

func (i interactor) PasswordValidator() validator.Func {
	return func(fl validator.FieldLevel) bool {
		password := fl.Field().String()

		if len(password) < 8 || len(password) > 64 {
			return false
		}

		if !stringHasNumber(password) {
			return false
		}

		if !stringHasUppercaseLetter(password) {
			return false
		}

		if !stringHasLowercaseLetter(password) {
			return false
		}

		return stringHasSpecialCharacter(password)
	}
}

func (i interactor) UniqueMailValidator() validator.Func {
	return func(fl validator.FieldLevel) bool {
		email := fl.Field().String()
		_, err := i.userRepo.FindUserByEmail(email)
		// If there is no error, it means the user exists so it is not unique therefore we return that there is an error.

		return err != nil
	}
}
