package driver

import (
	"context"
	"my-tracking-list-backend/core/domain"
)

type UserService interface {
	SaveUser(ctx context.Context, user domain.User) (domain.User, error)
	FindByEmail(ctx context.Context, email string) (domain.User, error)
	UserExistes(ctx context.Context, email string) (bool, error)
}
