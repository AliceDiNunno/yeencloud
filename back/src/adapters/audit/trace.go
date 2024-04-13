package audit

import (
	"time"

	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/google/uuid"
)

func (a *Audit) NewTrace(trigger string, data map[string]string) domain.AuditTraceID {
	trace := domain.AuditTrace{
		ID:          domain.AuditTraceID(uuid.New().String()),
		Trigger:     trigger,
		Content:     nil,
		TriggerData: data,
		Frame:       a.getFrame(2),
	}

	a.Log(trace.ID, NoStep).
		WithField(domain.LogFieldTraceTrigger, trigger).
		WithField(domain.LogFieldTimeStarted, time.Now()).
		Msg("New trace started")

	a.currentTraces[trace.ID] = &trace
	// Setting time at the end to avoid time drift at best as possible
	trace.StartedAt = time.Now().UnixMilli()
	return trace.ID
}

func (a *Audit) EndTrace(id domain.AuditTraceID) domain.AuditTrace {
	trace, exists := a.currentTraces[id]
	trace.EndedAt = time.Now().UnixMilli()

	duration := time.Duration(trace.EndedAt - trace.StartedAt)
	if !exists {
		a.Log(trace.ID, NoStep).WithLevel(domain.LogLevelWarn).Msg("Trace not found, aborting EndTrace")
		return domain.AuditTrace{}
	}

	lastStep := &trace.Content[len(trace.Content)-1]
	if lastStep.End.IsZero() {
		lastStep.End = time.Now()
	}

	a.Log(trace.ID, NoStep).WithFields(domain.LogFields{
		domain.LogFieldDuration:       duration,
		domain.LogFieldTraceStepCount: len(trace.Content),
	}).Msg("Trace ended")

	a.save(*trace)

	// Cleanup
	delete(a.currentTraces, id)
	return *trace
}
