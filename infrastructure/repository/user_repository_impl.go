package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"my-tracking-list-backend/core/domain"
	"my-tracking-list-backend/core/ports/driven"
)

var UserCollectionName = "users"

type UserRepositoryImpl struct {
	database *mongo.Database
}

func NewUserRepository(database *mongo.Database) driven.UserRepository {
	return &UserRepositoryImpl{
		database: database,
	}
}

func (repository UserRepositoryImpl) Persist(user domain.User) (domain.User, error) {
	user.ID = primitive.NewObjectID()
	res, err := repository.database.Collection(UserCollectionName).InsertOne(context.Background(), user)
	if err != nil {
		return domain.User{}, err
	}
	fmt.Printf("Id gerado: %v\n", res.InsertedID)
	return user, nil
}

func (repository UserRepositoryImpl) GetById(id int) (domain.User, error) {
	return domain.User{}, nil
}
