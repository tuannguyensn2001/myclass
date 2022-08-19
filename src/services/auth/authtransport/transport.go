package authtransport

type iService interface {
}

type httpTransport struct {
	service iService
}

func NewHttpTransport(service iService) *httpTransport {
	return &httpTransport{service: service}
}
