package audit

import (
	"back/src/core/domain"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"runtime"
	"strings"
	"time"
)

// #YC-23 TODO: add functions to step and returning it on addStep
// #YC-24 TODO: add end step function to measure time between steps
// #YC-25 TODO: add more data to each step if necessary (like error, or inter steps)

type Audit struct {
	SaveTrace domain.AuditSaveFunc

	currentTraces map[domain.AuditID]*root
}

type step struct {
	Next    *step
	Caller  map[string]interface{}
	Details []interface{}
}

type root struct {
	ID        domain.AuditID
	StartedAt int64
	EndedAt   int64

	Trigger     string
	TriggerData []interface{}

	Result []interface{}

	Content *step
}

func (a *Audit) NewTrace(trigger string, triggerdata ...interface{}) domain.AuditID {
	trace := root{
		ID:      domain.AuditID(uuid.New().String()),
		Trigger: trigger,
		Content: nil,
	}

	if len(triggerdata) > 0 {
		for _, data := range triggerdata {
			trace.TriggerData = append(trace.TriggerData, data)
		}
	}

	log.Info().Str("audit", trace.ID.String()).Time("start", time.Now()).Msg("Trace started")
	a.currentTraces[trace.ID] = &trace
	//setting time at the end to avoid time drift at best as possible
	trace.StartedAt = time.Now().UnixNano()
	return trace.ID
}

func (a *Audit) AddStep(id domain.AuditID, details ...interface{}) {
	trace, exists := a.currentTraces[id]

	if !exists {
		log.Info().Str("audit", id.String()).Msg("Trace not found, aborting add step")
		return
	}

	step := step{
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
		caller := map[string]interface{}{
			"file":     file,
			"line":     line,
			"function": functionName,
		}
		step.Caller = caller
	}

	if len(details) > 0 {
		for _, data := range details {
			step.Details = append(step.Details, data)
		}
	}

	if trace.Content == nil {
		trace.Content = &step
	} else {
		if trace.Content.Next == nil {
			trace.Content.Next = &step
		} else {
			last := trace.Content
			for last.Next != nil {
				last = last.Next
			}
			last.Next = &step
		}
	}
}

func (a *Audit) EndTrace(id domain.AuditID, result ...interface{}) {
	trace, exists := a.currentTraces[id]
	trace.EndedAt = time.Now().UnixNano()

	duration := time.Duration(trace.EndedAt-trace.StartedAt) * time.Millisecond
	if !exists {
		log.Info().Str("audit", id.String()).Msg("Trace not found, aborting end audit")
		return
	}

	if len(result) > 0 {
		for _, data := range result {
			trace.TriggerData = append(trace.TriggerData, data)
		}
	}

	steps := 0

	if trace.Content != nil {
		steps++
		for step := trace.Content; step.Next != nil; step = step.Next {
			steps++
		}
	}

	log.Info().Str("audit", id.String()).Dur("duration", duration/1000000).Int("steps", steps).Msg("Trace ended")

	a.saveTrace(*trace)

	//cleanup
	delete(a.currentTraces, id)
}

func (a *Audit) saveTrace(trace root) {
	j, _ := json.Marshal(trace)
	if a.SaveTrace != nil {
		a.SaveTrace(j)
	} else {
		log.Warn().Str("audit", trace.ID.String()).Msg("No save trace function defined, trace will be discarded.")
	}
}

func NewAuditer(saveTrace domain.AuditSaveFunc) *Audit {
	return &Audit{
		currentTraces: make(map[domain.AuditID]*root),
		SaveTrace:     saveTrace,
	}
}
