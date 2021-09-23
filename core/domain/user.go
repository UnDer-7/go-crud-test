package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id"`
	Email      string             `bson:"email"`
	Name       string             `bson:"name"`
	GivenName  string             `bson:"given_name"`
	FamilyName string             `bson:"family_name"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  *time.Time         `bson:"updated_at"`
}
