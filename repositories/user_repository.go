package repositories

import (
	"crud-test/app_error"
	"crud-test/model"
)

type UserRepository struct {
}

func (r *UserRepository) SaveUser(user model.User) (model.User, error) {
	//return model.User{
	//	Id:       1,
	//	Email:    user.Email,
	//	Password: user.Password,
	//}, nil
	return model.User{}, app_error.NewInternalServerError("NO DATABASE, Fool!!", nil)
}
