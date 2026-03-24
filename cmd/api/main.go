package main

import (
	"log"
	"net/http"

	"github.com/fatmalabidi/buzzfuzz/internal/api"
	"github.com/fatmalabidi/buzzfuzz/internal/handlers"
)

func main() {
	server := &handlers.Server{}
	mux := http.NewServeMux()
	api.HandlerFromMux(server, mux)
	log.Print("server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
