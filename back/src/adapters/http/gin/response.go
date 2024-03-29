package gin

import (
	"back/src/core/domain"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Response struct {
	StatusCode int

	Body  interface{}    `json:",omitempty"`
	Error *ResponseError `json:",omitempty"`
}

type ResponseError struct {
	Code    string
	Message string
}

func (server *ServiceHTTPServer) abortWithError(c *gin.Context, error domain.ErrorDescription) {
	lang := c.GetString("lang")

	msg := i18n.NewLocalizer(server.translator, lang)

	localized, _, _ := msg.LocalizeWithTag(&i18n.LocalizeConfig{
		MessageID: error.Code,
	})

	c.AbortWithStatusJSON(error.HttpCode, Response{
		StatusCode: error.HttpCode,
		Error: &ResponseError{
			Code:    error.Code,
			Message: localized,
		},
	})
}
