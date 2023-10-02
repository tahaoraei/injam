package locationservice

type InMemoryRepository interface {
	SaveUserLocation(userID string, od string) error
	GetUserLocation(userID string) (string, error)
}

type Repository interface {
	SaveUserLocation(userID int, longitude float64, latitude float64, timestamp int64) error
	GetUserLocation(userID int) (float64, float64, int64, error)
}

type Service struct {
	inMemoryRepo InMemoryRepository
	repo         Repository
}

func New(inMemoryRepo InMemoryRepository) Service {
	return Service{inMemoryRepo: inMemoryRepo}
}
