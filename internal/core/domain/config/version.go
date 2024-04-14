package config

type VersionConfig struct {
	SHA           string `json:"sha,omitempty"`
	Repository    string `json:"repository,omitempty"`
	RepositoryURL string `json:"repositoryUrl,omitempty"`

	Present bool `json:"present"`
}
