package driver

import "my-tracking-list-backend/core/domain"

type AuthService interface {
	Create(toke string) (domain.User, error)
}
