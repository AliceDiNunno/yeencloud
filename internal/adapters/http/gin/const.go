package gin

import "github.com/AliceDiNunno/yeencloud/internal/core/domain"

const (
	CtxAuditField          = "audit"
	CtxLanguageField       = "lang"
	CtxSessionField        = "session"
	CtxUserField           = "user"
	CtxProfileField        = "profile"
	CtxHTTPCodeField       = "http_code"
	CtxTraceField          = "trace_dump"
	CtxProfileMailField    = "mail"
	CtxOrganizationField   = "organization"
	CtxRequestContextField = "request_context"
)

const (
	HeaderAcceptLanguage = "Accept-Language"
	HeaderAuthorization  = "Authorization"
	HeaderUserAgent      = "User-Agent"
	HeaderContentType    = "Content-Type"
	HeaderContentLength  = "Content-Length"
)

const (
	MethodGet    = "GET"
	MethodPost   = "POST"
	MethodPut    = "PUT"
	MethodPatch  = "PATCH"
	MethodDelete = "DELETE"
	MethodOption = "OPTION"
)

var (
	LogScopeHttp             = domain.LogScope{Identifier: "http"}
	LogFieldHttpStatus       = domain.LogField{Scope: LogScopeHttp, Identifier: "status"}
	LogFieldHttpMethod       = domain.LogField{Scope: LogScopeHttp, Identifier: "method"}
	LogFieldHttpPath         = domain.LogField{Scope: LogScopeHttp, Identifier: "path"}
	LogFieldHttpHandler      = domain.LogField{Scope: LogScopeHttp, Identifier: "handler"}
	LogFieldHttpHandlerCount = domain.LogField{Scope: LogScopeHttp, Identifier: "handler_count"}
)
