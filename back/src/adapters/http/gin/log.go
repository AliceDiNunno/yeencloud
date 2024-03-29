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

		errorCode := ctx.GetString("http_code")
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

		fields := domain.LogFields{
			"http.status": status,
			"http.method": ctx.Request.Method,
			"http.path":   ctx.Request.URL.Path,
			"trace.id":    server.getTrace(ctx).String(),
		}

		tracedump, valid := trace.(domain.Request)
		if valid {
			fields["trace.dump"] = tracedump
		}

		profile, profileExists := ctx.Get(CtxProfileField)
		if profileExists {
			pprofile, ok := profile.(domain.Profile)

			if ok {
				fields["profile.id"] = pprofile.ID
				fields["profile.name"] = pprofile.Name
				fields["profile.email"] = ctx.GetString("mail")
			}
		}

		server.log.Log(level).WithFields(fields).Msg(message)
	}
}
