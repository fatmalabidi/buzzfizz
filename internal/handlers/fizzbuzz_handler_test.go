package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fatmalabidi/buzzfuzz/internal/api"
)

type mockFizzBuzzService struct {
}

func (m *mockFizzBuzzService) Generate(int1, int2, limit int, str1, str2 string) []string {
	return []string{"1", "2", "Fizz"}
}

func TestGetSequencesFizzbuzz_Success(t *testing.T) {
	mockService := &mockFizzBuzzService{}
	server := &Server{fizzBuzzService: mockService}

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	params := api.GetSequencesFizzbuzzParams{Int1: 3, Int2: 5, Limit: 10, Str1: "Fizz", Str2: "Buzz"}

	server.GetSequencesFizzbuzz(w, r, params)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}
}

func TestValidateFizzbuzzParams_InvalidInt1(t *testing.T) {
	params := api.GetSequencesFizzbuzzParams{Int1: -1, Int2: 5, Limit: 10, Str1: "Fizz", Str2: "Buzz"}
	err := validateFizzbuzzParams(params)
	if err == nil {
		t.Error("expected error for negative int1")
	}
}

func TestValidateFizzbuzzParams_EmptyStr1(t *testing.T) {
	params := api.GetSequencesFizzbuzzParams{Int1: 3, Int2: 5, Limit: 10, Str1: "", Str2: "Buzz"}
	err := validateFizzbuzzParams(params)
	if err == nil {
		t.Error("expected error for empty str1")
	}
}

func TestValidateFizzbuzzParams_NegativeLimit(t *testing.T) {
	params := api.GetSequencesFizzbuzzParams{Int1: 3, Int2: 5, Limit: -5, Str1: "Fizz", Str2: "Buzz"}
	err := validateFizzbuzzParams(params)
	if err == nil {
		t.Error("expected error for negative limit")
	}
}

func TestValidateFizzbuzzParams_Valid(t *testing.T) {
	params := api.GetSequencesFizzbuzzParams{Int1: 3, Int2: 5, Limit: 10, Str1: "Fizz", Str2: "Buzz"}
	err := validateFizzbuzzParams(params)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
