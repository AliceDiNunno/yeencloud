package domain

import "net/http"

// MARK: - Objects

type ProfileID string

// A profile represents the user's profile and everything that is not related to authentication
// So this is what we're referencing when we want to get organizations, services, settings, etc...
// As in the future the user will be moved to an authentication service, we want to keep them isolated

type Profile struct {
	ID       ProfileID `json:"profileId"`
	UserID   UserID    `json:"userId"`
	Name     string    `json:"name"`
	Language string    `json:"language"`
	Role     string    `json:"role"`
}

// MARK: - Requests

type UpdateProfile struct {
	Name     string `json:"name"`
	Language string `json:"language"`
}

// MARK: - Translatable
var (
	TranslatableProfileNotFound       = Translatable{Key: "ProfileNotFound"}
	TranslatableUnableToCreateProfile = Translatable{Key: "UnableToCreateProfile"}
)

// MARK: - Errors
var (
	ErrorProfileNotFound       = ErrorDescription{HttpCode: http.StatusNotFound, Code: TranslatableProfileNotFound}
	ErrorUnableToCreateProfile = ErrorDescription{HttpCode: http.StatusInternalServerError, Code: TranslatableUnableToCreateProfile}
)

// MARK: - Logs
var (
	LogScopeProfile     = LogScope{Identifier: "profile"}
	LogFieldProfileID   = LogField{Scope: LogScopeProfile, Identifier: "id"}
	LogFieldProfileMail = LogField{Scope: LogScopeProfile, Identifier: "mail"}
	LogFieldProfileName = LogField{Scope: LogScopeProfile, Identifier: "name"}
)

// MARK: - Functions

func (id ProfileID) String() string {
	return string(id)
}
