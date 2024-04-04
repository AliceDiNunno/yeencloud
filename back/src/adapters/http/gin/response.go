package gin

import (
	"net/http"

	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/gin-gonic/gin"
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

func (server *ServiceHTTPServer) reply(ctx *gin.Context, replyCall func(code int, obj any), code int, errDesc *domain.ErrorDescription, body interface{}) {
	ctx.Set(CtxHTTPCodeField, code)

	context := server.getContext(ctx)

	requestError := context.Usecases.EndRequest(code >= 200 && code < 300)

	if requestError != nil {
		body = requestError
		errDesc = &ErrorInternal
		code = errDesc.HttpCode
	}

	response := Response{
		StatusCode: code,
		Body:       body,
		RequestID:  server.getTrace(ctx).String(),
	}

	if errDesc != nil {
		lang := ctx.GetString(CtxLanguageField)

		response.Error = &ResponseError{
			Code:    errDesc.Code.RawKey(),
			Message: server.localize.GetLocalizedText(lang, errDesc.Code),
		}
	}

	replyCall(code, response)
}

func (server *ServiceHTTPServer) abortWithError(ctx *gin.Context, errorDescription domain.ErrorDescription, body ...interface{}) {
	server.reply(ctx, ctx.AbortWithStatusJSON, errorDescription.HttpCode, &errorDescription, body)
}

func (server *ServiceHTTPServer) success(ctx *gin.Context, body interface{}) {
	server.reply(ctx, ctx.JSON, http.StatusOK, nil, body)
}

func (server *ServiceHTTPServer) created(ctx *gin.Context, body interface{}) {
	server.reply(ctx, ctx.JSON, http.StatusCreated, nil, body)
}

func (server *ServiceHTTPServer) timedOut(ctx *gin.Context) {
	server.reply(ctx, ctx.JSON, http.StatusRequestTimeout, nil, nil)
}
