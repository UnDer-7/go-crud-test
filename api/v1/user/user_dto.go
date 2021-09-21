package user

import (
	"my-tracking-list-backend/core/domain"
	"time"
)

type responseUser struct {
	ID         string     `json:"id"`
	Email      string     `json:"email"`
	Name       string     `json:"name"`
	GivenName  string     `json:"givenName"`
	FamilyName string     `json:"familyName"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
}

func userToResponseUser(user domain.User) responseUser {
	return responseUser{
		ID:         user.ID.Hex(),
		Email:      user.Email,
		Name:       user.Name,
		GivenName:  user.GivenName,
		FamilyName: user.FamilyName,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}
}
