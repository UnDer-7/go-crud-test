package driven

import "my-tracking-list-backend/core/domain"

type UserRepository interface {
	Persist(user domain.User) (domain.User, error)
	GetByEmail(email string) (domain.User, error)
	ExistesByEmail(email string) (bool, error)
}
