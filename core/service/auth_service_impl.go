package service

import (
	"context"
	"my-tracking-list-backend/core/domain"
	"my-tracking-list-backend/core/ports/driven"
	"my-tracking-list-backend/core/ports/driver"
)

type authService struct {
	oauth       driven.OauthHandler
	userService driver.UserService
}

func NewAuthService(oauth driven.OauthHandler, userService driver.UserService) driver.AuthService {
	return &authService{oauth: oauth, userService: userService}
}

func (s authService) SignIn(ctx context.Context, token string) (domain.User, error) {
	tokenGoogle, err := s.oauth.DecodeGoogleToken(ctx, token)
	if err != nil {
		return domain.User{}, err
	}

	exists, err := s.userService.UserExists(ctx, tokenGoogle.Email)
	if err != nil {
		return domain.User{}, err
	}

	if exists {
		userFound, err := s.userService.FindByEmail(ctx, tokenGoogle.Email)
		if err != nil {
			return domain.User{}, err
		}
		return userFound, nil
	}

	userCreated, err := s.userService.SaveUser(ctx, domain.User{
		Email:      tokenGoogle.Email,
		Name:       tokenGoogle.Name,
		GivenName:  tokenGoogle.GivenName,
		FamilyName: tokenGoogle.FamilyName,
	})
	if err != nil {
		return domain.User{}, err
	}
	return userCreated, nil
}
