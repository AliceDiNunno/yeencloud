package validator

import (
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
)

var ValidationErrorPasswordMustContainAtLeastADigit = domain.Translatable{Key: "ValidationPasswordMustContainAtLeastADigit"}
var ValidationErrorPasswordMustContainAtLeastAnUppercaseLetter = domain.Translatable{Key: "ValidationPasswordMustContainAtLeastAnUppercaseLetter"}
var ValidationErrorPasswordMustContainAtLeastALowercaseLetter = domain.Translatable{Key: "ValidationPasswordMustContainAtLeastALowercaseLetter"}
var ValidationErrorPasswordMustContainAtLeastASpecialCharacter = domain.Translatable{Key: "ValidationPasswordMustContainAtLeastASpecialCharacter"}
var ValidationErrorPasswordMustBeBetween8And64Characters = domain.Translatable{Key: "ValidationPasswordMustBeBetween8And64Characters"}

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
func (validator *Validator) PasswordValidator(field FieldToValidate) []domain.Translatable {
	fieldErrors := []domain.Translatable{}
	password := field.FieldValue.String()

	if len(password) < 8 || len(password) > 64 {
		fieldErrors = append(fieldErrors, ValidationErrorPasswordMustBeBetween8And64Characters)
	}

	if !stringHasNumber(password) {
		fieldErrors = append(fieldErrors, ValidationErrorPasswordMustContainAtLeastADigit)
	}

	if !stringHasUppercaseLetter(password) {
		fieldErrors = append(fieldErrors, ValidationErrorPasswordMustContainAtLeastAnUppercaseLetter)
	}

	if !stringHasLowercaseLetter(password) {
		fieldErrors = append(fieldErrors, ValidationErrorPasswordMustContainAtLeastALowercaseLetter)
	}

	if !stringHasSpecialCharacter(password) {
		fieldErrors = append(fieldErrors, ValidationErrorPasswordMustContainAtLeastASpecialCharacter)
	}

	return fieldErrors
}
