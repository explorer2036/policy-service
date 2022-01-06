package model

import "time"

type Policy struct {
	ID                 string    `gorm:"column:id"`
	Name               string    `gorm:"column:name"`
	State              string    `gorm:"column:state"`
	ProviderName       string    `gorm:"column:provider_name"`
	ResourceType       string    `gorm:"column:resource_type"`
	ResourcesEvaluated string    `gorm:"column:resources_evaluated"`
	Tags               string    `gorm:"column:tags"`
	Steampipe          string    `gorm:"column:steampipe"`
	CreateTime         time.Time `gorm:"column:create_time"`
	UpdateTime         time.Time `gorm:"column:update_time"`
}
