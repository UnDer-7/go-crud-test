package service

import "crud-test/model"

type IUserService interface {
	SaveUser(user model.User) (model.User, error)
}