package rollbar

import "github.com/AliceDiNunno/yeencloud/src/core/domain"

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
