package userservice

type builder struct {
	repository iRepository
}

type iBuilder interface {
	Repository(repository iRepository) *builder
	Build() *service
}

func Init() *builder {
	return &builder{}
}

func (b *builder) Repository(repository iRepository) *builder {
	b.repository = repository
	return b
}

func (b *builder) Build() *service {
	return &service{
		repository: b.repository,
	}
}
