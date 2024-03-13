package config

type VersionConfig struct {
	SHA           string
	Repository    string
	RepositoryURL string

	Present bool
}
