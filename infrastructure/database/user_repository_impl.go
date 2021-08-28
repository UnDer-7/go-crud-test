package database

import (
	"crud-test/core/app_error"
	"crud-test/core/domain"
)

type UserRepositoryImpl struct {
	Database map[int]domain.User
}

func (repository UserRepositoryImpl) Persist(user domain.User) (domain.User, error) {
	nextId := len(repository.Database) + 1

	userSaved := domain.User{
		Id:       nextId,
		Email:    user.Email,
		Password: user.Password,
	}

	repository.Database[nextId] = userSaved

	return userSaved, nil
}

func (repository UserRepositoryImpl) GetById(id int) (domain.User, error) {
	user := repository.Database[id]

	if user.Email == "" {
		return domain.User{}, app_error.NewNotFoundError("user not found!!", nil)
	}
	return user, nil
}