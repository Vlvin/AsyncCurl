package simhttp

import (
	"net/http"
	"strings"
)

func HttpRequest(method string, url string, request_body string, client ...http.Client) (*http.Response, error) {
	request_body_reader := strings.NewReader(request_body)
	r, err := http.NewRequest(method, url, request_body_reader)
	if err != nil {
		return nil, err
	}
	if len(client) > 0 {
		return client[0].Do(r)
	} else {
		return http.DefaultClient.Do(r)
	}
}
