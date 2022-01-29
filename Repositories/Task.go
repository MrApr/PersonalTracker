package Repositories

import "time"

type Task struct {
	Id           int        `json:"id" gorm:"not null;primaryKey;autoIncrement"`
	CollectionId int        `json:"collection_id" gorm:"not null;type:UNSIGNED"`
	Collection   Collection `json:"collection"  gorm:"foreignKey:collection_id"`
	Title        string     `json:"title" gorm:"not null;type:varchar(255)"`
	IsCompleted  bool       `json:"is_completed" gorm:"not null;default:false"`
	DueToDate    time.Time  `json:"due_to_date;omitempty" gorm:"type:timestamp"`
	CreatedAt    time.Time  `json:"created_at" gorm:"autoCreateTime;type:timestamp"`
}
