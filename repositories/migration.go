package repositories

import "fmt"

func Migrate() bool {
	var models []interface{} = []interface{}{Collection{}, Task{}}
	db := ConnectToDb()
	for _, model := range models {
		err := db.AutoMigrate(model)
		if err != nil {
			fmt.Println(err)
			return false
		}
	}
	return true
}
