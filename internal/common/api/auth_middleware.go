package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mbobrovskyi/chat-management-go/internal/common/domain/http_error"
	"github.com/mbobrovskyi/chat-management-go/internal/common/domain/session"
	"github.com/mbobrovskyi/chat-management-go/internal/common/domain/user"
	"github.com/mbobrovskyi/chat-management-go/internal/infrastructure/server"
	"strings"
)

const (
	authTokenQueryParam = "token"
	headerAuthorization = "Authorization"
	bearerTokenType     = "Bearer"
)

type AuthMiddleware struct {
	userContract user.Contract
}

func (m *AuthMiddleware) Handle(ctx *fiber.Ctx) error {
	var (
		token string
		err   error
	)

	token, err = m.getTokenFromQuery(ctx)
	if err != nil {
		return err
	}

	if len(token) == 0 {
		token, err = m.getTokenFromHeader(ctx)
		if err != nil {
			return err
		}
	}

	if token == "" {
		return http_error.NewUnauthorizedError("invalid token")
	}

	user, err := m.userContract.GetCurrentUser(token)
	if err != nil {
		return err
	}

	if user == nil {
		return http_error.NewUnauthorizedError("session not found")
	}

	ctx.Context().SetUserValue("token", token)
	ctx.Context().SetUserValue("session", session.NewSession(user))

	return ctx.Next()
}

func (m *AuthMiddleware) getTokenFromQuery(ctx *fiber.Ctx) (string, error) {
	return ctx.Query(authTokenQueryParam), nil
}

func (m *AuthMiddleware) getTokenFromHeader(ctx *fiber.Ctx) (string, error) {
	authHeader := ctx.Get(headerAuthorization)
	if len(authHeader) == 0 {
		return "", http_error.NewUnauthorizedError("invalid token")
	}

	tokenParts := strings.Split(authHeader, " ")

	if len(tokenParts) < 2 {
		return "", http_error.NewUnauthorizedError("invalid token")
	}

	tokenPrefix := strings.ToLower(tokenParts[0])
	if strings.ToLower(tokenPrefix) != strings.ToLower(bearerTokenType) {
		return "", http_error.NewUnauthorizedError("invalid token type")
	}

	authToken := tokenParts[1]

	return authToken, nil
}

func NewAuthMiddleware(userContract user.Contract) server.Middleware {
	return &AuthMiddleware{
		userContract: userContract,
	}
}
