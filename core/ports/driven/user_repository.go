package driven

import (
	"context"
	"my-tracking-list-backend/core/domain"
)

//go:generate mockgen -source user_repository.go -destination ../../../test_helpers/mock/user_repository_mock.go -package=mocks
type UserRepository interface {
	Persist(ctx context.Context, user domain.User) (domain.User, error)
	GetByEmail(ctx context.Context, email string) (domain.User, error)
	ExistesByEmail(ctx context.Context, email string) (bool, error)
}
