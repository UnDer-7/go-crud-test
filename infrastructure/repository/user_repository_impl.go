package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"my-tracking-list-backend/core/app_error"
	"my-tracking-list-backend/core/domain"
	"my-tracking-list-backend/core/ports/driven"
	"time"
)

const UserCollectionName = "users"

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(database *mongo.Database) driven.UserRepository {
	return &userRepository{
		collection: database.Collection(UserCollectionName),
	}
}

func (r userRepository) Persist(user domain.User) (domain.User, error) {
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = nil

	_, err := r.collection.InsertOne(context.Background(), &user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r userRepository) GetByEmail(email string) (domain.User, error) {
	var user domain.User
	err := r.collection.
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

func (r userRepository) ExistesByEmail(email string) (bool, error) {
	limit := int64(1)
	total, err := r.collection.CountDocuments(context.Background(), bson.M{"email": email}, &options.CountOptions{Limit: &limit})
	if err != nil {
		return false, app_error.ThrowInternalServerError("Erro ao verificar se usuario existe por email", err)
	}

	return total == 1, nil
}
