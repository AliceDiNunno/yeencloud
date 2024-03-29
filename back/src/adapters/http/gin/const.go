package gin

import "github.com/AliceDiNunno/yeencloud/src/core/domain"

const (
	CtxAuditField       = "audit"
	CtxLanguageField    = "lang"
	CtxSessionField     = "session"
	CtxUserField        = "user"
	CtxProfileField     = "profile"
	CtxHTTPCodeField    = "http_code"
	CtxTraceField       = "trace_dump"
	CtxProfileMailField = "mail"
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

const (
	HttpLogField          = domain.LogField("http")
	HttpLogFieldStatus    = HttpLogField + ".status"
	HttpLogFieldMethod    = HttpLogField + ".method"
	HttpLogFieldPath      = HttpLogField + ".path"
	HttpHandlerField      = HttpLogField + ".handler"
	HttpHandlerCountField = HttpLogField + ".handler_count"
)
