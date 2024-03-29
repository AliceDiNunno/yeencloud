package env

import (
	configDomain "github.com/AliceDiNunno/yeencloud/src/core/domain/config"
	"os"
)

func (config *Config) GetRunContextConfig() configDomain.RunContextConfig {
	hostname, err := os.Hostname()

	if err != nil {
		hostname = "unknown"
	}

	dir, err := os.Getwd()

	if err != nil {
		dir = "unknown"
	}

	runContextConfig := configDomain.RunContextConfig{
		Environment:      config.GetEnvStringOrDefault("ENV", "development"),
		Hostname:         hostname,
		WorkingDirectory: dir,
	}

	return runContextConfig
}
