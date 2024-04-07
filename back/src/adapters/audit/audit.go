package audit

import (
	"encoding/json"

	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/AliceDiNunno/yeencloud/src/core/interactor"
)

// #YC-23 TODO: add functions to step and returning it on addStep
// #YC-24 TODO: add end step function to measure time between steps
// #YC-25 TODO: add more data to each step if necessary (like error, or inter steps)

var NoStep domain.AuditTraceStepID = domain.AuditTraceStepID("NoStep")

type Audit struct {
	SaveTrace domain.AuditSaveFunc
	Logger    interactor.Logger

	currentTraces map[domain.AuditTraceID]*domain.AuditTrace
}

func (a *Audit) save(trace domain.AuditTrace) {
	json, err := json.Marshal(trace)

	if err != nil {
		a.Log(trace.ID, NoStep).WithLevel(domain.LogLevelWarn).WithField(domain.LogFieldError, err).Msg("Error marshalling trace")

		return
	}

	if a.SaveTrace != nil {
		a.SaveTrace(json)
	} else {
		a.Log(trace.ID, NoStep).WithLevel(domain.LogLevelWarn).WithField(domain.LogFieldError, err).Msg("No save trace function defined, trace will be discarded.")
	}
}

func NewAuditer(logger interactor.Logger, saveTrace domain.AuditSaveFunc) *Audit {
	return &Audit{
		currentTraces: make(map[domain.AuditTraceID]*domain.AuditTrace),
		SaveTrace:     saveTrace,
		Logger:        logger,
	}
}
