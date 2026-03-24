package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fatmalabidi/buzzfizz/internal/api"
)

// GenerateFizzBuzz handles HTTP requests to generate a FizzBuzz sequence.
func (s *Server) GenerateFizzBuzz(w http.ResponseWriter, r *http.Request, params api.GenerateFizzBuzzParams) {
	err := validateFizzbuzzParams(params)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	result := s.fizzBuzzService.Generate(params.Int1, params.Int2, params.Limit, params.Str1, params.Str2)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		writeError(w, http.StatusInternalServerError, "error encoding response")
		return
	}
	s.statsService.Record(params.Int1, params.Int2, params.Limit, params.Str1, params.Str2)
}

func validateFizzbuzzParams(params api.GenerateFizzBuzzParams) error {
	if params.Int1 <= 0 || params.Int2 <= 0 {
		return fmt.Errorf("int1 and int2 should be positive")
	}

	if len(params.Str1) == 0 || len(params.Str2) == 0 {
		return fmt.Errorf("str1 and str2 should not be empty")
	}

	if params.Limit <= 0 {
		return fmt.Errorf("limit should be positive")
	}

	return nil
}
