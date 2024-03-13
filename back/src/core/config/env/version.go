package env

import configDomain "back/src/core/domain/config"

func (config *Config) GetVersionConfig() configDomain.VersionConfig {
	versionConfig := configDomain.VersionConfig{
		SHA:           config.GetEnvStringOrDefault("GITHUB_SHA", ""),
		Repository:    config.GetEnvStringOrDefault("GITHUB_REPOSITORY", ""),
		RepositoryURL: config.GetEnvStringOrDefault("GITHUB_REPOSITORY_URL", ""),

		Present: false,
	}

	if versionConfig.SHA != "" || versionConfig.Repository != "" || versionConfig.RepositoryURL != "" {
		versionConfig.Present = true
	}

	return versionConfig
}
