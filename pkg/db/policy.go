package db

import (
	"policy-server/pkg/db/model"

	"github.com/jinzhu/gorm"
)

const (
	// TablePolicy - database table 'policy'
	TablePolicy = "policy"
)

// QueryPolicy returns the policy by 'name'
func (s *Handler) QueryPolicy(name string) (*model.Policy, error) {
	policy := model.Policy{}

	res := s.db.Table(TablePolicy).Where("name = ?", name).First(&policy)
	if res.Error != nil {
		// if there is no record found
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, res.Error
	}

	return &policy, nil
}

// CreatePolicy inserts a new policy
func (s *Handler) CreatePolicy(policy *model.Policy) error {
	return s.db.Table(TablePolicy).Create(policy).Error
}

// UpdatePolicy updates a special policy
func (s *Handler) UpdatePolicy(policy *model.Policy) error {
	return s.db.Table(TablePolicy).Update(policy).Error
}

// DeletePolicyByName deletes the policy by name
func (s *Handler) DeletePolicy(name string) error {
	return s.db.Table(TablePolicy).Exec("Delete from policy where name = ?", name).Error
}
