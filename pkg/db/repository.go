package db

import "policy-service/pkg/db/model"

type Repository interface {
	// Policy CRUD
	QueryPolicies() ([]*model.Policy, error)
	QueryPolicy(id string) (*model.Policy, error)
	CreatePolicy(policy *model.Policy) error
	UpdatePolicy(policy *model.Policy) error
	DeletePolicy(id string) error

	// Tag CR
	CreateTag(tag *model.Tag) error
	QueryTag(id string) (*model.Tag, error)
	QueryTags() ([]*model.Tag, error)

	// Benchmark CRUD
	QueryBenchmarks() ([]*model.Benchmark, error)
	QueryBenchmark(id string) (*model.Benchmark, error)
	CreateBenchmark(benchmark *model.Benchmark) error
	UpdateBenchmark(benchmark *model.Benchmark) error
	DeleteBenchmark(id string) error

	// Provider CRU
	CreateProvider(provider *model.Provider) error
	UpdateProvider(provider *model.Provider) error
	QueryProvider(id string) (*model.Provider, error)
	QueryProviders() ([]*model.Provider, error)

	// ProviderType CRU
	CreateProviderType(providerType *model.ProviderType) error
	UpdateProviderType(providerType *model.ProviderType) error
	QueryProviderType(id string) (*model.ProviderType, error)
	QueryProviderTypes() ([]*model.ProviderType, error)

	Close()
}
