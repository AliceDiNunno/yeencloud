package interactor

type ClusterAdapter interface {
	IsRunningInsideCluster() bool
	IsConfigurationValid(ClusterConfiguration []byte) bool
}
