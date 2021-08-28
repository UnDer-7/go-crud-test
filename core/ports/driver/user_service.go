package driver

import "crud-test/core/domain"

type UserService interface {
	SaveUser(user domain.User) (domain.User, error)
	FindById(id int) (domain.User, error)
}
