package main

import (
	"fmt"
	"log"
	"net/http"
	"web-confleux/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	log.Printf("Starting server on port: %s...", cfg.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), nil)
	if err != nil {
		log.Fatalf("Unable to start server on port: %s", cfg.Port)
	}
}
