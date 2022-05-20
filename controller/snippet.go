package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rupinjairaj/snippet/entity"
	"github.com/rupinjairaj/snippet/errors"
	"github.com/rupinjairaj/snippet/service"
)

type SnippetController interface {
	GetSnippets(response http.ResponseWriter, request *http.Request)
	AddSnippet(response http.ResponseWriter, request *http.Request)
}

type snippetCtrl struct {
	snippetSer service.SnippetService
}

func NewSnippetController(snippetServ service.SnippetService) SnippetController {
	return &snippetCtrl{
		snippetSer: snippetServ,
	}
}

func (sc *snippetCtrl) GetSnippets(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")
	var tag entity.TagClient
	err := json.NewDecoder(request.Body).Decode(&tag)
	if err != nil {
		log.Printf("Failed to decode input tag name in request body")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the tag name"})
		return
	}

	snippets, err := sc.snippetSer.FindByTag(tag.Name)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the snippets"})
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(snippets)
}

func (sc *snippetCtrl) AddSnippet(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")

	var snippetClient entity.SnippetClient
	err := json.NewDecoder(request.Body).Decode(&snippetClient)
	if err != nil {
		log.Printf("Failed to decode input client snippet")
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error decoding the payload"})
		return
	}

	snippetFirestore, err := sc.snippetSer.Save(&snippetClient)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		log.Printf("Failed to store snippet")
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Failed to store snippet"})
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(snippetFirestore)
}
