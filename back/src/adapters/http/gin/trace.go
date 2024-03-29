package gin

import (
	"back/src/core/domain"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (server *ServiceHTTPServer) traceHandlerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trace := server.auditer.NewTrace(fmt.Sprintf("REST %s %s", ctx.Request.Method, ctx.Request.URL.Path),
			ctx.Request.Header, ctx.ClientIP(), ctx.GetHeader("User-Agent"))
		ctx.Set("audit", trace)
		ctx.Next()
		defer func() {
			server.auditer.EndTrace(trace)
		}()
	}
}

func (server *ServiceHTTPServer) getTrace(ctx *gin.Context) domain.AuditID {
	contextTrace, exists := ctx.Get("audit")
	if !exists {
		return domain.AuditID("")
	}

	return contextTrace.(domain.AuditID)
}
