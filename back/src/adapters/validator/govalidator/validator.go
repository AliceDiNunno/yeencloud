package govalidator

import (
	"back/src/core/domain"
	"back/src/core/usecases"
	"errors"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

type Validator struct {
	validate           *validator.Validate
	validateTranslator *ut.UniversalTranslator
}

func (v Validator) AddCustomValidation(tag string, fn validator.Func) {
	err := v.validate.RegisterValidation(tag, fn)
	if err != nil {
		log.Err(err).Str("tag", tag).Msg("Error registering custom validation")
		return
	}
}

func (v Validator) ValidateStruct(s interface{}) (bool, domain.ValidationErrors) {
	return v.ValidateStructWithLang(s, "en")
}

func (v Validator) ValidateStructWithLang(s interface{}, lang string) (bool, domain.ValidationErrors) {
	encounteredErrors := domain.ValidationErrors{}

	trans, _ := v.validateTranslator.GetTranslator(lang)

	err := v.validate.Struct(s)
	if err != nil {
		var validationErrors validator.ValidationErrors
		ok := errors.As(err, &validationErrors)

		if !ok {
			log.Err(err).Msg("Error validating struct")
			return false, encounteredErrors
		}

		for _, e := range validationErrors {
			if encounteredErrors[e.Field()] == nil {
				encounteredErrors[e.Field()] = []string{}
			}
			encounteredErrors[e.Field()] = append(encounteredErrors[e.Field()], e.Translate(trans))
		}
	}

	return err == nil, encounteredErrors
}

func NewValidator() usecases.Validator {
	v := Validator{}

	v.validate = validator.New()
	v.LoadLanguages()

	return v
}
