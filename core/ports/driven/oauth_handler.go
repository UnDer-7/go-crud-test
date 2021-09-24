package driven

import (
	"context"
	"my-tracking-list-backend/core/domain"
)

type OauthHandler interface {
	DecodeGoogleToken(ctx context.Context, token string) (domain.GoogleToken, error)
}
