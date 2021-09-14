package driver

import "my-tracking-list-backend/core/domain"

type UserService interface {
	SaveUser(user domain.User) (domain.User, error)
	FindById(id int) (domain.User, error)
}
