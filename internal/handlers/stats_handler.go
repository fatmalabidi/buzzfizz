package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

// GetMostFrequentRequest implements [api.ServerInterface].
func (s *Server) GetMostFrequentRequest(w http.ResponseWriter, r *http.Request) {
	stats, err := s.statsService.GetMostFrequent()
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(stats); err != nil {
		logrus.Print("failed to encode stats", err)
		return
	}
}
