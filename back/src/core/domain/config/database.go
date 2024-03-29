package config

type DatabaseConfig struct {
	Engine   string `json:"-"`
	Host     string `json:"-"`
	Port     int    `json:"-"`
	User     string `json:"-"`
	Password string `json:"-"`
	DbName   string `json:"-"`
}
