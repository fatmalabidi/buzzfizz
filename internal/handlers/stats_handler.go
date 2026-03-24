package handlers

import (
	"encoding/json"
	"net/http"
)

// GetMostFrequentRequest implements [api.ServerInterface].
func (s *Server) GetMostFrequentRequest(w http.ResponseWriter, r *http.Request) {
	stats, err := s.statsService.GetMostFrequent()
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	// write header only after encoding succeeds — or accept the limitation
	if err := json.NewEncoder(w).Encode(stats); err != nil {
		// log the error; can't change status at this point
		return
	}
}
