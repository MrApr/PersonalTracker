package services

import (
	"fmt"
	"github.com/MrApr/PersonalTracker/models"
	"github.com/MrApr/PersonalTracker/repositories"
	"github.com/MrApr/PersonalTracker/server"
	"net/http"
)

//GetCollectionRequest request
type GetCollectionRequest struct {
	Name string `json:"name"`
}

//CreateCollectionReq for creating a new collection
type CreateCollectionReq struct {
	Name string                `json:"name"`
	Type models.CollectionType `json:"type"`
}

//DeleteCollectionReq Request for deleting collection
type DeleteCollectionReq struct {
	Id int `json:"id"`
}

//UpdateCollectionReq for updating existing collection
type UpdateCollectionReq struct {
	*DeleteCollectionReq
	*CreateCollectionReq
}

//GetCollections gets and returns all collections
func GetCollections(req *server.Request) error {
	var getCollectionReq *GetCollectionRequest = new(GetCollectionRequest)

	err := req.ParseBody(getCollectionReq)
	if err != nil {
		return req.Status(http.StatusInternalServerError).Json(&server.Response{
			"message": err.Error(),
		})
	}

	collectionRepo := new(repositories.CollectionRepo)
	collections, err := collectionRepo.GetAll(getCollectionReq.Name)
	if err != nil {
		return req.Status(http.StatusNotFound).Json(&server.Response{
			"message": err.Error(),
		})
	}

	return req.Status(http.StatusOK).Json(&server.Response{
		"data": collections,
	})
}

//CreateNewCollection for new collection creation
func CreateNewCollection(req *server.Request) error {
	var createReqCollectionReq *CreateCollectionReq = new(CreateCollectionReq)

	err := req.ParseBody(createReqCollectionReq)
	if err != nil {
		return req.Status(http.StatusInternalServerError).Json(&server.Response{
			"message": err.Error(),
		})
	}

	collectionRepo := new(repositories.CollectionRepo)
	collectionRepo.Title = createReqCollectionReq.Name
	collectionRepo.Type = createReqCollectionReq.Type

	created := collectionRepo.Create()
	if !created {
		return req.Status(http.StatusInternalServerError).Json(&server.Response{
			"message": "Cannot create A new collection",
		})
	}

	return req.Status(http.StatusCreated).Json(&server.Response{
		"message": fmt.Sprintf("%s %s %s", "Collection", collectionRepo.Title, "Created success fully"),
	})
}

//UpdateCollection updates existing collection
func UpdateCollection(req *server.Request) error {
	var updateCollection *UpdateCollectionReq = new(UpdateCollectionReq)

	err := req.ParseBody(updateCollection)
	if err != nil {
		return req.Status(http.StatusInternalServerError).Json(&server.Response{
			"message": err.Error(),
		})
	}

	collectionRepo := new(repositories.CollectionRepo)
	err = collectionRepo.Get(updateCollection.Id)
	if err != nil {
		return req.Status(http.StatusNotFound).Json(&server.Response{
			"message": "Collection Not found",
		})
	}

	editedRepo := new(repositories.CollectionRepo)
	editedRepo.Title = updateCollection.Name
	editedRepo.Type = updateCollection.Type

	err = collectionRepo.Edit(editedRepo)
	if err != nil {
		return req.Status(http.StatusInternalServerError).Json(&server.Response{
			"message": err.Error(),
		})
	}

	return req.Status(http.StatusCreated).Json(&server.Response{
		"message": "created successfully",
	})
}

//DeleteCollection that exists in Db
func DeleteCollection(req *server.Request) error {
	var delCollectionReq *DeleteCollectionReq = new(DeleteCollectionReq)

	err := req.ParseBody(delCollectionReq)
	if err != nil {
		return req.Status(http.StatusInternalServerError).Json(&server.Response{
			"message": err.Error(),
		})
	}

	delRepo := new(repositories.CollectionRepo)
	err = delRepo.Get(delCollectionReq.Id)
	if err != nil {
		return req.Status(http.StatusNotFound).Json(&server.Response{
			"message": err.Error(),
		})
	}

	err = delRepo.Delete()
	if err != nil {
		return req.Status(http.StatusInternalServerError).Json(&server.Response{
			"message": err.Error(),
		})
	}

	return req.Status(http.StatusOK).Json(&server.Response{
		"message": "Collection deleted successfully",
	})
}