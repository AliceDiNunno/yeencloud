package domain

// MARK: - Objects

type LogLevel string

type LogField struct {
	Scope LogScope

	Identifier string
}
type LogFields map[LogField]interface{}

var (
	// MARK: - Config

	LogScopeConfig             = LogScope{Identifier: "config"}
	LogFieldConfigVersion      = LogField{Scope: LogScopeConfig, Identifier: "version"}
	LogFieldConfigDatabase     = LogField{Scope: LogScopeConfig, Identifier: "database"}
	LogFieldConfigHTTP         = LogField{Scope: LogScopeConfig, Identifier: "http"}
	LogFieldConfigRunContext   = LogField{Scope: LogScopeConfig, Identifier: "run_context"}
	LogFieldConfigLocalization = LogField{Scope: LogScopeConfig, Identifier: "localization"}
	LogFieldConfigMail         = LogField{Scope: LogScopeConfig, Identifier: "mail"}

	// MARK: - Time

	LogScopeTime        = LogScope{Identifier: "time"}
	LogFieldTimeStarted = LogField{Scope: LogScopeTime, Identifier: "started"}
	LogFieldTimeEnded   = LogField{Scope: LogScopeTime, Identifier: "ended"}
	LogFieldDuration    = LogField{Scope: LogScopeTime, Identifier: "duration"}
)

// MARK: - Log Levels

const (
	LogLevelDebug = LogLevel("DEBUG")
	LogLevelInfo  = LogLevel("INFO")
	LogLevelSQL   = LogLevel("SQL")
	LogLevelWarn  = LogLevel("WARN")
	LogLevelError = LogLevel("ERROR")
	LogLevelPanic = LogLevel("PANIC")
	LogLevelFatal = LogLevel("FATAL")
)

// MARK: - Functions

func (l LogLevel) String() string {
	return string(l)
}

func (l LogField) String() string {
	return l.Scope.String() + logSeparator + l.Identifier
}
