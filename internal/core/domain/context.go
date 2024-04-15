package domain

type RequestCallback func(interface{}, error)
type RequestContext struct {
	TraceID AuditTraceID
	StepID  AuditTraceStepID

	Profile *Profile

	Done RequestCallback
}
