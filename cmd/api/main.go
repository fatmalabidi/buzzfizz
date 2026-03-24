package main

import (
	"log"
	"net/http"

	"github.com/fatmalabidi/buzzfuzz/internal/api"
	"github.com/fatmalabidi/buzzfuzz/internal/fizzbuzz"
	"github.com/fatmalabidi/buzzfuzz/internal/handlers"
	"github.com/fatmalabidi/buzzfuzz/internal/stats"
)

func main() {
	fizzBuzzService := fizzbuzz.NewService()
	store := stats.NewStore()
	statsService := stats.NewService(store)

	if fizzBuzzService == nil || statsService == nil {
		log.Fatal("services must not be nil")
	}
	server := handlers.NewServer(*fizzBuzzService, *statsService)
	mux := http.NewServeMux()
	api.HandlerFromMux(server, mux)
	log.Print("server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
