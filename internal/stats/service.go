package stats

type Service struct {
	store *Store
}

func NewService(store *Store) *Service {
	return &Service{
		store: store,
	}
}

func (s *Service) Record(int1, int2, limit int, str1, str2 string) {
	s.store.Increment(int1, int2, limit, str1, str2)
}

func (s *Service) GetMostFrequent()(*RequestStat, error) {
	return s.store.GetMostFrequent()

}
