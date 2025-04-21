package gourl

import (
	"io"
	"iobound/simhttp"
	"net/http"
)

type Task struct {
	Url    string `json:"url"`
	Method string `json:"method"`
	Data   string `json:"data"`
}

func AwaitTask(task Task, w http.ResponseWriter) {
	r, err := simhttp.HttpRequest(task.Method, task.Url, task.Data)
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
