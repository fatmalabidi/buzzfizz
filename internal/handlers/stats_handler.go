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
	if err := json.NewEncoder(w).Encode(stats); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
