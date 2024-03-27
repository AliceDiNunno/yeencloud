package gin

import (
	"fmt"
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/gin-gonic/gin"
)

func (server *ServiceHTTPServer) ginLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		status := ctx.Writer.Status()

		message := fmt.Sprintf("%s %s %d", ctx.Request.Method, ctx.Request.URL.Path, status)

		errorCode := ctx.GetString("error_code")
		trace, _ := ctx.Get("trace_dump")

		if errorCode != "" {
			message += " - " + errorCode
		}

		level := domain.LogLevelInfo

		if status >= 400 && status < 500 {
			level = domain.LogLevelWarn
		} else if status >= 500 {
			level = domain.LogLevelError
		}

		server.log.Log(level).WithFields(map[string]interface{}{
			"status":  status,
			"method":  ctx.Request.Method,
			"path":    ctx.Request.URL.Path,
			"traceID": server.getTrace(ctx).String(),
			"trace":   trace,
		}).Msg(message)
	}
}

func (server *ServiceHTTPServer) printRoutes(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	server.log.Log(domain.LogLevelInfo).WithFields(map[string]interface{}{
		"handler":  handlerName,
		"handlers": nuHandlers,
	}).Msg(httpMethod + " " + absolutePath)
}
