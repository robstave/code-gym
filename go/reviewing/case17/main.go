// Case 17: Health check. Difficulty: Medium.
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func checkHealth(client *http.Client, url string) error {
	response, err := client.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return nil
}

func main() {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer server.Close()

	fmt.Println("Healthy:", checkHealth(server.Client(), server.URL) == nil)
}
