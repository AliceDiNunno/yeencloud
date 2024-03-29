package env

import configDomain "github.com/AliceDiNunno/yeencloud/src/core/domain/config"

func (config *Config) GetKubernetesConfig() configDomain.KubernetesConfig {
	return configDomain.KubernetesConfig{
		UsingInternalConfig: false,
		KubeconfigPath:      "",
	}
}
