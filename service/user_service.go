package service

import (
	"crud-test/app_error"
	"crud-test/model"
	"crud-test/repositories"
)

type UserService struct {
	Repository repositories.IUserRepository
}

func (service *UserService) SaveUser(user model.User) (model.User, error) {
	userSaved, err := service.Repository.SaveUser(user)
	if err != nil {
		return model.User{}, app_error.NewInternalServerError(
			"erro ao salvar no repository",
			err,
		)
	}

	return userSaved, nil
}
