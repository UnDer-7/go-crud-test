package database

import (
	"crud-test/core/app_error"
	"crud-test/core/domain"
	"crud-test/core/ports/driven"
	"fmt"
)

type UserRepositoryImpl struct {
	database map[int]domain.User
}

func NewUserRepository(database map[int]domain.User) driven.UserRepository {
	return &UserRepositoryImpl{
		database: database,
	}
}

func (repository UserRepositoryImpl) Persist(user domain.User) (domain.User, error) {
	nextId := len(repository.database) + 1

	userSaved := domain.User{
		Id:       nextId,
		Email:    user.Email,
		Password: user.Password,
	}

	repository.database[nextId] = userSaved

	return userSaved, nil
}

func (repository UserRepositoryImpl) GetById(id int) (domain.User, error) {
	user := repository.database[id]

	if user.Email == "" {
		return domain.User{},
			app_error.ThrowNotFoundError(
				"user not found!!",
				fmt.Sprintf("User with id %d is not in the database", id),
				nil)
	}
	return user, nil
}
