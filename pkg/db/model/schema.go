package model

import "time"

type Policy struct {
	ID                 string    `gorm:"column:id"`
	Name               string    `gorm:"column:name"`
	State              string    `gorm:"column:state"`
	Provider           string    `gorm:"column:provider"`
	ResourceType       string    `gorm:"column:resource_type"`
	ResourcesEvaluated string    `gorm:"column:resources_evaluated"`
	Tags               string    `gorm:"column:tags"`
	Steampipe          string    `gorm:"column:steampipe"`
	CreateTime         time.Time `gorm:"column:create_time"`
	UpdateTime         time.Time `gorm:"column:update_time"`
}

type Tags struct {
	ID         string    `gorm:"column:id"`
	Type       string    `gorm:"column:type"`
	Key        string    `gorm:"column:key"`
	Value      string    `gorm:"column:value"`
	State      string    `gorm:"column:state"`
	Provider   string    `gorm:"column:provider"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}
