package gin

import (
	"back/src/core/domain"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"net/http"
)

type Response struct {
	StatusCode int

	Body  interface{}    `json:",omitempty"`
	Error *ResponseError `json:",omitempty"`

	RequestID string `json:",omitempty"`
}

type ResponseError struct {
	Code    string
	Message string
}

func (server *ServiceHTTPServer) abortWithError(c *gin.Context, error domain.ErrorDescription, body ...interface{}) {
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
		RequestID: server.getTrace(c).String(),
		Body:      body,
	})
}

func (server *ServiceHTTPServer) success(c *gin.Context, body interface{}) {
	c.JSON(http.StatusOK, Response{
		StatusCode: http.StatusOK,
		Body:       body,
		RequestID:  server.getTrace(c).String(),
	})
}

func (server *ServiceHTTPServer) created(c *gin.Context, body interface{}) {
	c.JSON(http.StatusCreated, Response{
		StatusCode: http.StatusCreated,
		Body:       body,
		RequestID:  server.getTrace(c).String(),
	})
}
