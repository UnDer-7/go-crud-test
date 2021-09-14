package driven

import "my-tracking-list-backend/core/domain"

type OauthHandler interface {
	GoogleLogin(token string) (domain.GoogleToken, error)
}
