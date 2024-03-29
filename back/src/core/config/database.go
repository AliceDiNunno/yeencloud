package config

type DatabaseConfig struct {
	Engine   string
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

func (config *Config) GetDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Engine:   config.GetEnvStringOrDefault("DATABASE_ENGINE", "postgres"),
		Host:     config.GetEnvStringOrDefault("DATABASE_HOST", "localhost"),
		Port:     config.GetEnvIntOrDefault("DATABASE_PORT", 5432),
		User:     config.GetEnvStringOrDefault("DATABASE_USER", "postgres"),
		Password: config.GetEnvStringOrDefault("DATABASE_PASSWORD", "postgres"),
		DbName:   config.GetEnvStringOrDefault("DATABASE_NAME", ""),
	}
}
