package env

import (
	"github.com/AliceDiNunno/yeencloud/src/adapters/log/reporting/rollbar"
)

func (config *Config) GetRollbarConfig() rollbar.Config {
	rollbarConfig := rollbar.Config{
		Token: config.GetEnvStringOrDefault("ROLLBAR_TOKEN", ""),
	}

	return rollbarConfig
}
