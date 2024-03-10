package env

import configDomain "back/src/core/domain/config"

func (config *Config) GetKubernetesConfig() configDomain.KubernetesConfig {
	return configDomain.KubernetesConfig{
		UsingInternalConfig: config.GetEnvBoolOrDefault("KUBERNETES_INTERNAL_CONFIG", false),
		KubeconfigPath:      config.GetEnvStringOrDefault("KUBECONFIG_PATH", ""),
	}
}
