package driven

import (
	"context"
	"my-tracking-list-backend/core/domain"
)

//go:generate mockgen -source oauth_handler.go -destination ../../../test_helpers/mock/oauth_handler_mock.go -package=mocks
type OauthHandler interface {
	DecodeGoogleToken(ctx context.Context, token string) (domain.GoogleToken, error)
}
