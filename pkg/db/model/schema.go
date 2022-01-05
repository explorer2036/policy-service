package model

type Policy struct {
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
}
