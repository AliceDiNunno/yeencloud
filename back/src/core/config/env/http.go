package env

import configDomain "back/src/core/domain/config"

func (config *Config) GetHTTPConfig() configDomain.HTTPConfig {
	return configDomain.HTTPConfig{
		ListeningAddress: config.GetEnvStringOrDefault("LISTENING_ADDRESS", "0.0.0.0"),
		ListeningPort:    config.GetEnvIntOrDefault("PORT", 3000),

		FrontendURL: config.GetEnvStringOrDefault("FRONTEND_URL", "http://localhost:3000"),
	}
}
