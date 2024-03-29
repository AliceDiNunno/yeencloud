package config

type KubernetesConfig struct {
	UsingInternalConfig bool   `json:"-"`
	KubeconfigPath      string `json:"-"`
}
