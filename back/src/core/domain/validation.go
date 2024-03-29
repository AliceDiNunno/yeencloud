package domain

type ValidationFieldName string
type ValidationFieldError string

type ValidationErrors map[ValidationFieldName][]ValidationFieldError
