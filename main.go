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

func await_task(task Task, w http.ResponseWriter) {
	request_body := strings.NewReader(strings.ToUpper(task.Data))
	r, err := http.NewRequest(task.Method, task.Url, request_body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	response_body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(response_body)
}

func tesk_route(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	task := new(Task)
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	await_task(*task, w)

}

const port int = 8080

func main() {
	fmt.Println("Hello Go Servers")
	mux := http.NewServeMux()
	mux.HandleFunc("/", tesk_route)
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("SERVER ERROR: %s", err.Error())
	}
}
