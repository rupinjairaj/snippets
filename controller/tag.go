package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rupinjairaj/snippet/entity"
	"github.com/rupinjairaj/snippet/errors"
	"github.com/rupinjairaj/snippet/service"
)

type TagController interface {
	GetTags(response http.ResponseWriter, request *http.Request)
	AddTag(response http.ResponseWriter, request *http.Request)
}

type tagCtrl struct {
	tagSer service.TagService
}

func NewTagController(ser service.TagService) TagController {
	return &tagCtrl{
		tagSer: ser,
	}
}

func (sc *tagCtrl) GetTags(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")

	tags, err := sc.tagSer.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the tags"})
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(tags)
}

func (sc *tagCtrl) AddTag(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")

	var tag entity.TagClient
	err := json.NewDecoder(request.Body).Decode(&tag)
	if err != nil {
		log.Printf("Failed to decode input tag name in request body")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the tag name"})
		return
	}

	newTag, err := sc.tagSer.Create(tag.Name)
	if err != nil {
		log.Printf("Failed to create new tag")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error adding new tag"})
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(newTag)
}
