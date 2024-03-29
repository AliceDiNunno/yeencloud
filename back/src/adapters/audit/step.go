package audit

import (
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/AliceDiNunno/yeencloud/src/core/interactor"
	"github.com/google/uuid"
	"runtime"
	"strings"
	"time"
)

func (a *Audit) AddStep(id domain.AuditID, details ...interface{}) domain.StepID {
	trace, exists := a.currentTraces[id]

	if !exists {
		a.Logger.Log(domain.LogLevelInfo).WithField("audit", id.String()).Msg("Trace not found, aborting add step")
		return domain.StepID(uuid.Nil.String())
	}

	currentStep := domain.Step{
		ID:      domain.StepID(uuid.New().String()),
		Next:    nil,
		Caller:  map[string]interface{}{},
		Details: []interface{}{},
	}

	pc, file, line, ok := runtime.Caller(1)

	functionName := "<unknown>"
	fn := runtime.FuncForPC(pc)
	if fn != nil {
		fname := strings.Split(fn.Name(), "/")
		functionName = fname[len(fname)-1]
	}

	if ok {
		// TODO: create a struct instead of a map
		caller := map[string]interface{}{
			"file":     file,
			"line":     line,
			"function": functionName,
		}
		currentStep.Caller = caller
	}

	if len(details) > 0 {
		currentStep.Details = append(currentStep.Details, details...)
	}

	if trace.Content == nil {
		trace.Content = &currentStep
	} else {
		if trace.Content.Next == nil {
			trace.Content.Next = &currentStep
		} else {
			last := trace.Content
			for last.Next != nil {
				last = last.Next
			}
			last.Next = &currentStep
		}
	}

	a.Logger.Log(domain.LogLevelInfo).WithField("audit", id.String()).WithField("step", currentStep.ID.String()).Msg("Step added")
	currentStep.Start = time.Now()
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

	if trace.Content.ID == stepID {
		return trace.Content
	}

	current := trace.Content
	for current.Next != nil {
		if current.Next.ID == stepID {
			return current.Next
		}
		current = current.Next
	}

	return nil
}

func (audit *Audit) Log(auditID domain.AuditID, stepID domain.StepID) interactor.LogMessage {
	return audit.Logger.Log(domain.LogLevelInfo)
}

func (audit *Audit) EndStep(auditID domain.AuditID, stepID domain.StepID) {
	step := audit.findStep(auditID, stepID)

	if step != nil {
		step.End = time.Now()
	}
	audit.Logger.Log(domain.LogLevelInfo).WithField("audit", auditID).WithField("step", stepID).Msg("Step ended")

}
