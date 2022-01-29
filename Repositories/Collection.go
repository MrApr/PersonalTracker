package Repositories

type Collection struct {
	Id    uint   `json:"id" gorm:"not null;primaryKey;autoIncrement"`
	Title string `json:"title" gorm:"not null;type:varchar(255)"`
	Type  string `json:"type" gorm:"not null;type:varchar(255)"`
}
