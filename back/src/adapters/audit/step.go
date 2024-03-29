package audit

import (
	"time"

	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/AliceDiNunno/yeencloud/src/core/interactor"
	"github.com/google/uuid"
)

func (a *Audit) AddStep(id domain.AuditID, details ...interface{}) domain.StepID {
	trace, exists := a.currentTraces[id]

	if !exists {
		a.Logger.Log(domain.LogLevelInfo).WithField(domain.LogFieldTraceID, id.String()).Msg("Trace not found, aborting add step")
		return NoStep
	}

	currentStep := domain.Step{
		ID:      domain.StepID(uuid.New().String()),
		Caller:  a.getFrame(),
		Details: []interface{}{},
	}

	if len(details) > 0 {
		currentStep.Details = append(currentStep.Details, details...)
	}

	if len(trace.Content) > 0 && trace.Content[len(trace.Content)-1].End.IsZero() {
		trace.Content[len(trace.Content)-1].End = time.Now()
	}

	a.Logger.Log(domain.LogLevelInfo).WithField(domain.LogFieldTraceID, id.String()).WithField(domain.LogFieldStepID, currentStep.ID.String()).Msg("Step added")
	currentStep.Start = time.Now()
	trace.Content = append(trace.Content, currentStep)
	return currentStep.ID
}

func (audit *Audit) findStep(auditID domain.AuditID, stepID domain.StepID) *domain.Step {
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

func (audit *Audit) Log(auditID domain.AuditID, stepID domain.StepID) interactor.LogMessage {
	return audit.Logger.Log(domain.LogLevelInfo)
}

func (audit *Audit) EndStep(auditID domain.AuditID, stepID domain.StepID) {
	endTime := time.Now()
	step := audit.findStep(auditID, stepID)

	if step != nil {
		step.End = endTime
	}
	audit.Logger.Log(domain.LogLevelInfo).WithField(domain.LogFieldTraceID, auditID).WithField(domain.LogFieldStepID, stepID).Msg("Step ended")
}
