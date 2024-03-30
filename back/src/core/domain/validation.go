package domain

type ValidationFieldName string

type ValidationErrors map[ValidationFieldName][]Translatable

var ValidationErrorUserAlreadyExists = Translatable{Key: "ValidationUserAlreadyExists"}
