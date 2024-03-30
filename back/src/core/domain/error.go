package domain

import "net/http"

type ErrorDescription struct {
	HttpCode int          `json:"-"`
	Code     Translatable `json:"code"`
}

var (
	TranslatableUserNotFound    = Translatable{Key: "UserNotFound"}
	TranslatableProfileNotFound = Translatable{Key: "ProfileNotFound"}

	TranslatableUnableToHashPassword         = Translatable{Key: "UnableToHashPassword"}
	TranslatableUnableToGetUserOrganizations = Translatable{Key: "UnableToGetUserOrganizations"}
	TranslatableUserAlreadyExists            = Translatable{Key: "UserAlreadyExists"}
	TranslatableSessionNotFound              = Translatable{Key: "SessionNotFound"}
)

var (
	ErrorUserNotFound    = ErrorDescription{HttpCode: http.StatusNotFound, Code: TranslatableUserNotFound}
	ErrorProfileNotFound = ErrorDescription{HttpCode: http.StatusNotFound, Code: TranslatableProfileNotFound}

	ErrorUnableToHashPassword         = ErrorDescription{HttpCode: http.StatusInternalServerError, Code: TranslatableUnableToHashPassword}
	ErrorUnableToGetUserOrganizations = ErrorDescription{HttpCode: http.StatusInternalServerError, Code: TranslatableUnableToGetUserOrganizations}
	ErrorUserAlreadyExists            = ErrorDescription{HttpCode: http.StatusConflict, Code: TranslatableUserAlreadyExists}
	ErrorSessionNotFound              = ErrorDescription{HttpCode: http.StatusUnauthorized, Code: TranslatableSessionNotFound}
)
