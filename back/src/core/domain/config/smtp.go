package config

type MailConfig struct {
	Host string
	Port int

	From     string
	Password string

	TemplateDirectory string
}
