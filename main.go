package main

import (
	"encoding/json"
	"fmt"
	"iobound/gourl"
	"iobound/simhttp"
	"net/http"
)

func task_route(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	task := gourl.Task{}
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	gourl.AwaitTask(task, w)

}

const PORT int = 8080

func main() {
	fmt.Println("Hello Go Servers")
	server := simhttp.NewSimpleHTTPServer(
		fmt.Sprintf(":%d", PORT),
		simhttp.MkRoute("/", task_route),
	)
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("SERVER ERROR: %s", err.Error())
	}
}
