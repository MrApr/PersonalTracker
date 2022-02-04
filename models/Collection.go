package models

import "gorm.io/gorm"

type CollectionType string

const (
	DAILY   CollectionType = "daily"
	WEEKLY  CollectionType = "weekly"
	MONTHLY CollectionType = "monthly"
	YEARLY  CollectionType = "yearly"
)

//Collection defines a custom model structure for collection
type Collection struct {
	gorm.Model
	Title string         `json:"title" gorm:"not null;type:varchar(255)"`
	Type  CollectionType `json:"type" gorm:"not null; type:varchar(255)"`
}
