package domain

type RequestCallback func(interface{}, *ErrorDescription)
type RequestContext struct {
	TraceID AuditTraceID
	StepID  AuditTraceStepID

	Profile *Profile

	Done RequestCallback
}
