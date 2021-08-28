package service

import (
	"crud-test/core/domain"
	"crud-test/core/ports/driven"
	"fmt"
)

type UserServiceImpl struct {
	Repository driven.UserRepository
}

func (service UserServiceImpl) SaveUser(user domain.User) (domain.User, error) {
	if user.Id != 0 {
		// Throw error
	}

	userSave, err := service.Repository.Persist(user)
	if err != nil {
		fmt.Print(err)
		return domain.User{}, err
	}
	return userSave, nil
}

func (service UserServiceImpl) FindById(id int) (domain.User, error) {
	userFound, err := service.Repository.GetById(id)
	if err != nil {
		fmt.Print(err)
		return domain.User{}, err
	}
	return userFound, nil
}