package env

import (
	"github.com/AliceDiNunno/yeencloud/src/adapters/log/reporting/rollbar"
)

func (config *Config) GetRollbarConfig() rollbar.RollbarConfig {
	rollbarConfig := rollbar.RollbarConfig{
		Token: config.GetEnvStringOrDefault("ROLLBAR_TOKEN", ""),
	}

	return rollbarConfig
}
