package db

import "policy-server/pkg/db/model"

type Repository interface {
	// Policy CRUD
	QueryPolicy(name string) (*model.Policy, error)
	CreatePolicy(policy *model.Policy) error
	UpdatePolicy(policy *model.Policy) error
	DeletePolicy(name string) error

	// Tags CR
	CreateTags(tags *model.Tags) error
	QueryTagsByID(id string) (*model.Tags, error)
	QueryTagsByKeys(typ string, key string) (*model.Tags, error)

	Close()
}
