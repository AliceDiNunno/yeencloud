package audit

import (
	"time"

	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/AliceDiNunno/yeencloud/src/core/interactor"
	"github.com/google/uuid"
)

var DefaultSkip = 2

func (a *Audit) AddStep(id domain.AuditTraceID, skip int, details ...interface{}) domain.AuditTraceStepID {
	trace, exists := a.currentTraces[id]

	if !exists {
		a.Logger.Log(domain.LogLevelInfo).WithField(domain.LogFieldTraceID, id).Msg("Trace not found, aborting add step")
		return NoStep
	}

	currentStep := domain.AuditTraceStep{
		ID:      domain.AuditTraceStepID(uuid.New().String()),
		Caller:  a.getFrame(skip),
		Details: []interface{}{},
	}

	if len(details) > 0 {
		currentStep.Details = append(currentStep.Details, details...)
	}

	if len(trace.Content) > 0 && trace.Content[len(trace.Content)-1].End.IsZero() {
		trace.Content[len(trace.Content)-1].End = time.Now()
	}

	a.Logger.Log(domain.LogLevelInfo).WithField(domain.LogFieldTraceID, id).WithField(domain.LogFieldStepID, currentStep.ID).Msg("Step added")
	currentStep.Start = time.Now()
	trace.Content = append(trace.Content, currentStep)
	return currentStep.ID
}

func (audit *Audit) findStep(auditID domain.AuditTraceID, stepID domain.AuditTraceStepID) *domain.AuditTraceStep {
	trace, exists := audit.currentTraces[auditID]
	if !exists {
		return nil
	}

	if trace.Content == nil {
		return nil
	}

	for _, step := range trace.Content {
		if step.ID == stepID {
			return &step
		}
	}

	return nil
}

func (audit *Audit) Log(auditID domain.AuditTraceID, stepID domain.AuditTraceStepID) interactor.LogMessage {
	return audit.Logger.Log(domain.LogLevelInfo).WithField(domain.LogFieldTraceID, auditID).WithField(domain.LogFieldStepID, stepID)
}

func (audit *Audit) EndStep(auditID domain.AuditTraceID, stepID domain.AuditTraceStepID) {
	endTime := time.Now()
	step := audit.findStep(auditID, stepID)

	if step != nil {
		step.End = endTime
	}
	audit.Logger.Log(domain.LogLevelInfo).WithField(domain.LogFieldTraceID, auditID).WithField(domain.LogFieldStepID, stepID).Msg("Step ended")
}
