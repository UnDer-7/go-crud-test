package driver

import (
	"context"
	"my-tracking-list-backend/core/domain"
)

type AuthService interface {
	SignIn(ctx context.Context, token string) (domain.User, error)
}
