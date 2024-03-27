package validator

import (
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"reflect"
	"strings"
)

// Other validation packages did not satisfy me, so I decided to write my own.
// I need a validator that can be extensible, with custom validation errors and the ability to translate them.
// I used go-playground/validator which was almost perfect in that sense, but it was missing the ability to add custom validation errors.
// It won't be as feature-complete however as I will only implement the features I need.

type ValidationFunc func(FieldToValidate) []domain.ValidationFieldError
type ValidationFuncs map[string]ValidationFunc

var NoValidationFuncError = domain.ValidationFieldError("No validation function found for this field")

type Validator struct {
	ValidationFuncs ValidationFuncs
}

type FieldToValidate struct {
	FieldName  domain.ValidationFieldName
	FieldType  reflect.Type
	FieldValue reflect.Value
	Tag        string
}

func (validator Validator) RegisterValidation(tag string, fn ValidationFunc) {
	validator.ValidationFuncs[tag] = fn
}

/*
Validate runs checks on all the struct fields.
It returns true and an empty array if all the checks pass.
If at least one check fails, it returns false and an array of ValidationFieldError.
*/
func (validator Validator) Validate(s interface{}) (bool, []domain.ValidationFieldError) {
	var errors []domain.ValidationFieldError

	st := reflect.TypeOf(s)

	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		fieldTag := field.Tag.Get("validate")

		if fieldTag == "" {
			continue
		}

		fieldName := domain.ValidationFieldName(field.Name)
		validationSteps := strings.Split(fieldTag, ",")

		reflect.ValueOf(s).FieldByName(field.Name)

		fieldToValidate := FieldToValidate{
			FieldName:  fieldName,
			FieldType:  field.Type,
			FieldValue: reflect.ValueOf(s).FieldByName(field.Name),
			Tag:        fieldTag,
		}

		for _, step := range validationSteps {
			validationFunc, ok := validator.ValidationFuncs[step]
			if !ok {
				errors = append(errors, NoValidationFuncError)
				continue
			}

			validationErrors := validationFunc(fieldToValidate)
			if len(validationErrors) > 0 {
				return false, validationErrors
			}
		}
	}

	return len(errors) == 0, errors
}

func NewValidator() Validator {
	v := Validator{
		ValidationFuncs: make(ValidationFuncs),
	}

	v.RegisterValidation("email", v.EmailValidator)
	v.RegisterValidation("password", v.PasswordValidator)
	v.RegisterValidation("required", v.RequiredValidator)

	return v
}
