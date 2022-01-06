package mock

import (
	"errors"
	"fmt"
	"policy-service/pkg/db"
	"policy-service/pkg/db/model"

	"github.com/gofrs/uuid"
)

type Store struct {
	policies map[string]*model.Policy
}

func NewStore() db.Repository {
	policies := make(map[string]*model.Policy)

	for i := 0; i < 3; i++ {
		id := uuid.Must(uuid.NewV4()).String()
		policies[id] = &model.Policy{
			ID:                 id,
			Name:               fmt.Sprintf("policy-%02d", i),
			State:              "active",
			Provider:           uuid.Must(uuid.NewV4()).String(),
			ResourceType:       fmt.Sprintf("resource-type-%02d", i),
			ResourcesEvaluated: fmt.Sprintf("resources-evaluated-%02d", i),
			Tags:               []string{uuid.Must(uuid.NewV4()).String()},
			Steampipe:          fmt.Sprintf("steampipe-%02d", i),
		}
	}

	return &Store{
		policies: policies,
	}
}

func (s *Store) QueryPolicy(id string) (*model.Policy, error) {
	for key, policy := range s.policies {
		if key == id {
			return policy, nil
		}
	}
	return nil, nil
}

func (s *Store) QueryPolicies() ([]*model.Policy, error) {
	policies := []*model.Policy{}
	for _, policy := range s.policies {
		policies = append(policies, policy)
	}
	return policies, nil
}

func (s *Store) CreatePolicy(policy *model.Policy) error {
	s.policies[policy.ID] = policy
	return nil
}

func (s *Store) UpdatePolicy(policy *model.Policy) error {
	s.policies[policy.ID] = policy
	return nil
}

func (s *Store) DeletePolicy(id string) error {
	delete(s.policies, id)
	return nil
}

func (s *Store) CreateTag(tag *model.Tag) error {
	return errors.New("not implemented")
}
func (s *Store) QueryTag(id string) (*model.Tag, error) {
	return nil, errors.New("not implemented")
}
func (s *Store) QueryTags() ([]*model.Tag, error) {
	return nil, errors.New("not implemented")
}

// Benchmark CRUD
func (s *Store) QueryBenchmarks() ([]*model.Benchmark, error) {
	return nil, errors.New("not implemented")
}
func (s *Store) QueryBenchmark(id string) (*model.Benchmark, error) {
	return nil, errors.New("not implemented")
}
func (s *Store) CreateBenchmark(benchmark *model.Benchmark) error {
	return errors.New("not implemented")
}
func (s *Store) UpdateBenchmark(benchmark *model.Benchmark) error {
	return errors.New("not implemented")
}
func (s *Store) DeleteBenchmark(id string) error {
	return errors.New("not implemented")
}

// Provider CRU
func (s *Store) CreateProvider(provider *model.Provider) error {
	return errors.New("not implemented")
}
func (s *Store) UpdateProvider(provider *model.Provider) error {
	return errors.New("not implemented")
}
func (s *Store) QueryProvider(id string) (*model.Provider, error) {
	return nil, errors.New("not implemented")
}
func (s *Store) QueryProviders() ([]*model.Provider, error) {
	return nil, errors.New("not implemented")
}

// ProviderType CRU
func (s *Store) CreateProviderType(providerType *model.ProviderType) error {
	return errors.New("not implemented")
}
func (s *Store) UpdateProviderType(providerType *model.ProviderType) error {
	return errors.New("not implemented")
}
func (s *Store) QueryProviderType(id string) (*model.ProviderType, error) {
	return nil, errors.New("not implemented")
}
func (s *Store) QueryProviderTypes() ([]*model.ProviderType, error) {
	return nil, errors.New("not implemented")
}

func (s *Store) Close() {}
