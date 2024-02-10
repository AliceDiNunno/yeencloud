package config

type VersionConfig struct {
	SHA           string
	Repository    string
	RepositoryURL string
}

func (config *Config) GetVersionConfig() VersionConfig {
	return VersionConfig{
		SHA:           config.GetEnvStringOrDefault("GITHUB_SHA", ""),
		Repository:    config.GetEnvStringOrDefault("GITHUB_REPOSITORY", ""),
		RepositoryURL: config.GetEnvStringOrDefault("GITHUB_REPOSITORY_URL", ""),
	}
}
