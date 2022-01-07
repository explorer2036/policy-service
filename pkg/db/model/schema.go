package model

import (
	"time"

	"github.com/lib/pq"
)

type Policy struct {
	ID                 string         `gorm:"column:id"`
	Name               string         `gorm:"column:name"`
	State              string         `gorm:"column:state"`
	Provider           string         `gorm:"column:provider"`
	ResourceType       string         `gorm:"column:resource_type"`
	ResourcesEvaluated string         `gorm:"column:resources_evaluated"`
	Tags               pq.StringArray `gorm:"column:tags"`
	Steampipe          string         `gorm:"column:steampipe"`
	CreateTime         time.Time      `gorm:"column:create_time"`
	UpdateTime         time.Time      `gorm:"column:update_time"`
}

type Tag struct {
	ID         string    `gorm:"column:id"`
	Type       string    `gorm:"column:type"`
	Key        string    `gorm:"column:key"`
	Value      string    `gorm:"column:value"`
	State      string    `gorm:"column:state"`
	Provider   string    `gorm:"column:provider"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}

type ProviderType struct {
	ID          string    `gorm:"column:id"`
	Name        string    `gorm:"column:name"`
	State       string    `gorm:"column:state"`
	Description string    `gorm:"column:description"`
	CreateTime  time.Time `gorm:"column:create_time"`
	UpdateTime  time.Time `gorm:"column:update_time"`
}

type Provider struct {
	ID           string    `gorm:"column:id"`
	Name         string    `gorm:"column:name"`
	Url          string    `gorm:"column:url"`
	ProviderType string    `gorm:"column:provider_type"`
	State        string    `gorm:"column:state"`
	Description  string    `gorm:"column:description"`
	CreateTime   time.Time `gorm:"column:create_time"`
	UpdateTime   time.Time `gorm:"column:update_time"`
}

type Benchmark struct {
	ID                 string         `gorm:"column:id"`
	Name               string         `gorm:"column:name"`
	State              string         `gorm:"column:state"`
	Provider           string         `gorm:"column:provider"`
	ResourceType       string         `gorm:"column:resource_type"`
	ResourcesEvaluated string         `gorm:"column:resources_evaluated"`
	Tags               pq.StringArray `gorm:"column:tags"`
	Policies           pq.StringArray `gorm:"column:policies"`
	Description        string         `gorm:"column:description"`
	CreateTime         time.Time      `gorm:"column:create_time"`
	UpdateTime         time.Time      `gorm:"column:update_time"`
}
