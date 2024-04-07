package domain

// MARK: - Objects

type ErrorDescription struct {
	HttpCode  int                     `json:"-"`
	Code      Translatable            `json:"code"`
	Arguments TranslatableArgumentMap `json:"-"`
}

// MARK: - Log

var LogFieldError = LogField{Identifier: "error"}
