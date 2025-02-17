package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HelloHandler(wrt http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		wrt.WriteHeader(http.StatusAccepted)
		wrt.Write([]byte("We have received your greeting"))
	default:
		http.Error(wrt, "Method not allowed",
			http.StatusMethodNotAllowed)
	}
}

func CreateItemHandler(wrt http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		var reqData struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		}
		if json.NewDecoder(req.Body).Decode(&reqData) == nil {
			fmt.Printf("Got create-item request: %+v\n", reqData)
			wrt.WriteHeader(http.StatusCreated)
		}
	default:
		http.Error(wrt, "Method not allowed",
			http.StatusMethodNotAllowed)
	}
}

func main() {
	// be aware that "/hello" matches all URLs with that prefix!
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/items", CreateItemHandler)
	http.ListenAndServe(":8080", nil)
}
