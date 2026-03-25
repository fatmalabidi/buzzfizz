package stats

//go:generate mockgen -destination=mocks/mock_service.go -package=mocks github.com/fatmalabidi/buzzfizz/internal/stats Service
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
