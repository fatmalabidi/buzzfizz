package stats_test

import (
	"fmt"
	"testing"

	"github.com/fatmalabidi/buzzfizz/internal/services/stats"
	"github.com/stretchr/testify/assert"
)

func TestNewStore(t *testing.T) {
	store := stats.NewStore()
	assert.NotNil(t, store, "NewStore returned nil")
	assert.NotNil(t, store.Stats, "stats map not initialized")
}

func TestIncrement(t *testing.T) {
	store := stats.NewStore()
	store.Increment(1, 2, 10, "fizz", "buzz")
	assert.Equal(t, 1, len(store.Stats), fmt.Sprintf("expected 1 stat, got %d", len(store.Stats)))

	key := stats.BuildKey(1, 2, 10, "fizz", "buzz")
	stat, exists := store.Stats[key]
	assert.True(t, exists, "stat not found in store")
	assert.Equal(t, 1, stat.Hits, fmt.Sprintf("expected Hits=1, got %d", stat.Hits))
}

func TestIncrementDuplicate(t *testing.T) {
	store := stats.NewStore()
	store.Increment(1, 2, 10, "fizz", "buzz")
	store.Increment(1, 2, 10, "fizz", "buzz")
	assert.Equal(t, 1, len(store.Stats), "expected 1 stat")

	key := stats.BuildKey(1, 2, 10, "fizz", "buzz")
	assert.Equal(t, 2, store.Stats[key].Hits, "expected Hits=2")
}

func TestGetMostFrequentStore(t *testing.T) {
	store := stats.NewStore()
	store.Increment(1, 2, 10, "fizz", "buzz")
	store.Increment(1, 2, 10, "fizz", "buzz")
	store.Increment(3, 4, 20, "foo", "bar")

	stat, err := store.GetMostFrequent()
	assert.NoError(t, err, "unexpected error")
	assert.Equal(t, 2, stat.Hits, "expected most frequent Hits=2")
}

func TestGetMostFrequentEmpty(t *testing.T) {
	store := stats.NewStore()
	_, err := store.GetMostFrequent()
	assert.Error(t, err, "expected error for empty store")
}
