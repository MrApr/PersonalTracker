package services

import (
	"github.com/MrApr/PersonalTracker/repositories"
	"github.com/MrApr/PersonalTracker/server"
	"net/http"
	"time"
)

//DeleteTaskReq for deleting existing task
type DeleteTaskReq struct {
	Id int `json:"id"`
}

//CreateTaskReq for a new task
type CreateTaskReq struct {
	CollectionId int       `json:"collection_id"`
	Title        string    `json:"title"`
	IsCompleted  bool      `json:"is_completed"`
	DueToDate    time.Time `json:"due_to_date;omitempty"`
}

//EditTaskReq for editing existing task
type EditTaskReq struct {
	*DeleteTaskReq
	*CreateTaskReq
}

//GetTasks that defined in system
func GetTasks(req *server.Request) error {
	taskRepo := new(repositories.TaskRepo)
	tasks, err := taskRepo.GetAll()

	if err != nil {
		return req.Status(http.StatusNotFound).Json(&server.Response{
			"message": err.Error(),
		})
	}

	return req.Status(http.StatusOK).Json(&server.Response{
		"data": tasks,
	})
}

//CreateTask in db
func CreateTask(req *server.Request) error {
	var createTaskReq *CreateTaskReq = new(CreateTaskReq)

	err := req.ParseBody(createTaskReq)
	if err != nil {
		return req.Status(http.StatusInternalServerError).Json(&server.Response{
			"message": err.Error(),
		})
	}

	createTaskRepo := new(repositories.TaskRepo)
	createTaskRepo.Title = createTaskReq.Title
	createTaskRepo.CollectionId = createTaskReq.CollectionId
	createTaskRepo.IsCompleted = createTaskReq.IsCompleted
	createTaskRepo.DueToDate = createTaskReq.DueToDate

	err = createTaskRepo.Create()
	if err != nil {
		return req.Status(http.StatusInternalServerError).Json(&server.Response{
			"message": err.Error(),
		})
	}

	return req.Status(http.StatusCreated).Json(&server.Response{
		"message": "Created successfully",
	})
}

//EditTask that exists in db
func EditTask(req *server.Request) error {
	var editTaskReq *EditTaskReq = new(EditTaskReq)

	err := req.ParseBody(editTaskReq)
	if err != nil {
		return req.Status(http.StatusInternalServerError).Json(&server.Response{
			"message": err.Error(),
		})
	}

	taskRepo := new(repositories.TaskRepo)

	err = taskRepo.Get(editTaskReq.Id)
	if err != nil {
		return req.Status(http.StatusNotFound).Json(&server.Response{
			"message": err.Error(),
		})
	}

	newTask := new(repositories.TaskRepo)
	newTask.Title = editTaskReq.Title
	newTask.CollectionId = editTaskReq.CollectionId
	newTask.IsCompleted = editTaskReq.IsCompleted
	newTask.DueToDate = editTaskReq.DueToDate

	err = taskRepo.Edit(newTask)
	if err != nil {
		return req.Status(http.StatusInternalServerError).Json(&server.Response{
			"message": err.Error(),
		})
	}

	return req.Status(http.StatusCreated).Json(&server.Response{
		"message": "Edit successfully",
	})
}

//DeleteTask that exists in db
func DeleteTask(req *server.Request) error {
	var deleteTaskReq *DeleteTaskReq = new(DeleteTaskReq)

	err := req.ParseBody(deleteTaskReq)
	if err != nil {
		return req.Status(http.StatusInternalServerError).Json(&server.Response{
			"message": err.Error(),
		})
	}

	taskRepo := new(repositories.TaskRepo)

	err = taskRepo.Get(deleteTaskReq.Id)
	if err != nil {
		return req.Status(http.StatusNotFound).Json(&server.Response{
			"message": err.Error(),
		})
	}

	err = taskRepo.Delete()
	if err != nil {
		return req.Status(http.StatusInternalServerError).Json(&server.Response{
			"message": err.Error(),
		})
	}

	return req.Status(http.StatusOK).Json(&server.Response{
		"message": "Task deleted successfully",
	})
}
