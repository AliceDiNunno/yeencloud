package rollbar

import (
	"errors"
	"fmt"
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

func (r *Middleware) fieldToRollbarField(fields domain.LogFields) map[string]interface{} {
	rollbarFields := make(map[string]interface{})

	for key, value := range fields {
		rollbarFields[key.String()] = value
	}

	return rollbarFields
}

func (r *Middleware) Log(message log.Message) {
	if r.config.Token == "" {
		return
	}

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

	if r.isLoggable(r.translateLogLevel(message.Level)) {
		additionnalFields := make(map[string]interface{})

		for key, v := range message.Fields {
			if key == "trace.dump" {
				trace, traceOk := v.(domain.Request)

				if traceOk {
					for triggerKey, triggerValue := range trace.TriggerData {
						additionnalFields[triggerKey] = triggerValue
					}

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
			}
		}

		rollbar.Log(message.Level.String(), errors.New(message.Message), r.fieldToRollbarField(message.Fields), additionnalFields)
	}
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
