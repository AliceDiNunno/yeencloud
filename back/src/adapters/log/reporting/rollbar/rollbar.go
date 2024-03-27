package rollbar

import (
	"errors"
	"github.com/AliceDiNunno/yeencloud/src/adapters/log"
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/AliceDiNunno/yeencloud/src/core/domain/config"
	"github.com/rogpeppe/go-internal/modfile"
	"github.com/rollbar/rollbar-go"
	"io/ioutil"
	"strings"
)

type RollbarConfig struct {
	Token string
}

type RollbarMiddleware struct {
	config RollbarConfig
}

type RollbarLogLevel string

const (
	RollbarLogLevelDebug    RollbarLogLevel = "debug"
	RollbarLogLevelInfo     RollbarLogLevel = "info"
	RollbarLogLevelWarning  RollbarLogLevel = "warning"
	RollbarLogLevelError    RollbarLogLevel = "error"
	RollbarLogLevelCritical RollbarLogLevel = "critical"
)

func (z *RollbarMiddleware) translateLogLevel(level domain.LogLevel) RollbarLogLevel {
	switch level {
	case domain.LogLevelInfo:
		return RollbarLogLevelInfo
	case domain.LogLevelWarn:
		return RollbarLogLevelWarning
	case domain.LogLevelError:
		return RollbarLogLevelError
	case domain.LogLevelFatal, domain.LogLevelPanic:
		return RollbarLogLevelCritical
	default:
		return RollbarLogLevelDebug
	}
}

func (z *RollbarMiddleware) isLoggable(level RollbarLogLevel) bool {
	switch level {
	case RollbarLogLevelWarning, RollbarLogLevelError, RollbarLogLevelCritical:
		return true
	}

	return false
}

func (z *RollbarMiddleware) Log(message log.Message) {
	if z.config.Token == "" {
		return
	}
	/*
		trace, exists := message.Fields["trace"]
		if !exists {
			trace = nil
		}

		message.Fields["trace"] = nil
	*/
	if z.isLoggable(z.translateLogLevel(message.Level)) {
		rollbar.Log(message.Level.String(), errors.New(message.Message), map[string]interface{}(message.Fields))
	}
}

func GetModuleName() string {
	goModBytes, err := ioutil.ReadFile("go.mod")
	if err != nil {
		return "unknown"
	}

	modName := modfile.ModulePath(goModBytes)

	return modName
}

func NewRollbarMiddleware(config RollbarConfig, runContext config.RunContextConfig, version config.VersionConfig) *RollbarMiddleware {
	configMw := &RollbarMiddleware{
		config: config,
	}

	rollbar.SetToken(config.Token)

	rollbar.SetEnvironment(runContext.Environment)
	codeversion := version.SHA
	if codeversion == "" {
		codeversion = "main"
	}
	rollbar.SetCodeVersion(codeversion)        // optional Git hash/branch/tag (required for GitHub integration)
	rollbar.SetServerHost(runContext.Hostname) // optional override; defaults to hostname
	s, _ := strings.CutSuffix(runContext.WorkingDirectory, "/back")
	rollbar.SetServerRoot(s) // path of project (required for GitHub integration and non-project stacktrace collapsing)

	return configMw
}
