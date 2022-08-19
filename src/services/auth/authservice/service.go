package authservice

type iRepository interface {
}

type service struct {
	repository iRepository
}

func NewService(repository iRepository) *service {
	return &service{repository: repository}
}
