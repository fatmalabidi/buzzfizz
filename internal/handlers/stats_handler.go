package handlers

import "net/http"

// GetStats implements [api.ServerInterface].
func (s *Server) GetStats(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message":"ok"}`))
}
