package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"my-tracking-list-backend/core/app_error"
	"my-tracking-list-backend/core/domain"
	"my-tracking-list-backend/core/ports/driven"
	"time"
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
	user.CreatedAt = time.Now()
	user.UpdatedAt = nil

	_, err := r.database.Collection(UserCollectionName).InsertOne(context.Background(), &user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r UserRepositoryImpl) GetByEmail(email string) (domain.User, error) {
	var user domain.User
	err := r.database.
		Collection(UserCollectionName).
		FindOne(context.Background(), bson.M{"email": email}).
		Decode(&user)

	if err != nil && errors.Is(err, mongo.ErrNoDocuments) {
		return domain.User{}, app_error.ThrowNotFoundError(
			"Usuario nao encontrado",
			"Usuario com email informado nao foi cadastrado",
			err,
		)
	} else if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
