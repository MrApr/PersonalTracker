package repositories

import (
	"fmt"
	"github.com/MrApr/PersonalTracker/models"
)

func Migrate() bool {
	var migrationModels []interface{} = []interface{}{Collection{}, Task{}}
	for _, model := range migrationModels {
		err := models.DB.AutoMigrate(model)
		if err != nil {
			fmt.Println(err)
			return false
		}
	}
	return true
}
