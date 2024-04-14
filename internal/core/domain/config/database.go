package config

type DatabaseConfig struct {
	Engine   string
	Host     string
	Port     int
	User     string
	Password string `json:"-"`
	DbName   string
}
