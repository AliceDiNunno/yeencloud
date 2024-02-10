package usecases

func (i interactor) hasKubernetesCluster() bool {
	if i.cluster.IsRunningInsideCluster() {
		return true
	}

	kubeconfig := i.getKubeconfigFromSettings()

	if kubeconfig == "" {
		return false
	}

	return i.cluster.IsConfigurationValid([]byte(kubeconfig))
}
