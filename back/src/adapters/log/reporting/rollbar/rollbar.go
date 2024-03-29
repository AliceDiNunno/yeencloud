package rollbar

import (
	"errors"
	"fmt"
	"github.com/AliceDiNunno/yeencloud/src/adapters/log"
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/AliceDiNunno/yeencloud/src/core/domain/config"
	"github.com/davecgh/go-spew/spew"
	"github.com/rollbar/rollbar-go"
	"runtime"
)

type Config struct {
	Token string
}

type Middleware struct {
	config Config
}

type LogLevel string

const (
	LogLevelDebug    LogLevel = "debug"
	LogLevelInfo     LogLevel = "info"
	LogLevelWarning  LogLevel = "warning"
	LogLevelError    LogLevel = "error"
	LogLevelCritical LogLevel = "critical"
)

func (r *Middleware) translateLogLevel(level domain.LogLevel) LogLevel {
	switch level {
	case domain.LogLevelInfo:
		return LogLevelInfo
	case domain.LogLevelWarn:
		return LogLevelWarning
	case domain.LogLevelError:
		return LogLevelError
	case domain.LogLevelFatal, domain.LogLevelPanic:
		return LogLevelCritical
	default:
		return LogLevelDebug
	}
}

func (r *Middleware) isLoggable(level LogLevel) bool {
	return level == LogLevelWarning ||
		level == LogLevelError ||
		level == LogLevelCritical
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

	dump, present := message.Fields["trace.dump"]

	if present {
		trace, ok := dump.(domain.Request)

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

	spew.Dump(r.Traceback(message))

	profileID, profilePresent := message.Fields["profile.id"]
	profileStr, ok := profileID.(domain.ProfileID)
	if profilePresent && ok {
		email, mailFound := message.Fields["profile.email"]
		name, nameFound := message.Fields["profile.name"]

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
		if key == "trace.dump" {
			trace, traceOk := v.(domain.Request)

			if traceOk {
				for triggerKey, triggerValue := range trace.TriggerData {
					additionnalFields[triggerKey] = triggerValue
				}

				//TODO: Add trace data to callstack (it is not traced otherwise and is relevant)
				additionnalFields["trace.trigger"] = trace.Trigger
				additionnalFields["trace.id"] = trace.ID.String()
				additionnalFields["trace.start"] = trace.StartedAt
				additionnalFields["trace.end"] = trace.EndedAt
				additionnalFields["trace.steps"] = len(trace.Content)
				additionnalFields["trace.result"] = trace.Result

				for i, step := range trace.Content {
					istr := fmt.Sprintf("trace.%d", i)
					additionnalFields[istr+".caller"] = step.Caller
					additionnalFields[istr+".details"] = step.Details
					additionnalFields[istr+".start"] = step.Start
					additionnalFields[istr+".end"] = step.End
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
