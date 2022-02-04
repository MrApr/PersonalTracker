package repositories

import (
	"fmt"
	"github.com/MrApr/PersonalTracker/models"
)

//TaskRepo Defines task repository for task model
type TaskRepo models.Task

type ITaskRepository interface {
	//Get and return it as single task
	Get(taskId int) error
	//GetAll that are created and return them
	GetAll() (*[]TaskRepo, error)
	//Create creates a new task and returns it
	Create() error
	//Edit that exists in db and
	Edit(newData *TaskRepo) error
	//Delete existing task in db
	Delete() error
}

//Get and return it as single task
func (task *TaskRepo) Get(taskId int) error {
	result := models.DB.Where("id = ? ", taskId).First(task)
	if result.Error != nil {
		return fmt.Errorf("%s: %s", "Unable to find task error", result.Error)
	}
	return nil
}

//GetAll that are created and return them
func (task *TaskRepo) GetAll() (*[]TaskRepo, error) {
	var tasks *[]TaskRepo
	result := models.DB.Find(tasks)
	if result.Error != nil {
		return nil, fmt.Errorf("%s: %s", "Unable to find tasks error", result.Error)
	}
	return tasks, nil
}

//Create creates a new task and returns it
func (task *TaskRepo) Create() error {
	result := models.DB.Create(task)
	if result.Error != nil {
		return fmt.Errorf("%s: %s", "Unable to create a new task with error", result.Error)
	}
	return nil
}

//Edit that exists in db and
func (task *TaskRepo) Edit(newData *TaskRepo) error {
	task.Title = newData.Title
	task.IsCompleted = newData.IsCompleted
	task.DueToDate = newData.DueToDate
	task.CollectionId = newData.CollectionId

	result := models.DB.Save(task)
	if result.Error != nil {
		return fmt.Errorf("%s: %s", "Unable to update a new task with error", result.Error)
	}
	return nil
}

//Delete existing task in db
func (task *TaskRepo) Delete() error {
	result := models.DB.Delete(task)
	if result.Error != nil {
		return fmt.Errorf("%s: %s", "Unable to update a new task with error", result.Error)
	}
	return nil
}
