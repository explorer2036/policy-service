package db

import (
	"policy-server/pkg/db/model"

	"github.com/jinzhu/gorm"
)

const (
	// TablePolicy - database table 'policy'
	TablePolicy = "policy"
)

// FindPolicyByName queries the policy by 'name'
func (s *Handler) FindPolicyByName(name string) (*model.Policy, error) {
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
