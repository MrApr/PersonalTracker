package models

import (
	"fmt"
)

//Migrate migrating db models into database
func Migrate() bool {
	var migrationModels []interface{} = []interface{}{Collection{}, Task{}}
	for _, model := range migrationModels {
		err := DB.AutoMigrate(&model)
		if err != nil {
			fmt.Println(err)
			return false
		}
	}
	return true
}
