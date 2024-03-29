package config

type HTTPConfig struct {
	ListeningAddress string
	ListeningPort    int

	FrontendURL string
}

func (config *Config) GetHTTPConfig() HTTPConfig {
	return HTTPConfig{
		ListeningAddress: config.GetEnvStringOrDefault("LISTENING_ADDRESS", "0.0.0.0"),
		ListeningPort:    config.GetEnvIntOrDefault("PORT", 3000),

		FrontendURL: config.GetEnvStringOrDefault("FRONTEND_URL", "http://localhost:3000"),
	}
}
