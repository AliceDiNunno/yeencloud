package gin

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/AliceDiNunno/yeencloud/internal/core/domain"
	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int `json:"status"`

	Body  interface{}    `json:"body,omitempty"`
	Error *ResponseError `json:"error,omitempty"`

	RequestID string `json:"requestId,omitempty"`
}

type ResponseError struct {
	Code        string `json:"code"`
	Message     string `json:"message"`
	Translation string `json:"translation,omitempty"`
}

func (server *ServiceHTTPServer) reply(ctx *gin.Context, replyCall func(code int, obj any), code int, body interface{}, err error) {
	if ctx.Writer.Written() {
		return
	}

	ctx.Set(CtxHTTPCodeField, code)

	context := server.getContext(ctx)

	requestError := context.Usecases.EndRequest(code >= 200 && code < 300)

	if requestError != nil {
		body = requestError
	}

	response := Response{
		StatusCode: code,
		Body:       body,
		RequestID:  server.getTrace(ctx).String(),
	}

	var ercode domain.Translatable
	var translatableErr domain.TranslatableError
	if errors.As(err, &translatableErr) {
		ercode = translatableErr.RawKey()
	}

	if err != nil {
		lang := ctx.GetString(CtxLanguageField)

		errorStr := err.Error()
		errs := strings.Split(errorStr, "\n")

		if os.Getenv("ENV") == "production" || os.Getenv("ENV") == "prod" {
			if len(errs) > 1 {
				errorStr = errs[0]
			}
		}

		response.Error = &ResponseError{
			Code:        ercode.RawKey(),
			Message:     errorStr,
			Translation: server.localize.GetLocalizedText(lang, ercode, nil),
		}
	}

	replyCall(code, response)
}

func (server *ServiceHTTPServer) abortWithError(ctx *gin.Context, err error, body ...interface{}) {
	var code domain.RestErrorCode
	if errors.As(err, &code) {
		server.reply(ctx, ctx.AbortWithStatusJSON, code.RestCode(), body, err)
		return
	}

	server.reply(ctx, ctx.AbortWithStatusJSON, 500, body, err)
}

func (server *ServiceHTTPServer) success(ctx *gin.Context, body interface{}) {
	server.reply(ctx, ctx.JSON, http.StatusOK, body, nil)
}

func (server *ServiceHTTPServer) created(ctx *gin.Context, body interface{}) {
	server.reply(ctx, ctx.JSON, http.StatusCreated, body, nil)
}

func (server *ServiceHTTPServer) timedOut(ctx *gin.Context) {
	server.reply(ctx, ctx.JSON, http.StatusRequestTimeout, nil, nil)
}
