package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/fatmalabidi/buzzfuzz/internal/api"
	"github.com/fatmalabidi/buzzfuzz/internal/fizzbuzz"
	"github.com/fatmalabidi/buzzfuzz/internal/stats"
)

var _ api.ServerInterface = new(Server)

type Server struct {
	fizzBuzzService fizzbuzz.Service
	statsService    stats.Service
}

func NewServer(fizzBuzzService fizzbuzz.Service, statsService stats.Service) *Server {
	return &Server{
		fizzBuzzService: fizzBuzzService,
		statsService:    statsService,
	}
}
func writeError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(api.Error{Code: code, Message: message})
}
