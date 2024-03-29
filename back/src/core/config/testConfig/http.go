package env

import configDomain "back/src/core/domain/config"

func (config *Config) GetHTTPConfig() configDomain.HTTPConfig {
	return configDomain.HTTPConfig{
		ListeningAddress: "0.0.0.0",
		ListeningPort:    3456,
		FrontendURL:      "*",
	}
}
