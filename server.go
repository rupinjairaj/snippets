package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/rupinjairaj/snippet/controller"
	"github.com/rupinjairaj/snippet/repository"
	"github.com/rupinjairaj/snippet/service"
	"github.com/rupinjairaj/snippet/web"
)

var (
	uuid              repository.UUID              = repository.NewUUIDGen()
	tagRepository     repository.TagRepository     = repository.NewFirestoreTagRepo(uuid)
	snippetRepository repository.SnippetRepository = repository.NewFirestoreSnippetRepo(tagRepository, uuid)
	tagService        service.TagService           = service.NewTagService(tagRepository)
	snippetService    service.SnippetService       = service.NewSnippetService(snippetRepository)
	snippetController controller.SnippetController = controller.NewSnippetController(snippetService)
	tagController     controller.TagController     = controller.NewTagController(tagService)
	httpRouter        web.Router                   = web.NewChiRouter()
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	httpRouter.Get("/health", func(response http.ResponseWriter, request *http.Request) {
		json.NewEncoder(response).Encode(fmt.Sprintf("Site is up and running, %v", time.Now()))
	})
	httpRouter.Get("/snippet/{tagName}", snippetController.GetSnippets)
	httpRouter.Post("/snippet", snippetController.AddSnippet)
	httpRouter.Get("/tag", tagController.GetTags)
	httpRouter.Post("/tag", tagController.AddTag)
	httpRouter.Serve(":" + port)
}
