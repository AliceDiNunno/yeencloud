package validator

import (
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
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

// PasswordValidator validates a password using certain rules
// We could use a regex to validate it but then we can't have precise error messages for each rule
func (validator *Validator) PasswordValidator(field FieldToValidate) []domain.ValidationFieldError {
	fieldErrors := []domain.ValidationFieldError{}
	password := field.FieldValue.String()

	if len(password) < 8 || len(password) > 64 {
		fieldErrors = append(fieldErrors, "Password must be between 8 and 64 characters")
	}

	if !stringHasNumber(password) {
		fieldErrors = append(fieldErrors, "Password must contain at least one number")
	}

	if !stringHasUppercaseLetter(password) {
		fieldErrors = append(fieldErrors, "Password must contain at least one uppercase letter")
	}

	if !stringHasLowercaseLetter(password) {
		fieldErrors = append(fieldErrors, "Password must contain at least one lowercase letter")
	}

	if !stringHasSpecialCharacter(password) {
		fieldErrors = append(fieldErrors, "Password must contain at least one special character")
	}

	return fieldErrors
}
