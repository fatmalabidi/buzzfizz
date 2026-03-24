package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fatmalabidi/buzzfuzz/internal/api"
)

func (s *Server) GetSequencesFizzbuzz(w http.ResponseWriter, r *http.Request, params api.GetSequencesFizzbuzzParams) {
	err := validateFizzbuzzParams(params)
	if err != nil {
		writeError(w, err.Error())
	}

	result := s.fizzBuzzService.Generate(params.Int1, params.Int2, params.Limit, params.Str1, params.Str2)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func validateFizzbuzzParams(params api.GetSequencesFizzbuzzParams) error {
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
