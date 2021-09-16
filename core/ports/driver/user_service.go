package driver

import "my-tracking-list-backend/core/domain"

type UserService interface {
	SaveUser(user domain.User) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
	UserExistes(email string) (bool, error)
}
