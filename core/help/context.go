package help

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my-tracking-list-backend/core/app_error"
	"strings"
)

const keyCurrentUserEmail = "github.com/UnDer-7/my-tracking-list-backend/help_context_CURRENT_USER_EMAIL"

func SetCurrentUserEmail(ctx *gin.Context, email string) error {
	if ctx.Value(keyCurrentUserEmail) != nil {
		return app_error.ThrowInternalServerError(
			fmt.Sprintf("Context ja possui a keyCurrentUserEmail: %s", keyCurrentUserEmail),
			nil)
	}

	ctx.Set(keyCurrentUserEmail, strings.TrimSpace(email))
	return nil
}

func GetCurrentUserEmail(ctx *gin.Context) (string, error) {
	value := ctx.Value(keyCurrentUserEmail)
	if value == nil {
		return "", app_error.ThrowInternalServerError(
			fmt.Sprintf("Context nao possui chave %s", keyCurrentUserEmail),
			nil,
		)
	}

	email, ok := value.(string)
	if !ok {
		return "", app_error.ThrowInternalServerError(
			fmt.Sprintf("Valor retornado do context com a chave %s nao eh do tipo string", keyCurrentUserEmail),
			nil,
		)
	}

	if strings.TrimSpace(email) == "" {
		return "", app_error.ThrowInternalServerError(
			fmt.Sprintf("Valor retornado do context com a chave %s eh uma string vazia", keyCurrentUserEmail),
			nil,
		)
	}

	return email, nil
}
