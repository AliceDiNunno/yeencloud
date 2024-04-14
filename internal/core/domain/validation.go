package domain

// MARK: - Objects

type ValidationFieldName string
type ValidationErrors map[ValidationFieldName][]Translatable

// MARK: - Translatable

var ValidationErrorUserAlreadyExists = Translatable{Key: "ValidationUserAlreadyExists"}
