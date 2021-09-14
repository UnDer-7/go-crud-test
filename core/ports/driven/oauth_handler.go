package driven

import "my-tracking-list-backend/core/domain"

type OauthHandler interface {
	DecodeGoogleToken(token string) (domain.GoogleToken, error)
}
