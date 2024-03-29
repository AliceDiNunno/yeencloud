package domain

import "net/http"

type ErrorDescription struct {
	HttpCode int    `json:"-"`
	Code     string `json:"code"`
}

var (
	ErrorNotFound   = ErrorDescription{HttpCode: http.StatusNotFound, Code: "PageNotFound"}
	ErrorNoMethod   = ErrorDescription{HttpCode: http.StatusMethodNotAllowed, Code: "MethodNotAllowed"}
	ErrorBadRequest = ErrorDescription{HttpCode: http.StatusBadRequest, Code: "BadRequest"}
	ErrorInternal   = ErrorDescription{HttpCode: http.StatusInternalServerError, Code: "InternalError"}

	ErrorUserNotFound    = ErrorDescription{HttpCode: http.StatusNotFound, Code: "UserNotFound"}
	ErrorProfileNotFound = ErrorDescription{HttpCode: http.StatusNotFound, Code: "ProfileNotFound"}

	ErrorUnableToHashPassword         = ErrorDescription{HttpCode: http.StatusInternalServerError, Code: "UnableToHashPassword"}
	ErrorUnableToGetUserOrganizations = ErrorDescription{HttpCode: http.StatusInternalServerError, Code: "UnableToGetUserOrganizations"}
	ErrorUserAlreadyExists            = ErrorDescription{HttpCode: http.StatusConflict, Code: "UserAlreadyExists"}

	// #YC-14 TODO: this should be moved to the adapter/http/gin package directly
	ErrorAuthenticationTokenMissing = ErrorDescription{HttpCode: http.StatusUnauthorized, Code: "AuthenticationTokenMissing"}
	ErrorSessionNotFound            = ErrorDescription{HttpCode: http.StatusUnauthorized, Code: "SessionNotFound"}
)
