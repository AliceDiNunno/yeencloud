package audit

import (
	"encoding/json"
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/AliceDiNunno/yeencloud/src/core/interactor"
	"github.com/google/uuid"

	"time"
)

// #YC-23 TODO: add functions to step and returning it on addStep
// #YC-24 TODO: add end step function to measure time between steps
// #YC-25 TODO: add more data to each step if necessary (like error, or inter steps)

type Audit struct {
	SaveTrace domain.AuditSaveFunc
	Logger    interactor.Logger

	currentTraces map[domain.AuditID]*domain.Request
}

func (a *Audit) NewTrace(trigger string, triggerdata map[string]string) domain.AuditID {
	trace := domain.Request{
		ID:          domain.AuditID(uuid.New().String()),
		Trigger:     trigger,
		Content:     nil,
		TriggerData: triggerdata,
		Frame:       a.getFrame(),
	}

	a.Log(trace.ID, "").WithField("trigger", trigger).WithField("start", time.Now()).Msg("New trace started")

	a.currentTraces[trace.ID] = &trace
	// Setting time at the end to avoid time drift at best as possible
	trace.StartedAt = time.Now().UnixMilli()
	return trace.ID
}

func (a *Audit) EndTrace(id domain.AuditID) domain.Request {
	trace, exists := a.currentTraces[id]
	trace.EndedAt = time.Now().UnixMilli()

	duration := time.Duration(trace.EndedAt - trace.StartedAt)
	if !exists {
		a.Log(trace.ID, "").WithLevel(domain.LogLevelWarn).Msg("Trace not found, aborting EndTrace")
		return domain.Request{}
	}

	lastStep := &trace.Content[len(trace.Content)-1]
	if lastStep.End.IsZero() {
		lastStep.End = time.Now()
	}

	a.Log(trace.ID, "").WithFields(domain.LogFields{
		"duration": duration,
		"steps":    len(trace.Content),
	}).Msg("Trace ended")

	a.saveTrace(*trace)

	// Cleanup
	delete(a.currentTraces, id)
	return *trace
}

func (a *Audit) saveTrace(trace domain.Request) {
	json, err := json.Marshal(trace)

	if err != nil {
		a.Log(trace.ID, "").WithLevel(domain.LogLevelWarn).WithFields(domain.LogFields{
			"error": err,
		}).Msg("Error marshalling trace")

		return
	}

	if a.SaveTrace != nil {
		a.SaveTrace(json)
	} else {
		a.Log(trace.ID, "").WithLevel(domain.LogLevelWarn).WithFields(domain.LogFields{
			"error": err,
		}).Msg("No save trace function defined, trace will be discarded.")
	}
}

func NewAuditer(logger interactor.Logger, saveTrace domain.AuditSaveFunc) *Audit {
	return &Audit{
		currentTraces: make(map[domain.AuditID]*domain.Request),
		SaveTrace:     saveTrace,
		Logger:        logger,
	}
}
