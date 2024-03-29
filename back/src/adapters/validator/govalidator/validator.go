package govalidator

import (
	"back/src/core/domain"
	"back/src/core/usecases"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate           *validator.Validate
	validateTranslator *ut.UniversalTranslator
}

func (v Validator) AddCustomValidation(tag string, fn validator.Func) {
	v.validate.RegisterValidation(tag, fn)
}

func NewValidator() usecases.Validator {
	v := Validator{}

	v.validate = validator.New()
	v.LoadLanguages()

	return v
}

func (v Validator) ValidateStruct(s interface{}) (bool, domain.ValidationErrors) {
	return v.ValidateStructWithLang(s, "en")
}

func (v Validator) ValidateStructWithLang(s interface{}, lang string) (bool, domain.ValidationErrors) {
	validationErrors := domain.ValidationErrors{}

	trans, _ := v.validateTranslator.GetTranslator(lang)

	err := v.validate.Struct(s)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			if validationErrors[e.Field()] == nil {
				validationErrors[e.Field()] = []string{}
			}
			validationErrors[e.Field()] = append(validationErrors[e.Field()], e.Translate(trans))
		}
	}

	return err == nil, validationErrors
}
