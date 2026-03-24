package main

import (
	"log"
	"net/http"

	"github.com/fatmalabidi/buzzfizz/internal/api"
	"github.com/fatmalabidi/buzzfizz/internal/fizzbuzz"
	"github.com/fatmalabidi/buzzfizz/internal/handlers"
	"github.com/fatmalabidi/buzzfizz/internal/stats"
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

	mux.Handle("/api/", http.StripPrefix("/api/", http.FileServer(http.Dir("./api"))))
	mux.Handle("/docs", http.RedirectHandler("/docs/", http.StatusMovedPermanently))
	mux.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(http.Dir("./web/swagger"))))

	api.HandlerFromMux(server, mux)
	log.Println("server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
