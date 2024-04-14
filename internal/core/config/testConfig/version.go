package env

import configDomain "github.com/AliceDiNunno/yeencloud/internal/core/domain/config"

func (config *Config) GetVersionConfig() configDomain.VersionConfig {
	return configDomain.VersionConfig{
		SHA:           "",
		Repository:    "",
		RepositoryURL: "",
	}
}
