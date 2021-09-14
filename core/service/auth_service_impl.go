package service

import (
	"fmt"
	"my-tracking-list-backend/core/ports/driven"
)

type AuthServiceImpl struct {
	oauth driven.OauthHandler
}

func NewAuthService(oauth driven.OauthHandler) *AuthServiceImpl {
	return &AuthServiceImpl{oauth: oauth}
}

func (s AuthServiceImpl) Login(token string) error {
	tokenGoogle, err := s.oauth.DecodeGoogleToken(token)
	if err != nil {
		return err
	}

	fmt.Printf("Token Incoming: %v\n", tokenGoogle.Email)
	// Validar na base de dados
	return nil
}
