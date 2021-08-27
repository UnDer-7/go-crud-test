package repositories

import (
	"crud-test/model"
)

type IUserRepository interface {
	SaveUser(user model.User) (model.User, error)
}