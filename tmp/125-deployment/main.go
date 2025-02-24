package main

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

const (
	port = ":8081"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "Hello from Docker!! 🚀")
	if err != nil {
		return
	}
}

func main() {
	_, _ = uuid.NewUUID() // just to have and use a library
	fmt.Printf("🖥️ Starting server on `%s`\n", port)
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
