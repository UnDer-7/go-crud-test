package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
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

func (r UserRepositoryImpl) Persist(user domain.User) (domain.User, error) {
	user.ID = primitive.NewObjectID()
	res, err := r.database.Collection(UserCollectionName).InsertOne(context.Background(), user)
	if err != nil {
		return domain.User{}, err
	}
	fmt.Printf("Id gerado: %v\n", res.InsertedID)
	return user, nil
}

func (r UserRepositoryImpl) GetByEmail(email string) (domain.User, error) {
	var user domain.User
	err := r.database.
		Collection(UserCollectionName).
		FindOne(context.Background(), bson.M{"email": email}).
		Decode(&user)
	if err != nil {
		return domain.User{}, err
	}

	// todo: tratar not found, ele jogar um erro :(
	return user, nil
}
