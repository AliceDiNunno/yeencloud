package gin

import (
	"github.com/AliceDiNunno/yeencloud/src/core/usecases"
	"github.com/gin-gonic/gin"
)

type RequestContext struct {
	TraceID  string
	Usecases usecases.Usecases
}

func (server *ServiceHTTPServer) RequestContextMiddleware(ctx *gin.Context) {
	context := &RequestContext{
		TraceID:  server.getTrace(ctx).String(),
		Usecases: server.ucs.StartRequest(),
	}

	ctx.Set(CtxRequestContextField, context)
}

func (server *ServiceHTTPServer) getContext(ctx *gin.Context) *RequestContext {
	contextField := ctx.MustGet(CtxRequestContextField)

	if contextField == nil {
		println("No context found")
		return nil
	}

	context, ok := contextField.(*RequestContext)

	if !ok {
		println("Context is not of the right type")
		return nil
	}

	return context
}

func (server *ServiceHTTPServer) usecases(ctx *gin.Context) usecases.Usecases {
	return server.getContext(ctx).Usecases
}
