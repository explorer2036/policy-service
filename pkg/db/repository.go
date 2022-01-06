package db

import "policy-server/pkg/db/model"

type Repository interface {
	// Policy CRUD
	FindPolicyByName(name string) (*model.Policy, error)
	CreatePolicy(policy *model.Policy) error

	Close()
}
