package authorizationservice

type Repository interface {
	GetPermission(pubUserID int, subUserID int) (bool, error)
	AddPermission(pubUserID int, subUserID int) (bool, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) CheckPermission(pubUserID int, subUserID int) (bool, error) {
	return s.repo.GetPermission(pubUserID, subUserID)
}

func (s Service) AddPermission(pubUserID int, subUserID int) (bool, error) {
	return s.repo.AddPermission(pubUserID, subUserID)
}
