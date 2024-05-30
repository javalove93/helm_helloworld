package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	env := os.Getenv("COLOR")
	fmt.Fprintf(w, "Hello, world from %s!\n", env)
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server started on port 80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
