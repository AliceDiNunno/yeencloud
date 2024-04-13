package env

import (
	configDomain "github.com/AliceDiNunno/yeencloud/src/core/domain/config"
)

func (config *Config) GetMailConfig() configDomain.MailConfig {
	mailConfig := configDomain.MailConfig{
		Host:              config.GetEnvStringOrDefault("SMTP_HOST", ""),
		Port:              config.GetEnvIntOrDefault("SMTP_PORT", 0),
		From:              config.GetEnvStringOrDefault("SMTP_FROM", ""),
		Password:          config.GetEnvStringOrDefault("SMTP_PASSWORD", ""),
		TemplateDirectory: config.GetEnvStringOrDefault("SMTP_TEMPLATE_DIRECTORY", "./template/mail/"),
	}

	return mailConfig
}
