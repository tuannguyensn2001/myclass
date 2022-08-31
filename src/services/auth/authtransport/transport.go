package authtransport

import (
	"context"
	"myclass/src/services/auth/authstruct"
)

type iService interface {
	Register(ctx context.Context, input authstruct.RegisterInput) error
}

type httpTransport struct {
	service iService
}

func NewHttpTransport(service iService) *httpTransport {
	return &httpTransport{service: service}
}
