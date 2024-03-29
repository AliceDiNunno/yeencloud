package audit

import (
	"back/src/core/domain"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"time"
)

// #YC-23 TODO: add functions to step and returning it on addStep
// #YC-24 TODO: add end step function to measure time between steps
// #YC-25 TODO: add more data to each step if necessary (like error, or inter steps)

type Audit struct {
	SaveTrace domain.AuditSaveFunc

	currentTraces map[domain.AuditID]*root
}

func (a *Audit) NewTrace(trigger string, triggerdata ...interface{}) domain.AuditID {
	trace := root{
		ID:      domain.AuditID(uuid.New().String()),
		Trigger: trigger,
		Content: nil,
	}

	if len(triggerdata) > 0 {
		trace.TriggerData = append(trace.TriggerData, triggerdata...)
	}

	log.Info().Str(domain.LogFieldAudit, trace.ID.String()).Time("start", time.Now()).Msg("Trace started")
	a.currentTraces[trace.ID] = &trace
	// Setting time at the end to avoid time drift at best as possible
	trace.StartedAt = time.Now().UnixNano()
	return trace.ID
}

func (a *Audit) EndTrace(id domain.AuditID, result ...interface{}) {
	trace, exists := a.currentTraces[id]
	trace.EndedAt = time.Now().UnixNano()

	duration := time.Duration(trace.EndedAt-trace.StartedAt) * time.Millisecond
	if !exists {
		log.Info().Str(domain.LogFieldAudit, id.String()).Msg("Trace not found, aborting end audit")
		return
	}

	if len(result) > 0 {
		trace.TriggerData = append(trace.TriggerData, result...)
	}

	steps := 0

	if trace.Content != nil {
		steps++
		for step := trace.Content; step.Next != nil; step = step.Next {
			steps++
		}
	}

	log.Info().Str(domain.LogFieldAudit, id.String()).Dur("duration", duration/1000000).Int("steps", steps).Msg("Trace ended")

	a.saveTrace(*trace)

	// Cleanup
	delete(a.currentTraces, id)
}

func (a *Audit) saveTrace(trace root) {
	json, err := json.Marshal(trace)

	if err != nil {
		log.Warn().Str(domain.LogFieldAudit, trace.ID.String()).Err(err).Msg("Error marshalling trace")
		return
	}

	if a.SaveTrace != nil {
		a.SaveTrace(json)
	} else {
		log.Warn().Str(domain.LogFieldAudit, trace.ID.String()).Msg("No save trace function defined, trace will be discarded.")
	}
}

func NewAuditer(saveTrace domain.AuditSaveFunc) *Audit {
	return &Audit{
		currentTraces: make(map[domain.AuditID]*root),
		SaveTrace:     saveTrace,
	}
}
