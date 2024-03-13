package domain

import "net/http"

type ErrorDescription struct {
	HttpCode int
	Code     string
}

type Error struct {
	ErrorDescription

	Stack          Stack
	Child          *Error                 `json:",omitempty"`
	AdditionalData map[string]interface{} `json:",omitempty"`
	Fingerprint    string
}

var (
	ErrorNotFound   = ErrorDescription{HttpCode: http.StatusNotFound, Code: "PageNotFound"}
	ErrorNoMethod   = ErrorDescription{HttpCode: http.StatusMethodNotAllowed, Code: "MethodNotAllowed"}
	ErrorBadRequest = ErrorDescription{HttpCode: http.StatusBadRequest, Code: "BadRequest"}

	ErrorUserNotFound    = ErrorDescription{HttpCode: http.StatusNotFound, Code: "UserNotFound"}
	ErrorProfileNotFound = ErrorDescription{HttpCode: http.StatusNotFound, Code: "ProfileNotFound"}

	ErrorUnableToGetUserOrganizations = ErrorDescription{HttpCode: http.StatusInternalServerError, Code: "UnableToGetUserOrganizations"}

	// #YC-14 TODO: this should be moved to the adapter/http/gin package directly
	ErrorAuthenticationTokenMissing = ErrorDescription{HttpCode: http.StatusUnauthorized, Code: "AuthenticationTokenMissing"}
	ErrorSessionNotFound            = ErrorDescription{HttpCode: http.StatusUnauthorized, Code: "SessionNotFound"}
)
