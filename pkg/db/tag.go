package db

import (
	"policy-service/pkg/db/model"

	"github.com/jinzhu/gorm"
)

const (
	// TableTag - database table 'tag'
	TableTag = "tag"
)

// QueryTags returns the tags
func (s *Handler) QueryTags() ([]*model.Tag, error) {
	tags := []*model.Tag{}

	res := s.db.Table(TableTag).Find(&tags)
	if res.Error != nil {
		// if there is no record found
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, res.Error
	}

	return tags, nil
}

// QueryTag returns the tag by 'id'
func (s *Handler) QueryTag(id string) (*model.Tag, error) {
	tag := model.Tag{}

	res := s.db.Table(TableTag).Where("id = ?", id).First(&tag)
	if res.Error != nil {
		// if there is no record found
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, res.Error
	}

	return &tag, nil
}

// CreateTag inserts a new tag
func (s *Handler) CreateTag(tag *model.Tag) error {
	return s.db.Table(TableTag).Create(tag).Error
}
