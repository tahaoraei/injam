package validator

type Repository interface {
	IsPhoneNumberUnique(number string) (bool, error)
}

type Validator struct {
	repo Repository
}

func New(repo Repository) Validator {
	return Validator{repo: repo}
}
