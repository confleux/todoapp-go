package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	PORT = "3000"
)

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = PORT
	}

	log.Printf("Starting server on port: %s...", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatalf("Unable to start server on port: %s", port)
	}
}
