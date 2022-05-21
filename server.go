package main

import (
	"os"

	"github.com/rupinjairaj/snippet/controller"
	"github.com/rupinjairaj/snippet/repository"
	"github.com/rupinjairaj/snippet/service"
	"github.com/rupinjairaj/snippet/web"
)

var (
	tagRepository     repository.TagRepository     = repository.NewFirestoreTagRepo()
	snippetRepository repository.SnippetRepository = repository.NewFirestoreSnippetRepo(tagRepository)
	tagService        service.TagService           = service.NewTagService(tagRepository)
	snippetService    service.SnippetService       = service.NewSnippetService(snippetRepository)
	snippetController controller.SnippetController = controller.NewSnippetController(snippetService)
	tagController     controller.TagController     = controller.NewTagController(tagService)
	httpRouter        web.Router                   = web.NewChiRouter()
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":9000"
	}

	httpRouter.Get("/snippet", snippetController.GetSnippets)
	httpRouter.Post("/snippet", snippetController.AddSnippet)
	httpRouter.Get("/tag", tagController.GetTags)
	httpRouter.Post("/tag", tagController.AddTag)
	httpRouter.Serve(port)
}