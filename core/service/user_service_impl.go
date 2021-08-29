package service

import (
	"crud-test/core/domain"
	"crud-test/core/ports/driven"
	"crud-test/core/ports/driver"
	"fmt"
)

type UserServiceImpl struct {
	repository driven.UserRepository
}

func NewUserService(repository driven.UserRepository) driver.UserService {
	return &UserServiceImpl{
		repository: repository,
	}
}

func (service UserServiceImpl) SaveUser(user domain.User) (domain.User, error) {
	if user.Id != 0 {
		// Throw error
	}

	userSave, err := service.repository.Persist(user)
	if err != nil {
		fmt.Print(err)
		return domain.User{}, err
	}
	return userSave, nil
}

func (service UserServiceImpl) FindById(id int) (domain.User, error) {
	userFound, err := service.repository.GetById(id)
	if err != nil {
		return domain.User{}, err
	}
	return userFound, nil
}
