package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/fatmalabidi/buzzfizz/internal/api"
	"github.com/fatmalabidi/buzzfizz/internal/services/fizzbuzz"
	"github.com/fatmalabidi/buzzfizz/internal/services/stats"
)

var _ api.ServerInterface = new(Server)

type Server struct {
	FizzBuzzService fizzbuzz.Service
	StatsService    stats.Service
}

func NewServer(fizzBuzzService fizzbuzz.Service, statsService stats.Service) *Server {
	return &Server{
		FizzBuzzService: fizzBuzzService,
		StatsService:    statsService,
	}
}
func writeError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(api.ErrorResponse{Error: message}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
