package web

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type chiRouter struct {
	chiDispatcher *chi.Mux
}

func NewChiRouter() Router {
	return &chiRouter{
		chiDispatcher: chi.NewRouter(),
	}
}

func (c *chiRouter) Get(uri string, f func(response http.ResponseWriter, request *http.Request)) {
	c.chiDispatcher.Get(uri, f)
}

func (c *chiRouter) Post(uri string, f func(response http.ResponseWriter, request *http.Request)) {
	c.chiDispatcher.Post(uri, f)
}

func (c *chiRouter) Serve(port string) {
	log.Println("Chi HTTP server running on port %v\n", port)
	http.ListenAndServe(port, c.chiDispatcher)
}
