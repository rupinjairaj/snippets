package web

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type chiRouter struct {
	chiDispatcher *chi.Mux
}

func NewChiRouter() Router {
	router := chi.NewRouter()
	router.Use(
		cors.Handler(cors.Options{
			AllowOriginFunc: AllowOriginFunc,
			AllowedMethods:  []string{"GET", "POST"},
			AllowedHeaders:  []string{"*"},
		}))
	return &chiRouter{
		chiDispatcher: router,
	}
}

func (c *chiRouter) Get(uri string, f func(response http.ResponseWriter, request *http.Request)) {
	c.chiDispatcher.Get(uri, f)
}

func (c *chiRouter) Post(uri string, f func(response http.ResponseWriter, request *http.Request)) {
	c.chiDispatcher.Post(uri, f)
}

func AllowOriginFunc(request *http.Request, origin string) bool {
	return true
}

func (c *chiRouter) Serve(port string) {
	log.Printf("Chi HTTP server running on port %v\n", port)
	http.ListenAndServe(port, c.chiDispatcher)
}
