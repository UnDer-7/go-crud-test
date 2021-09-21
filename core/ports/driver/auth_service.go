package driver

import "my-tracking-list-backend/core/domain"

type AuthService interface {
	SignIn(token string) (domain.User, error)
}
