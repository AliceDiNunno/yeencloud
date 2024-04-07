package gin

import (
	"fmt"
	"strconv"

	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type requestTrace struct {
	UserAgent      string
	AcceptLanguage string
	IP             string
	Method         string
	Path           string
	Status         string
}

func (rt requestTrace) data() map[string]string {
	return map[string]string{
		"UserAgent":      rt.UserAgent,
		"AcceptLanguage": rt.AcceptLanguage,
		"IP":             rt.IP,
		"Method":         rt.Method,
		"Path":           rt.Path,
		"Status":         rt.Status,
	}
}

func (server *ServiceHTTPServer) traceHandlerMiddleware(ctx *gin.Context) {
	requestData := requestTrace{
		UserAgent:      ctx.GetHeader(HeaderUserAgent),
		AcceptLanguage: ctx.GetHeader(HeaderAcceptLanguage),
		IP:             ctx.ClientIP(),
		Method:         ctx.Request.Method,
		Path:           ctx.Request.URL.Path,
		Status:         strconv.Itoa(ctx.Writer.Status()),
	}

	trace := server.auditer.NewTrace(fmt.Sprintf("REST %s %s", ctx.Request.Method, ctx.Request.URL.Path),
		requestData.data())
	ctx.Set(CtxAuditField, trace)
	ctx.Next()
	defer func() {
		dump := server.auditer.EndTrace(trace)
		ctx.Set(CtxTraceField, dump)
	}()
}

func (server *ServiceHTTPServer) getTrace(ctx *gin.Context) domain.AuditTraceID {
	contextTrace, exists := ctx.Get(CtxAuditField)
	if !exists {
		return domain.AuditTraceID("")
	}

	trace, valid := contextTrace.(domain.AuditTraceID)

	if !valid {
		server.log.Log(domain.LogLevelWarn).Msg("Trace not found in context, returning empty ID")
		return domain.AuditTraceID(uuid.Nil.String())
	}

	return trace
}
