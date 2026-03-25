package stats

//go:generate mockgen -source=service.go -destination=../../mocks/stats/mock_service.go -package=statsmocks
type Service interface {
	GetMostFrequent() (*RequestStat, error)
	Record(int1, int2, limit int, str1, str2 string)
}

type service struct {
	store *Store
}

func NewService(store *Store) Service {
	return &service{
		store: store,
	}
}

func (s *service) Record(int1, int2, limit int, str1, str2 string) {
	s.store.Increment(int1, int2, limit, str1, str2)
}

func (s *service) GetMostFrequent() (*RequestStat, error) {
	return s.store.GetMostFrequent()

}
