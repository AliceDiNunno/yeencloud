package interactor

import "github.com/AliceDiNunno/yeencloud/src/core/domain"

type Audit interface {
	// TODO: change data map keys to LogField for constitency
	NewTrace(trigger string, data map[string]string) domain.AuditTraceID
	AddStep(id domain.AuditTraceID, details ...interface{}) domain.AuditTraceStepID

	EndStep(id domain.AuditTraceID, step domain.AuditTraceStepID)
	EndTrace(id domain.AuditTraceID) domain.AuditTrace

	Log(id domain.AuditTraceID, step domain.AuditTraceStepID) LogMessage
	DumpTrace(id domain.AuditTraceID) *domain.AuditTrace
}
