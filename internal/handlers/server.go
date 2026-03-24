package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/fatmalabidi/buzzfuzz/internal/api"
	"github.com/fatmalabidi/buzzfuzz/internal/fizzbuzz"
)

var _ api.ServerInterface = new(Server)

type Server struct {
	fizzBuzzService fizzbuzz.FizzBuzzService
}

func NewServer(fizzBuzzService fizzbuzz.FizzBuzzService) *Server {
	return &Server{
		fizzBuzzService: fizzBuzzService,
	}
}
func writeError(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(api.Error{Code: http.StatusBadRequest, Message: message})
}
