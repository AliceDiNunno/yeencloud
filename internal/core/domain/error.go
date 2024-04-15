package domain

import (
	"fmt"
)

type RestErrorCode interface {
	RestCode() int
}

// MARK: - Usecases Error

type UsecasesError struct {
	Msg string
	Key Translatable
}

func (e *UsecasesError) Error() string {
	return fmt.Sprintf("usecases: %v", e.Msg)
}

func (e *UsecasesError) RawKey() Translatable {
	return e.Key
}

// MARK: - Resource Not Found

type ResourceNotFoundError struct {
	Id   string
	Type string
}

func (e *ResourceNotFoundError) Error() string {
	return fmt.Sprintf("ressource not found: '%v' of type '%v'", e.Id, e.Type)
}

func (e *ResourceNotFoundError) RestCode() int {
	return 404
}

// MARK: - Log
var LogFieldError = LogField{Identifier: "error"}
