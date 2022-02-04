package repositories

import (
	"fmt"
	"github.com/MrApr/PersonalTracker/models"
)

//CollectionRepo is an allies for collection
type CollectionRepo models.Collection

//ICollectionRepository Determines collection repository interface behaviors
type ICollectionRepository interface {
	//GetCollection gets and return single collection
	GetCollection(id int) (*CollectionRepo, error)
	//GetAllCollections Get and return all collections
	GetAllCollections(name string) (*[]CollectionRepo, error)
	//GetAllCollectionTypes get and return all existing collection types
	GetAllCollectionTypes() []models.CollectionType
	//CreateCollection creates a new collection in db
	CreateCollection() bool
	//EditCollection edits existing collection in db
	EditCollection(editedData *CollectionRepo) error
	//DeleteCollection deletes existing collection in db
	DeleteCollection() error
}

//GetCollection Get and returns single collection
func (col *CollectionRepo) GetCollection(id int) (*CollectionRepo, error) { //Todo define custom error types for better err handling
	result := models.DB.Where("id = ?", id).First(col)
	if result.Error != nil {
		//errors.Is(result.Error, gorm.ErrRecordNotFound)//Todo pay attention to this line
		return nil, fmt.Errorf("%s: %s", "Query execution failed with error", result.Error)
	}
	return col, nil
}

//GetAllCollections Get and return all collections
func (col *CollectionRepo) GetAllCollections(name string) (*[]CollectionRepo, error) {
	var collections *[]CollectionRepo
	result := models.DB.Where("name LIKE ?", "%"+name+"%").Find(collections)
	if result.Error != nil {
		return nil, fmt.Errorf("%s: %s", "Unable to fetch batch collections with error", result.Error)
	}
	return collections, nil
}

//GetAllCollectionTypes get and return all existing collection types
func (col *CollectionRepo) GetAllCollectionTypes() []models.CollectionType {
	var collTypes []models.CollectionType
	collTypes = append(collTypes, models.DAILY)
	collTypes = append(collTypes, models.WEEKLY)
	collTypes = append(collTypes, models.MONTHLY)
	collTypes = append(collTypes, models.YEARLY)
	return collTypes
}

func (col *CollectionRepo) CreateCollection() bool {
	res := models.DB.Create(col)
	if res.Error != nil {
		return false
	}
	return true
}

//EditCollection edits existing collection in db
func (col *CollectionRepo) EditCollection(editedData *CollectionRepo) error {
	col.Title = editedData.Title
	col.Type = editedData.Type
	result := models.DB.Save(col)
	if result.Error != nil {
		return fmt.Errorf("%s: %s", "Unable to update existing data in db with error", result.Error)
	}
	return nil
}

//DeleteCollection deletes existing collection in db
func (col *CollectionRepo) DeleteCollection() error {
	result := models.DB.Delete(col)
	if result.Error != nil {
		return fmt.Errorf("%s: %s", "Unable to delete Model with error", result.Error)
	}
	return nil
}
