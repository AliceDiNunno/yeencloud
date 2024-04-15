package gin

import (
	"fmt"

	"github.com/AliceDiNunno/yeencloud/internal/core/domain"
	"github.com/gin-gonic/gin"
)

// MARK: - PageNotFoundError
var TranslatablePageNotFound = domain.Translatable{Key: "PageNotFound"}

type PageNotFoundError struct {
	Msg string
	Key domain.Translatable

	Method string
	Path   string
}

func (e *PageNotFoundError) Error() string {
	arg := ""
	if e.Path != "" {
		arg = fmt.Sprintf("(%v %v)", e.Method, e.Path)
	}
	return fmt.Sprintf("http: page not found: %v %v", e.Msg, arg)
}

func (e *PageNotFoundError) RawKey() domain.Translatable {
	return TranslatablePageNotFound
}

// MARK: BadRequestError
var TranslatableBadRequest = domain.Translatable{Key: "BadRequest"}

type BadRequestError struct {
}

func (e *BadRequestError) Error() string {
	return "http: bad request"
}

func (e *BadRequestError) RawKey() domain.Translatable {
	return TranslatableBadRequest
}

// MARK: - InternalServerError
var TranslatableInternalServerError = domain.Translatable{Key: "InternalServerError"}

type InternalServerError struct {
}

func (e *InternalServerError) Error() string {
	return "http: internal server error"
}

func (e *InternalServerError) RawKey() domain.Translatable {
	return TranslatableInternalServerError
}

// MARK: - UnauthorizedError
// No authentication credentials were provided.

var TranslatableUnauthorizedError = domain.Translatable{Key: "Unauthorized"}

type UnauthorizedError struct {
}

func (e *UnauthorizedError) Error() string {
	return "http: unauthorized"
}

func (e *UnauthorizedError) RawKey() domain.Translatable {
	return TranslatableUnauthorizedError
}

func (e *UnauthorizedError) RestCode() int {
	return 401
}

// MARK: - ForbiddenError
// The authenticated user lacks the necessary permissions to perform this action.

var TranslatableForbiddenError = domain.Translatable{Key: "Forbidden"}

type ForbiddenError struct {
}

func (e *ForbiddenError) Error() string {
	return "http: forbidden"
}

func (e *ForbiddenError) RawKey() domain.Translatable {
	return TranslatableForbiddenError
}

func (e *ForbiddenError) RestCode() int {
	return 403
}

// MARK: - AuthenticationRequiredError
var TranslatableAuthenticationRequired = domain.Translatable{Key: "AuthenticationRequired"}

type AuthenticationRequiredError struct {
}

func (e *AuthenticationRequiredError) Error() string {
	return "http: authentication required"
}

func (e *AuthenticationRequiredError) RawKey() domain.Translatable {
	return TranslatableAuthenticationRequired
}

// MARK: - MethodRequiredError
var TranslatableMethodNotAllowed = domain.Translatable{Key: "MethodNotAllowed"}

type MethodNotAllowedError struct {
}

func (e *MethodNotAllowedError) Error() string {
	return fmt.Sprintf("http: method not allowed")
}

func (e *MethodNotAllowedError) RawKey() domain.Translatable {
	return TranslatablePageNotFound
}

func pageNotFoundError(method string, path string) error {
	page := PageNotFoundError{}
	page.Method = method
	page.Path = path
	return &page
}

func (server *ServiceHTTPServer) recoverFromPanic(c *gin.Context, err interface{}) {
	server.abortWithError(c, &InternalServerError{}, err)
}
