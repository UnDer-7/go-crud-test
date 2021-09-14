package driven

import "crud-test/core/domain"

type OauthHandler interface {
	GoogleLogin(token string) (domain.GoogleToken, error)
}
