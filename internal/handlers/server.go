package handlers

import (
	"net/http"

	"github.com/fatmalabidi/buzzfuzz/internal/api"
)

var _ api.ServerInterface = new(Server)

type Server struct{}

// GetStats implements [api.ServerInterface].
func (s *Server) GetStats(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message":"ok"}`))
}

func (s *Server) GetSequencesFizzbuzz(w http.ResponseWriter, r *http.Request, params api.GetSequencesFizzbuzzParams) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"ok"}`))
}
