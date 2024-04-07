package gin

import (
	"fmt"

	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/gin-gonic/gin"
)

func (server *ServiceHTTPServer) ginLogger(ctx *gin.Context) {
	ctx.Next()

	status := ctx.Writer.Status()
	httpCode := ctx.GetString(CtxHTTPCodeField)
	trace, _ := ctx.Get(CtxTraceField)

	message := fmt.Sprintf("%s %s %d", ctx.Request.Method, ctx.Request.URL.Path, status)
	if httpCode != "" {
		message += " - " + httpCode
	}

	level := domain.LogLevelInfo
	if status >= 400 && status < 500 {
		level = domain.LogLevelWarn
	} else if status >= 500 {
		level = domain.LogLevelError
	}

	fields := domain.LogFields{
		LogFieldHttpStatus:     status,
		LogFieldHttpMethod:     ctx.Request.Method,
		LogFieldHttpPath:       ctx.Request.URL.Path,
		domain.LogFieldTraceID: server.getTrace(ctx).String(),
	}

	tracedump, valid := trace.(domain.AuditTrace)
	if valid {
		fields[domain.LogFieldTraceDump] = tracedump
	}
	profile, profileExists := ctx.Get(CtxProfileField)
	if profileExists {
		pprofile, ok := profile.(domain.Profile)

		if ok {
			fields[domain.LogFieldTraceID] = pprofile.ID
			fields[domain.LogFieldProfileName] = pprofile.Name
			fields[domain.LogFieldProfileMail] = ctx.GetString("mail")
		}
	}

	server.log.Log(level).WithFields(fields).Msg(message)
}
