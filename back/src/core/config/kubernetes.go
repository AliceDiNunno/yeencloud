package config

type KubernetesConfig struct {
	UsingInternalConfig bool
	KubeconfigPath      string
}

func (config *Config) GetKubernetesConfig() KubernetesConfig {
	return KubernetesConfig{
		UsingInternalConfig: config.GetEnvBoolOrDefault("KUBERNETES_INTERNAL_CONFIG", false),
		KubeconfigPath:      config.GetEnvStringOrDefault("KUBECONFIG_PATH", ""),
	}
}
