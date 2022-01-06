package db

import (
	"policy-server/pkg/db/model"

	"github.com/jinzhu/gorm"
)

const (
	// TableTags - database table 'tags'
	TableTags = "tags"
)

// QueryTagsByID returns the tags by 'id'
func (s *Handler) QueryTagsByID(id string) (*model.Tags, error) {
	tags := model.Tags{}

	res := s.db.Table(TableTags).Where("id = ?", id).First(&tags)
	if res.Error != nil {
		// if there is no record found
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, res.Error
	}

	return &tags, nil
}

// QueryTagsByKeys returns the tags by 'type and key'
func (s *Handler) QueryTagsByKeys(typ string, key string) (*model.Tags, error) {
	tags := model.Tags{}

	res := s.db.Table(TableTags).Where("type = ? and key = ?", typ, key).First(&tags)
	if res.Error != nil {
		// if there is no record found
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, res.Error
	}

	return &tags, nil
}

// CreateTags inserts a new tags
func (s *Handler) CreateTags(tags *model.Tags) error {
	return s.db.Table(TableTags).Create(tags).Error
}
