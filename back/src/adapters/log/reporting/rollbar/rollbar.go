package rollbar

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/AliceDiNunno/yeencloud/src/adapters/log"
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/AliceDiNunno/yeencloud/src/core/domain/config"
	"github.com/rollbar/rollbar-go"
)

type Config struct {
	Token string
}

type Middleware struct {
	config Config
}

func (r *Middleware) Traceback(message log.Message) []runtime.Frame {
	var frames []runtime.Frame

	for skip := 0; ; skip++ {
		pc, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}

		callerfunc := runtime.FuncForPC(pc)

		frames = append(frames, runtime.Frame{
			PC:       pc,
			File:     file,
			Line:     line,
			Function: callerfunc.Name(),
		})
	}

	dump, present := message.Fields[domain.LogFieldTraceDump]

	if present {
		trace, ok := dump.(domain.AuditTrace)

		if ok {
			for _, step := range trace.Content {
				frames = append(frames, step.Caller)
			}
		}
	}

	return frames
}

func (r *Middleware) Log(message log.Message) {
	if r.config.Token == "" {
		return
	}

	if !r.isLoggable(r.translateLogLevel(message.Level)) {
		return
	}

	rollbar.SetStackTracer(func(error) ([]runtime.Frame, bool) {
		return r.Traceback(message), true
	})

	profileID, profilePresent := message.Fields[domain.LogFieldProfileID]
	profileStr, ok := profileID.(domain.ProfileID)
	if profilePresent && ok {
		email, mailFound := message.Fields[domain.LogFieldProfileMail]
		name, nameFound := message.Fields[domain.LogFieldProfileName]

		if mailFound && nameFound {
			emailStr, emailOk := email.(string)
			nameStr, nameOk := name.(string)

			if emailOk && nameOk {
				rollbar.SetPerson(profileStr.String(), nameStr, emailStr)
			}
		}
	}

	additionnalFields := make(map[string]interface{})

	for key, v := range message.Fields {
		if key == domain.LogFieldTraceDump {
			trace, traceOk := v.(domain.AuditTrace)

			if traceOk {
				for triggerKey, triggerValue := range trace.TriggerData {
					additionnalFields[triggerKey] = triggerValue
				}

				additionnalFields[domain.LogFieldTraceTrigger.String()] = trace.Trigger
				additionnalFields[domain.LogFieldTraceID.String()] = trace.ID.String()
				additionnalFields[domain.LogFieldTimeStarted.String()] = trace.StartedAt
				additionnalFields[domain.LogFieldTimeEnded.String()] = trace.EndedAt
				additionnalFields[domain.LogFieldTraceStepCount.String()] = len(trace.Content)
				additionnalFields[domain.LogFieldTraceResult.String()] = trace.Result

				for i, step := range trace.Content {
					additionnalFields[fmt.Sprintf(domain.LogFieldTraceStepCaller.String(), i)] = step.Caller
					additionnalFields[fmt.Sprintf(domain.LogFieldTraceStepDetails.String(), i)] = step.Details
					additionnalFields[fmt.Sprintf(domain.LogFieldTraceStepStart.String(), i)] = step.Start
					additionnalFields[fmt.Sprintf(domain.LogFieldTraceStepEnd.String(), i)] = step.End
				}
			}
		} else {
			additionnalFields[key.String()] = v
		}
	}

	rollbar.Log(message.Level.String(), errors.New(message.Message), additionnalFields)
}

func NewRollbarMiddleware(config Config, runContext config.RunContextConfig, version config.VersionConfig) *Middleware {
	configMw := &Middleware{
		config: config,
	}

	rollbar.SetToken(config.Token)

	rollbar.SetEnvironment(runContext.Environment)
	codeversion := version.SHA
	if codeversion == "" {
		codeversion = "main"
	}

	rollbar.SetCodeVersion(codeversion)
	rollbar.SetServerHost(runContext.Hostname)
	rollbar.SetServerRoot(runContext.WorkingDirectory)

	return configMw
}
