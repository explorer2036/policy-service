package db

import (
	"policy-service/pkg/db/model"

	"github.com/jinzhu/gorm"
)

const (
	// TableProviderType - database table 'provider_type'
	TableProviderType = "provider_type"
)

// QueryProviderTypes returns the provider types
func (s *Handler) QueryProviderTypes() ([]*model.ProviderType, error) {
	providerTypes := []*model.ProviderType{}

	res := s.db.Table(TableProviderType).Find(&providerTypes)
	if res.Error != nil {
		// if there is no record found
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, res.Error
	}

	return providerTypes, nil
}

// QueryProviderType returns the provider type by 'id'
func (s *Handler) QueryProviderType(id string) (*model.ProviderType, error) {
	providerType := model.ProviderType{}

	res := s.db.Table(TableProviderType).Where("id = ?", id).First(&providerType)
	if res.Error != nil {
		// if there is no record found
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, res.Error
	}

	return &providerType, nil
}

// CreateProviderType inserts a new provider type
func (s *Handler) CreateProviderType(providerType *model.ProviderType) error {
	return s.db.Table(TableProviderType).Create(providerType).Error
}

// UpdateProviderType updates a special provider type
func (s *Handler) UpdateProviderType(providerType *model.ProviderType) error {
	return s.db.Table(TableProviderType).Update(providerType).Error
}
