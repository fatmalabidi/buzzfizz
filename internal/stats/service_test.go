package stats_test

import (
	"errors"
	"testing"

	"github.com/fatmalabidi/buzzfizz/internal/stats"
	"github.com/stretchr/testify/assert"
)

func TestStoreIncrement(t *testing.T) {
	testStruct := struct {
		name  string
		int1  int
		int2  int
		limit int
		str1  string
		str2  string
	}{
		name:  "increment",
		int1:  3,
		int2:  5,
		limit: 20,
		str1:  "fizz",
		str2:  "buzz",
	}

	store := stats.NewStore()
	store.Increment(testStruct.int1, testStruct.int2, testStruct.limit, testStruct.str1, testStruct.str2)
	store.Increment(testStruct.int1, testStruct.int2, testStruct.limit, testStruct.str1, testStruct.str2)

	stat, err := store.GetMostFrequent()
	assert.NoError(t, err, "expected no error got", err)
	assert.Equal(t, stat.Hits, 2)
}

func TestGetMostFrequent(t *testing.T) {
	store := stats.NewStore()
	store.Increment(3, 5, 20, "fizz", "buzz")
	store.Increment(3, 5, 20, "fizz", "buzz")
	store.Increment(7, 11, 100, "foo", "bar")

	most, err := store.GetMostFrequent()
	assert.NoError(t, err, "expected no error got", err)
	assert.Equal(t, most.Hits, 2)
	assert.Equal(t, most.Int1, 3)
}

func TestGetMostFrequent_EmptyStats(t *testing.T) {
	store := stats.NewStore()

	most, err := store.GetMostFrequent()
	assert.Error(t, err, errors.New("no stats was found"))
	assert.Nil(t, most)
}
