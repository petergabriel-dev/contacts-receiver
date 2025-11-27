package contact

import "context"

type Service interface {
	ListContacts(ctx context.Context) error
}

type svc struct {
}

func NewService() Service {
	return &svc{}
}

func (s *svc) ListContacts(ctx context.Context) error {
	return nil
}
