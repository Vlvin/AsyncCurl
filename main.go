package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Task struct {
	Url    string `json:"url"`
	Method string `json:"method"`
	Data   string `json:"data"`
}

func http_request(method string, url string, request_body string) (*http.Response, error) {
	request_body_reader := strings.NewReader(strings.ToUpper(request_body))
	r, err := http.NewRequest(method, url, request_body_reader)
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(r)
}

func await_task(task Task, w http.ResponseWriter) {
	r, err := http_request(task.Method, task.Url, task.Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response_body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(response_body)
	defer r.Body.Close()
}

func task_route(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	task := new(Task)
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	await_task(*task, w)

}

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

const PORT int = 8080

func main() {
	fmt.Println("Hello Go Servers")
	server := NewSimpleHTTPServer(
		fmt.Sprintf(":%d", PORT),
		MkRoute("/", task_route),
	)
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("SERVER ERROR: %s", err.Error())
	}
}
