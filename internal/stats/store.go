package stats

import (
	"errors"
	"strconv"
	"sync"
)

type Store struct {
	mux   sync.RWMutex
	stats map[string]*RequestStat
}

type RequestStat struct {
	Int1, Int2, Limit int
	Str1, Str2        string
	Hits              int
}

func NewStore() *Store {
	return &Store{
		stats: make(map[string]*RequestStat),
	}
}

func (s *Store) GetMostFrequent() (*RequestStat, error) {
	s.mux.RLock()
	defer s.mux.RUnlock()

	if len(s.stats) == 0 {
		return nil, errors.New("no stats was found")
	}

	var maxHits int
	var mostFrequent *RequestStat

	for _, stat := range s.stats {
		if stat.Hits > maxHits {
			mostFrequent = stat
			maxHits = stat.Hits
		}
	}
	return mostFrequent, nil
}

func buildKey(int1, int2, limit int, str1, str2 string) string {
	return strconv.Itoa(int1) + "-" + strconv.Itoa(int2) + "-" + strconv.Itoa(limit) + "-" + str1 + "-" + str2
}

func (s *Store) Increment(int1, int2, limit int, str1, str2 string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	key := buildKey(int1, int2, limit, str1, str2)

	if stat, ok := s.stats[key]; ok {
		stat.Hits++
		return
	}

	s.stats[key] = &RequestStat{
		Int1:  int1,
		Int2:  int2,
		Limit: limit,
		Str1:  str1,
		Str2:  str2,
		Hits:  1,
	}
}
