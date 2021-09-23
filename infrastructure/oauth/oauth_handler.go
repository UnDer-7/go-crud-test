package oauth

import (
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/api/oauth2/v1"
	"my-tracking-list-backend/core/app_error"
	"my-tracking-list-backend/core/domain"
	"my-tracking-list-backend/core/ports/driven"
	"net/http"
)

type oauthHandler struct{}

func NewOauthHandler() driven.OauthHandler {
	return &oauthHandler{}
}

func (h oauthHandler) DecodeGoogleToken(tokenStr string) (domain.GoogleToken, error) {
	oauthService, err := oauth2.New(http.DefaultClient)
	if err != nil {
		return domain.GoogleToken{}, app_error.ThrowInternalServerError(
			"Erro ao criar Oauth2 Client", err,
		)
	}

	tokenInfoCall := oauthService.Tokeninfo()
	tokenInfoCall.IdToken(tokenStr)

	if _, err := tokenInfoCall.Do(); err != nil {
		return domain.GoogleToken{}, app_error.ThrowBadRequestError(
			"Erro ao realizar validar login",
			"Token informado eh invalido",
			err,
		)
	}

	token, _, er := new(jwt.Parser).ParseUnverified(tokenStr, &domain.GoogleToken{})
	if er != nil {
		return domain.GoogleToken{}, app_error.ThrowInternalServerError(
			"Erro ao decodificar token",
			err,
		)
	}

	if googleToken, ok := token.Claims.(*domain.GoogleToken); ok {
		return *googleToken, nil
	}
	return domain.GoogleToken{}, app_error.ThrowInternalServerError(
		"token nao eh do tipo domain.GoogleToken",
		nil,
	)
}
