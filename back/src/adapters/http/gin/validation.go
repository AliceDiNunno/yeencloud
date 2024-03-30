package gin

import (
	"github.com/gin-gonic/gin"
)

func (server *ServiceHTTPServer) validate(i interface{}, c *gin.Context) bool {
	server.auditer.AddStep(server.getTrace(c))

	succeeded, validationErrors := server.validator.Validate(i)

	if succeeded {
		return true
	}

	// TODO: translation here
	server.abortWithError(c, ErrorBadRequest, validationErrors)

	return false
}
