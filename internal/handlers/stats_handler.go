package handlers

import (
	"encoding/json"
	"net/http"
)

// GetStats implements [api.ServerInterface].
func (s *Server) GetStats(w http.ResponseWriter, r *http.Request) {
	stats, err := s.statsService.GetMostFrequent()
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stats)
}
