package gin

import (
	"back/src/core/domain"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"net/http"
)

type Response struct {
	StatusCode int `json:"status"`

	Body  interface{}    `json:"body,omitempty"`
	Error *ResponseError `json:"error,omitempty"`

	RequestID string `json:"requestId,omitempty"`
}

type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (server *ServiceHTTPServer) abortWithError(c *gin.Context, errorDescription domain.ErrorDescription, body ...interface{}) {
	lang := c.GetString("lang")

	msg := i18n.NewLocalizer(server.translator, lang)

	localized, _, _ := msg.LocalizeWithTag(&i18n.LocalizeConfig{
		MessageID: errorDescription.Code,
	})

	c.AbortWithStatusJSON(errorDescription.HttpCode, Response{
		StatusCode: errorDescription.HttpCode,
		Error: &ResponseError{
			Code:    errorDescription.Code,
			Message: localized,
		},
		RequestID: server.getTrace(c).String(),
		Body:      body,
	})
}

func (server *ServiceHTTPServer) success(ctx *gin.Context, body interface{}) {
	ctx.JSON(http.StatusOK, Response{
		StatusCode: http.StatusOK,
		Body:       body,
		RequestID:  server.getTrace(ctx).String(),
	})
}

func (server *ServiceHTTPServer) created(ctx *gin.Context, body interface{}) {
	ctx.JSON(http.StatusCreated, Response{
		StatusCode: http.StatusCreated,
		Body:       body,
		RequestID:  server.getTrace(ctx).String(),
	})
}
