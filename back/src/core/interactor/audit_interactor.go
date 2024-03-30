package interactor

import "github.com/AliceDiNunno/yeencloud/src/core/domain"

type Audit interface {
	// TODO: change data map keys to LogField for constitency
	NewTrace(trigger string, data map[string]string) domain.AuditID
	AddStep(id domain.AuditID, details ...interface{}) domain.StepID

	EndStep(id domain.AuditID, step domain.StepID)
	EndTrace(id domain.AuditID) domain.Request

	Log(id domain.AuditID, step domain.StepID) LogMessage
	DumpTrace(id domain.AuditID) *domain.Request
}
