package domain

import "net/http"

// MARK: - Translatable
var (
	TranslatableUnableToSendVerificationMail = Translatable{Key: "UnableToSendVerificationMail"}
)

// MARK: - Errors
var (
	ErrorUnableToSendVerificationMail = ErrorDescription{HttpCode: http.StatusInternalServerError, Code: TranslatableUnableToSendVerificationMail}
)
