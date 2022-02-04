package repositories

import (
	"fmt"
	"github.com/MrApr/PersonalTracker/models"
)

//CollectionRepo is an allies for collection
type CollectionRepo models.Collection

//ICollectionRepository Determines collection repository interface behaviors
type ICollectionRepository interface {
	//Get gets and return single collection
	Get(id int) error
	//GetAll Get and return all collections
	GetAll(name string) (*[]CollectionRepo, error)
	//CollectionTypes get and return all existing collection types
	CollectionTypes() []models.CollectionType
	//Create creates a new collection in db
	Create() bool
	//Edit edits existing collection in db
	Edit(editedData *CollectionRepo) error
	//Delete deletes existing collection in db
	Delete() error
}

//Get and returns single collection
func (col *CollectionRepo) Get(id int) error {
	result := models.DB.Where("id = ?", id).First(col)
	if result.Error != nil {
		//errors.Is(result.Error, gorm.ErrRecordNotFound)//Todo pay attention to this line
		return fmt.Errorf("%s: %s", "Query execution failed with error", result.Error)
	}
	return nil
}

//GetAll and return all collections
func (col *CollectionRepo) GetAll(name string) (*[]CollectionRepo, error) {
	var collections *[]CollectionRepo
	result := models.DB.Where("name LIKE ?", "%"+name+"%").Find(collections)
	if result.Error != nil {
		return nil, fmt.Errorf("%s: %s", "Unable to fetch batch collections with error", result.Error)
	}
	return collections, nil
}

//CollectionTypes get and return all existing collection types
func (col *CollectionRepo) CollectionTypes() []models.CollectionType {
	var collTypes []models.CollectionType
	collTypes = append(collTypes, models.DAILY)
	collTypes = append(collTypes, models.WEEKLY)
	collTypes = append(collTypes, models.MONTHLY)
	collTypes = append(collTypes, models.YEARLY)
	return collTypes
}

//Create a new collection into db
func (col *CollectionRepo) Create() bool {
	res := models.DB.Create(col)
	if res.Error != nil {
		return false
	}
	return true
}

//Edit existing collection in db
func (col *CollectionRepo) Edit(editedData *CollectionRepo) error {
	col.Title = editedData.Title
	col.Type = editedData.Type
	result := models.DB.Save(col)
	if result.Error != nil {
		return fmt.Errorf("%s: %s", "Unable to update existing data in db with error", result.Error)
	}
	return nil
}

//Delete existing collection in db
func (col *CollectionRepo) Delete() error {
	result := models.DB.Delete(col)
	if result.Error != nil {
		return fmt.Errorf("%s: %s", "Unable to delete Model with error", result.Error)
	}
	return nil
}
