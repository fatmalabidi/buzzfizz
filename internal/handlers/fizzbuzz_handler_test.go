package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fatmalabidi/buzzfizz/internal/api"
	"github.com/fatmalabidi/buzzfizz/internal/handlers"
	fizzbuzzmocks "github.com/fatmalabidi/buzzfizz/internal/mocks/fizzbuzz"
	statsmocks "github.com/fatmalabidi/buzzfizz/internal/mocks/stats"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGenerateFizzBuzz_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFizz := fizzbuzzmocks.NewMockService(ctrl)
	mockStats := statsmocks.NewMockService(ctrl)

	params := api.GenerateFizzBuzzParams{Int1: 3, Int2: 5, Limit: 100, Str1: "Fizz", Str2: "Buzz"}

	mockFizz.EXPECT().
		Generate(params.Int1, params.Int2, params.Limit, params.Str1, params.Str2).
		Return([]string{"1", "2", "Fizz"}).
		Times(1)

	mockStats.EXPECT().
		Record(params.Int1, params.Int2, params.Limit, params.Str1, params.Str2).
		Times(1)

	server := &handlers.Server{
		FizzBuzzService: mockFizz,
		StatsService:    mockStats,
	}

	w := httptest.NewRecorder()
	server.GenerateFizzBuzz(w, httptest.NewRequest(http.MethodGet, "/", nil), params)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGenerateFizzBuzz_InvalidParams(t *testing.T) {
	server := &handlers.Server{FizzBuzzService: nil, StatsService: nil}
	params := api.GenerateFizzBuzzParams{Int1: -1, Int2: 5, Limit: 100, Str1: "Fizz", Str2: "Buzz"}

	w := httptest.NewRecorder()
	server.GenerateFizzBuzz(w, httptest.NewRequest("GET", "/", nil), params)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGenerateFizzBuzz_StatsRecorded(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFizz := fizzbuzzmocks.NewMockService(ctrl)
	mockStats := statsmocks.NewMockService(ctrl)

	mockFizz.EXPECT().Generate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]string{"1", "2", "Fizz"}).Times(1)
	mockStats.EXPECT().Record(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(1)

	server := &handlers.Server{
		FizzBuzzService: mockFizz,
		StatsService:    mockStats,
	}

	params := api.GenerateFizzBuzzParams{Int1: 3, Int2: 5, Limit: 100, Str1: "Fizz", Str2: "Buzz"}
	w := httptest.NewRecorder()
	server.GenerateFizzBuzz(w, httptest.NewRequest("GET", "/", nil), params)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGenerateFizzBuzz_ValidParams_ReturnsOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFizz := fizzbuzzmocks.NewMockService(ctrl)
	mockStats := statsmocks.NewMockService(ctrl)

	params := api.GenerateFizzBuzzParams{
		Int1: 3, Int2: 5, Limit: 100, Str1: "Fizz", Str2: "Buzz",
	}

	mockFizz.EXPECT().Generate(params.Int1, params.Int2, params.Limit, params.Str1, params.Str2).
		Return([]string{"1", "2", "Fizz"}).Times(1)
	mockStats.EXPECT().Record(params.Int1, params.Int2, params.Limit, params.Str1, params.Str2).
		Times(1)

	server := &handlers.Server{
		FizzBuzzService: mockFizz,
		StatsService:    mockStats,
	}
	w := httptest.NewRecorder()

	server.GenerateFizzBuzz(w, httptest.NewRequest(http.MethodGet, "/", nil), params)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGenerateFizzBuzz_InvalidInt1_ReturnsBadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFizz := fizzbuzzmocks.NewMockService(ctrl)
	mockStats := statsmocks.NewMockService(ctrl)

	params := api.GenerateFizzBuzzParams{
		Int1: -1, Int2: 5, Limit: 100, Str1: "Fizz", Str2: "Buzz",
	}

	mockFizz.EXPECT().Generate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
	mockStats.EXPECT().Record(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

	server := &handlers.Server{FizzBuzzService: mockFizz, StatsService: mockStats}
	w := httptest.NewRecorder()

	server.GenerateFizzBuzz(w, httptest.NewRequest(http.MethodGet, "/", nil), params)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGenerateFizzBuzz_InvalidStr1_ReturnsBadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFizz := fizzbuzzmocks.NewMockService(ctrl)
	mockStats := statsmocks.NewMockService(ctrl)

	params := api.GenerateFizzBuzzParams{
		Int1: 3, Int2: 5, Limit: 100, Str1: "", Str2: "Buzz",
	}

	mockFizz.EXPECT().Generate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
	mockStats.EXPECT().Record(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

	server := &handlers.Server{FizzBuzzService: mockFizz, StatsService: mockStats}
	w := httptest.NewRecorder()

	server.GenerateFizzBuzz(w, httptest.NewRequest(http.MethodGet, "/", nil), params)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGenerateFizzBuzz_InvalidLimit_ReturnsBadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFizz := fizzbuzzmocks.NewMockService(ctrl)
	mockStats := statsmocks.NewMockService(ctrl)

	params := api.GenerateFizzBuzzParams{
		Int1: 3, Int2: 5, Limit: 0, Str1: "Fizz", Str2: "Buzz",
	}

	mockFizz.EXPECT().Generate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
	mockStats.EXPECT().Record(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

	server := &handlers.Server{FizzBuzzService: mockFizz, StatsService: mockStats}
	w := httptest.NewRecorder()

	server.GenerateFizzBuzz(w, httptest.NewRequest(http.MethodGet, "/", nil), params)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
