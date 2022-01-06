package db

import (
	"policy-service/pkg/db/model"

	"github.com/jinzhu/gorm"
)

const (
	// TableProvider - database table 'provider'
	TableProvider = "provider"
)

// QueryProviders returns the providers
func (s *Handler) QueryProviders() ([]*model.Provider, error) {
	providers := []*model.Provider{}

	res := s.db.Table(TableProvider).Find(&providers)
	if res.Error != nil {
		// if there is no record found
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, res.Error
	}

	return providers, nil
}

// QueryProvider returns the provider by 'id'
func (s *Handler) QueryProvider(id string) (*model.Provider, error) {
	provider := model.Provider{}

	res := s.db.Table(TableProvider).Where("id = ?", id).First(&provider)
	if res.Error != nil {
		// if there is no record found
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, res.Error
	}

	return &provider, nil
}

// CreateProvider inserts a new provider
func (s *Handler) CreateProvider(provider *model.Provider) error {
	return s.db.Table(TableProvider).Create(provider).Error
}

// UpdateProvider updates a special provider
func (s *Handler) UpdateProvider(provider *model.Provider) error {
	return s.db.Table(TableProvider).Update(provider).Error
}
