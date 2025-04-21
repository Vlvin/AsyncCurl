package simhttp

import (
	"net/http"
)

type RouteHandler struct {
	route   string
	handler func(http.ResponseWriter, *http.Request)
}

func MkRoute(route string, handler func(http.ResponseWriter, *http.Request)) RouteHandler {
	return RouteHandler{
		route,
		handler,
	}
}

func NewSimpleHTTPServer(Address string, handlers ...RouteHandler) http.Server {
	mux := http.NewServeMux()
	for _, handler := range handlers {
		mux.HandleFunc(handler.route, handler.handler)
	}
	return http.Server{
		Addr:    Address,
		Handler: mux,
	}

}
