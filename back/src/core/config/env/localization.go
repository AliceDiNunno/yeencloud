package env

import configDomain "github.com/AliceDiNunno/yeencloud/src/core/domain/config"

func (config *Config) GetLocalizationConfig() configDomain.LocalizationConfig {
	return configDomain.LocalizationConfig{
		DefaultLang: config.GetEnvStringOrDefault("API_LANG", "en-US"),
	}
}
