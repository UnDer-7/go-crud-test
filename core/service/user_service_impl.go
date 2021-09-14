package service

import (
	"fmt"
	"my-tracking-list-backend/core/app_error"
	"my-tracking-list-backend/core/domain"
	"my-tracking-list-backend/core/ports/driven"
	"my-tracking-list-backend/core/ports/driver"
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
	if !user.ID.IsZero() {
		return domain.User{}, app_error.ThrowBusinessError(
			"Um error ocorreu, tente novamente",
			"Nao pode criar usuario com Id",
		)
	}

	userSave, err := service.repository.Persist(user)
	if err != nil {
		fmt.Print(err)
		return domain.User{}, err
	}
	return userSave, nil
}

func (service UserServiceImpl) FindByEmail(email string) (domain.User, error) {
	// todo: validar email
	userFound, err := service.repository.GetByEmail(email)
	if err != nil {
		return domain.User{}, err
	}
	return userFound, nil
}
