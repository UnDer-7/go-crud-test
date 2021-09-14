package driven

import "my-tracking-list-backend/core/domain"

type UserRepository interface {
	Persist(user domain.User) (domain.User, error)
	GetById(id int) (domain.User, error)
}
