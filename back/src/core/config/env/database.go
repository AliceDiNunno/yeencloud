package env

import configDomain "github.com/AliceDiNunno/yeencloud/src/core/domain/config"

func (config *Config) GetDatabaseConfig() configDomain.DatabaseConfig {
	return configDomain.DatabaseConfig{
		Engine:   config.GetEnvStringOrDefault("DATABASE_ENGINE", "postgres"),
		Host:     config.GetEnvStringOrDefault("DATABASE_HOST", "localhost"),
		Port:     config.GetEnvIntOrDefault("DATABASE_PORT", 5432),
		User:     config.GetEnvStringOrDefault("DATABASE_USER", "postgres"),
		Password: config.GetEnvStringOrDefault("DATABASE_PASSWORD", "postgres"),
		DbName:   config.GetEnvStringOrDefault("DATABASE_NAME", ""),
	}
}
