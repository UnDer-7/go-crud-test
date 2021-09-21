package service

import (
	"my-tracking-list-backend/core/domain"
	"my-tracking-list-backend/core/ports/driven"
	"my-tracking-list-backend/core/ports/driver"
)

type AuthServiceImpl struct {
	oauth driven.OauthHandler
	userService driver.UserService
}

func NewAuthService(oauth driven.OauthHandler, userService driver.UserService) *AuthServiceImpl {
	return &AuthServiceImpl{oauth: oauth, userService: userService}
}

func (s AuthServiceImpl) SignIn(token string) (domain.User, error) {
	tokenGoogle, err := s.oauth.DecodeGoogleToken(token)
	if err != nil {
		return domain.User{}, err
	}

	exists, err := s.userService.UserExistes(tokenGoogle.Email)
	if err != nil {
		return domain.User{}, err
	}

	if exists {
		userFound, err := s.userService.FindByEmail(tokenGoogle.Email)
		if err != nil {
			return domain.User{}, err
		}
		return userFound, nil
	}

	userCreated, err := s.userService.SaveUser(domain.User{
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
