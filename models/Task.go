package models

import (
	"gorm.io/gorm"
	"time"
)

//Task for model structure for task entity in system
type Task struct {
	gorm.Model
	CollectionId int        `json:"collection_id" gorm:"not null;type:UNSIGNED"`
	Collection   Collection `json:"collection"  gorm:"foreignKey:collection_id"`
	Title        string     `json:"title" gorm:"not null;type:varchar(255)"`
	IsCompleted  bool       `json:"is_completed" gorm:"not null;default:false"`
	DueToDate    time.Time  `json:"due_to_date;omitempty" gorm:"type:timestamp"`
}
