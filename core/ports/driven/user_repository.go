package driven

import "crud-test/core/domain"

type UserRepository interface {
	Persist(user domain.User) (domain.User, error)
	GetById(id int) (domain.User, error)
}