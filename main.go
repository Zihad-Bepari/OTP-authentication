package main

import (
	"fmt"
	"myproject/handlers"
	"net/http"
	"os"
)

func main() {

	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(handlers.IndexHandler))
	mux.Handle("/verify", http.HandlerFunc(handlers.VerifyHandler))

	port := ":8080"

	fmt.Println("Server running at http://localhost", port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		fmt.Println("Error starting the server")
		os.Exit(1)
	}
}
