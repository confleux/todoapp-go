package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "3000"
	}

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatalf("Unable to start server on port: %s", port)
	}

	log.Printf("Starting server on port: %s", port)
}
