package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fatmalabidi/buzzfizz/internal/handlers"
	statsmocks "github.com/fatmalabidi/buzzfizz/internal/mocks/stats"
	"github.com/fatmalabidi/buzzfizz/internal/services/stats"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetMostFrequentRequest_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := statsmocks.NewMockService(ctrl)
	mostFrequent := &stats.RequestStat{
		Int1:  2,
		Int2:  2,
		Limit: 3,
		Str1:  "test1",
		Str2:  "test2",
		Hits:  5,
	}
	mockService.EXPECT().GetMostFrequent().Return(mostFrequent, nil)

	server := &handlers.Server{StatsService: mockService}

	req := httptest.NewRequest("GET", "/stats", nil)
	w := httptest.NewRecorder()

	server.GetMostFrequentRequest(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
	mockService.EXPECT().Record(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

	var result stats.RequestStat
	assert.NoError(t, json.NewDecoder(w.Body).Decode(&result))
	assert.Equal(t, 5, result.Hits)
	assert.Equal(t, "test1", result.Str1)
	assert.Equal(t, "test2", result.Str2)
}
