package mock

import (
	"errors"
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
			Provider:           uuid.Must(uuid.NewV4()).String(),
			ResourceType:       fmt.Sprintf("resource-type-%02d", i),
			ResourcesEvaluated: fmt.Sprintf("resources-evaluated-%02d", i),
			Tags:               uuid.Must(uuid.NewV4()).String(),
			Steampipe:          fmt.Sprintf("steampipe-%02d", i),
		}
	}

	return &Store{
		policies: policies,
	}
}

func (s *Store) QueryPolicy(name string) (*model.Policy, error) {
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

func (s *Store) UpdatePolicy(policy *model.Policy) error {
	s.policies[policy.Name] = policy
	return nil
}

func (s *Store) DeletePolicy(name string) error {
	delete(s.policies, name)
	return nil
}

func (s *Store) CreateTags(tags *model.Tags) error {
	return errors.New("not implemented")
}

func (s *Store) QueryTagsByID(id string) (*model.Tags, error) {
	return nil, errors.New("not implemented")
}

func (s *Store) QueryTagsByKeys(typ string, key string) (*model.Tags, error) {
	return nil, errors.New("not implemented")
}

func (s *Store) Close() {}
