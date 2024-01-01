package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mbobrovskyi/chat-management-go/pkg/domain/baseerror"
	"github.com/mbobrovskyi/chat-management-go/pkg/domain/entities/abstracts"
	"github.com/mbobrovskyi/chat-management-go/pkg/infrastructure/httpserver"
	"strings"
)

const (
	authTokenQueryParam = "token"
	headerAuthorization = "Authorization"
	bearerTokenType     = "Bearer"
)

type AuthMiddleware struct {
	userContract abstracts.UserClient
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
		return baseerror.NewUnauthorizedError("invalid token")
	}

	user, err := m.userContract.GetCurrent(token)
	if err != nil {
		return err
	}

	if user == nil {
		return baseerror.NewUnauthorizedError("user session not found")
	}

	ctx.Context().SetUserValue("token", token)
	ctx.Context().SetUserValue("user", user)

	return ctx.Next()
}

func (m *AuthMiddleware) getTokenFromQuery(ctx *fiber.Ctx) (string, error) {
	return ctx.Query(authTokenQueryParam), nil
}

func (m *AuthMiddleware) getTokenFromHeader(ctx *fiber.Ctx) (string, error) {
	authHeader := ctx.Get(headerAuthorization)
	if len(authHeader) == 0 {
		return "", baseerror.NewUnauthorizedError("invalid token")
	}

	tokenParts := strings.Split(authHeader, " ")

	if len(tokenParts) < 2 {
		return "", baseerror.NewUnauthorizedError("invalid token")
	}

	tokenPrefix := strings.ToLower(tokenParts[0])
	if strings.ToLower(tokenPrefix) != strings.ToLower(bearerTokenType) {
		return "", baseerror.NewUnauthorizedError("invalid token type")
	}

	authToken := tokenParts[1]

	return authToken, nil
}

func NewAuthMiddleware(userContract abstracts.UserClient) httpserver.Middleware {
	return &AuthMiddleware{
		userContract: userContract,
	}
}
