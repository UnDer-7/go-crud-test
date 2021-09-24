package driven

import (
	"context"
	"my-tracking-list-backend/core/domain"
)

type UserRepository interface {
	Persist(ctx context.Context, user domain.User) (domain.User, error)
	GetByEmail(ctx context.Context, email string) (domain.User, error)
	ExistesByEmail(ctx context.Context, email string) (bool, error)
}
