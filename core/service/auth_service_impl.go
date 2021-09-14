package service

import (
	"crud-test/core/ports/driven"
	"fmt"
)

type AuthServiceImpl struct {
	oauth driven.OauthHandler
}

func NewAuthService(oauth driven.OauthHandler) *AuthServiceImpl {
	return &AuthServiceImpl{oauth: oauth}
}

func (s AuthServiceImpl) Login(token string) error {
	tokenGoogle, err := s.oauth.GoogleLogin(token)
	if err != nil {
		return err
	}

	fmt.Printf("Token Incoming: %v\n", tokenGoogle.Email)
	// Validar na base de dados
	return nil
}
