package web

import (
	"net/http"
)

type Router interface {
	Get(uri string, f func(response http.ResponseWriter, request *http.Request))
	Post(uri string, f func(response http.ResponseWriter, request *http.Request))
	Serve(port string)
}
