package db

import (
	"bitbucket.org/stack-rox/apollo/pkg/api/generated/api/v1"
)

// Storage is the interface for the persistent storage
type Storage interface {
	Load() error
	Close()

	AlertStorage
	BenchmarkResultsStorage
	BenchmarkStorage
	ClusterStorage
	DeploymentStorage
	ImagePolicyStorage
	ImageStorage
	NotifierStorage
	RegistryStorage
	ScannerStorage
}

// AlertStorage provides storage functionality for alerts.
type AlertStorage interface {
	GetAlert(id string) (*v1.Alert, bool, error)
	GetAlerts(request *v1.GetAlertsRequest) ([]*v1.Alert, error)
	AddAlert(alert *v1.Alert) error
	UpdateAlert(alert *v1.Alert) error
	RemoveAlert(id string) error
}

// BenchmarkStorage provides storage functionality for benchmarks results.
type BenchmarkStorage interface {
	GetBenchmark(name string) (*v1.Benchmark, bool, error)
	GetBenchmarks(request *v1.GetBenchmarksRequest) ([]*v1.Benchmark, error)
	AddBenchmark(benchmark *v1.Benchmark) error
	UpdateBenchmark(benchmark *v1.Benchmark) error
	RemoveBenchmark(name string) error
}

// BenchmarkResultsStorage provides storage functionality for benchmarks results.
type BenchmarkResultsStorage interface {
	GetBenchmarkResult(id string) (*v1.BenchmarkResult, bool, error)
	GetBenchmarkResults(request *v1.GetBenchmarkResultsRequest) ([]*v1.BenchmarkResult, error)
	AddBenchmarkResult(benchmark *v1.BenchmarkResult) error
}

// ClusterStorage provides storage functionality for clusters.
type ClusterStorage interface {
	GetCluster(name string) (*v1.Cluster, bool, error)
	GetClusters() ([]*v1.Cluster, error)
	AddCluster(cluster *v1.Cluster) error
	UpdateCluster(cluster *v1.Cluster) error
	RemoveCluster(name string) error
}

// DeploymentStorage provides storage functionality for deployments.
type DeploymentStorage interface {
	GetDeployment(id string) (*v1.Deployment, bool, error)
	GetDeployments(request *v1.GetDeploymentsRequest) ([]*v1.Deployment, error)
	AddDeployment(deployment *v1.Deployment) error
	UpdateDeployment(deployment *v1.Deployment) error
	RemoveDeployment(id string) error
}

// ImagePolicyStorage provides storage functionality for image policies.
type ImagePolicyStorage interface {
	GetImagePolicies(request *v1.GetImagePoliciesRequest) ([]*v1.ImagePolicy, error)
	AddImagePolicy(*v1.ImagePolicy) error
	UpdateImagePolicy(*v1.ImagePolicy) error
	RemoveImagePolicy(string) error
}

// ImageStorage provide storage functionality for images.
type ImageStorage interface {
	GetImages(request *v1.GetImagesRequest) ([]*v1.Image, error)
	AddImage(image *v1.Image) error
	UpdateImage(image *v1.Image) error
	RemoveImage(id string) error
}

// NotifierStorage provide storage functionality for notifiers
type NotifierStorage interface {
	GetNotifier(name string) (*v1.Notifier, bool, error)
	GetNotifiers(request *v1.GetNotifiersRequest) ([]*v1.Notifier, error)
	AddNotifier(notifier *v1.Notifier) error
	UpdateNotifier(notifier *v1.Notifier) error
	RemoveNotifier(name string) error
}

// RegistryStorage provide storage functionality for registries.
type RegistryStorage interface {
	GetRegistry(name string) (*v1.Registry, bool, error)
	GetRegistries(request *v1.GetRegistriesRequest) ([]*v1.Registry, error)
	AddRegistry(registry *v1.Registry) error
	UpdateRegistry(registry *v1.Registry) error
	RemoveRegistry(name string) error
}

// ScannerStorage provide storage functionality for scanner.
type ScannerStorage interface {
	GetScanner(name string) (*v1.Scanner, bool, error)
	GetScanners(request *v1.GetScannersRequest) ([]*v1.Scanner, error)
	AddScanner(scanner *v1.Scanner) error
	UpdateScanner(scanner *v1.Scanner) error
	RemoveScanner(name string) error
}
