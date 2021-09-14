package repository

import (
	"my-tracking-list-backend/core/domain"
	"my-tracking-list-backend/core/ports/driven"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	database *gorm.DB
}

func NewUserRepository(database *gorm.DB) driven.UserRepository {
	return &UserRepositoryImpl{
		database: database,
	}
}

func (repository UserRepositoryImpl) Persist(user domain.User) (domain.User, error) {
	err := repository.database.Create(&user).Error
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (repository UserRepositoryImpl) GetById(id int) (domain.User, error) {
	return domain.User{}, nil
}
