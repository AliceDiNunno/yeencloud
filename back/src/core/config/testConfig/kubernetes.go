package env

import configDomain "back/src/core/domain/config"

func (config *Config) GetKubernetesConfig() configDomain.KubernetesConfig {
	return configDomain.KubernetesConfig{
		UsingInternalConfig: false,
		KubeconfigPath:      "",
	}
}
