package env

import configDomain "back/src/core/domain/config"

func (config *Config) GetVersionConfig() configDomain.VersionConfig {
	return configDomain.VersionConfig{
		SHA:           "",
		Repository:    "",
		RepositoryURL: "",
	}
}
