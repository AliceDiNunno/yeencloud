package env

import configDomain "back/src/core/domain/config"

func (config *Config) GetDatabaseConfig() configDomain.DatabaseConfig {
	return configDomain.DatabaseConfig{
		Engine:   "postgres",
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "postgres",
		DbName:   "yeencloud_test",
	}
}
