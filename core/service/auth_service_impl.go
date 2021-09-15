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

func (s AuthServiceImpl) Create(toke string) (domain.User, error) {
	tokenGoogle, err := s.oauth.DecodeGoogleToken(toke)
	if err != nil {
		return domain.User{}, err
	}

	usr, err := s.userService.SaveUser(domain.User{
		Email:      tokenGoogle.Email,
		Name:       tokenGoogle.Name,
		GivenName:  tokenGoogle.GivenName,
		FamilyName: tokenGoogle.FamilyName,
	})
	if err != nil {
		return domain.User{}, err
	}
	return usr, nil
}
