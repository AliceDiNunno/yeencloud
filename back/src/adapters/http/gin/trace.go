package gin

import (
	"fmt"
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (server *ServiceHTTPServer) traceHandlerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trace := server.auditer.NewTrace(fmt.Sprintf("REST %s %s", ctx.Request.Method, ctx.Request.URL.Path),
			ctx.Request.Header, ctx.ClientIP(), ctx.GetHeader(HeaderUserAgent))
		ctx.Set(CtxAuditField, trace)
		ctx.Next()
		defer func() {
			dump := server.auditer.EndTrace(trace)
			ctx.Set("trace_dump", dump)
		}()
	}
}

func (server *ServiceHTTPServer) getTrace(ctx *gin.Context) domain.AuditID {
	contextTrace, exists := ctx.Get(CtxAuditField)
	if !exists {
		return domain.AuditID("")
	}

	trace, err := contextTrace.(domain.AuditID)

	if !err {
		server.log.Log(domain.LogLevelWarn).Msg("Trace not found in context, returning empty ID")
		return domain.AuditID(uuid.Nil.String())
	}

	return trace
}
