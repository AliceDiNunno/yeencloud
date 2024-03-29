package config

type HTTPConfig struct {
	ListeningAddress string `json:"-"`
	ListeningPort    int    `json:"-"`

	FrontendURL string `json:"-"`
}
