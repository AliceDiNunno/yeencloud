package k8s

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Cluster struct {
}

func (c Cluster) IsRunningInsideCluster() bool {
	_, err := rest.InClusterConfig()

	return err == nil
}

func (c Cluster) IsConfigurationValid(rawConfig []byte) bool {
	kubeConfig, err := clientcmd.RESTConfigFromKubeConfig(rawConfig)
	if err != nil {
		return false
	}

	_, err = kubernetes.NewForConfig(kubeConfig)

	return err == nil
}

func NewCluster() *Cluster {
	return &Cluster{}
}
