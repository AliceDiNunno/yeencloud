package env

import configDomain "github.com/AliceDiNunno/yeencloud/internal/core/domain/config"

func (config *Config) GetKubernetesConfig() configDomain.KubernetesConfig {
	return configDomain.KubernetesConfig{
		UsingInternalConfig: config.GetEnvBoolOrDefault("KUBERNETES_INTERNAL_CONFIG", false),
		KubeconfigPath:      config.GetEnvStringOrDefault("KUBECONFIG_PATH", ""),
	}
}
