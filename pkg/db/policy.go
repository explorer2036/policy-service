package db

import (
	"policy-service/pkg/db/model"

	"github.com/jinzhu/gorm"
)

const (
	// TablePolicy - database table 'policy'
	TablePolicy = "policy"
)

// QueryPolicies returns the policies
func (s *Handler) QueryPolicies() ([]*model.Policy, error) {
	policies := []*model.Policy{}

	res := s.db.Table(TablePolicy).Find(&policies)
	if res.Error != nil {
		// if there is no record found
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, res.Error
	}

	return policies, nil
}

// QueryPolicy returns the policy by 'id'
func (s *Handler) QueryPolicy(id string) (*model.Policy, error) {
	policy := model.Policy{}

	res := s.db.Table(TablePolicy).Where("id = ?", id).First(&policy)
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

// DeletePolicy deletes the policy by id
func (s *Handler) DeletePolicy(id string) error {
	return s.db.Table(TablePolicy).Exec("Delete from policy where id = ?", id).Error
}
