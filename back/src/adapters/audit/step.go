package audit

import (
	"back/src/core/domain"
	"github.com/rs/zerolog/log"
	"runtime"
	"strings"
)

type step struct {
	Next    *step
	Caller  map[string]interface{}
	Details []interface{}
}

func (a *Audit) AddStep(id domain.AuditID, details ...interface{}) {
	trace, exists := a.currentTraces[id]

	if !exists {
		log.Info().Str(domain.LogFieldAudit, id.String()).Msg("Trace not found, aborting add step")
		return
	}

	currentStep := step{
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
}
