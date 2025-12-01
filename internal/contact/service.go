package contact

import (
	repo "contacts/internal/adapters/sqlite/sqlc"
	"context"
)

type Service interface {
	ListContacts(ctx context.Context) ([]repo.Contact, error)
	GetContactByID(ctx context.Context, id int64) (repo.Contact, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) ListContacts(ctx context.Context) ([]repo.Contact, error) {
	return s.repo.ListContacts(ctx)
}

func (s *svc) GetContactByID(ctx context.Context, id int64) (repo.Contact, error) {
	return s.repo.GetContactByID(ctx, id)
}
