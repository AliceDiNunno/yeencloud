package gin

import "github.com/AliceDiNunno/yeencloud/src/core/domain"

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
	HttpLogField          = domain.LogField{Name: "http"}
	HttpLogFieldStatus    = domain.LogField{Parent: &HttpLogField, Name: "status"}
	HttpLogFieldMethod    = domain.LogField{Parent: &HttpLogField, Name: "method"}
	HttpLogFieldPath      = domain.LogField{Parent: &HttpLogField, Name: "path"}
	HttpHandlerField      = domain.LogField{Parent: &HttpLogField, Name: "handler"}
	HttpHandlerCountField = domain.LogField{Parent: &HttpLogField, Name: "handler_count"}
)
