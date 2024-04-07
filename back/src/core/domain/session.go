package domain

import "net/http"

// MARK: - Objects

type Session struct {
	Token    string `json:"token"`
	ExpireAt int64  `json:"expireAt"`
	IP       string `json:"ip"`
	UserID   UserID `json:"userId"`
}

// MARK: - Requests

type NewSession struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	OTPCode  string `json:"otpCode"`

	// Origin of the request to be filled internally (can, and will probably be an IP)
	Origin string `json:"-"` // TODO: remove it and use the request context
}

// MARK: - Translatable

var (
	TranslatableSessionNotFound = Translatable{Key: "SessionNotFound"}
)

// MARK: - Errors

var (
	ErrorSessionNotFound = ErrorDescription{HttpCode: http.StatusUnauthorized, Code: TranslatableSessionNotFound}
)

// MARK: - Logs
var (
	LogScopeSession = LogScope{Identifier: "session"}

	LogScopeSessionRequest = LogScope{Parent: &LogScopeSession, Identifier: "request"}

	LogFieldSessionRequestMail = LogField{Scope: LogScopeSessionRequest, Identifier: "mail"}
)

// MARK: - Functions

// Secure : remove sensitive data from the request.
func (n NewSession) Secure() NewSession {
	n.Password = ""
	// OTP code should be safe as it is time sensitive.
	return n
}
