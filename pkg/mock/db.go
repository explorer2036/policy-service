package mock

import (
	"fmt"
	"policy-server/pkg/db"
	"policy-server/pkg/db/model"

	"github.com/gofrs/uuid"
)

type Store struct {
	policies map[string]*model.Policy
}

func NewStore() db.Repository {
	policies := make(map[string]*model.Policy)

	for i := 0; i < 3; i++ {
		name := fmt.Sprintf("policy-%02d", i)
		policies[name] = &model.Policy{
			ID:                 uuid.Must(uuid.NewV4()).String(),
			Name:               name,
			State:              "active",
			ProviderName:       fmt.Sprintf("provider-name-%02d", i),
			ResourceType:       fmt.Sprintf("resource-type-%02d", i),
			ResourcesEvaluated: fmt.Sprintf("resources-evaluated-%02d", i),
			Tags:               fmt.Sprintf("tags-%02d", i),
			Steampipe:          fmt.Sprintf("steampipe-%02d", i),
		}
	}

	return &Store{
		policies: policies,
	}
}

func (s *Store) FindPolicyByName(name string) (*model.Policy, error) {
	for key, policy := range s.policies {
		if key == name {
			return policy, nil
		}
	}
	return nil, nil
}

func (s *Store) CreatePolicy(policy *model.Policy) error {
	s.policies[policy.Name] = policy
	return nil
}

func (s *Store) Close() {}
