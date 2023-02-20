package middleware

import (
	"context"
	"fmt"
	"github.com/dainoguchi/bstodo-api/internal/infra/auth0"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strings"
)

type AuthMiddleware interface {
	EnsureValidToken(next echo.HandlerFunc) echo.HandlerFunc
}

type authMiddleware struct {
	jwtValidator auth0.JwtValidator
}

func NewAuthMiddleware(jwtValidator auth0.JwtValidator) AuthMiddleware {
	return &authMiddleware{jwtValidator: jwtValidator}
}

// EnsureValidToken is a middleware that will check the validity of our JWT.
func (a *authMiddleware) EnsureValidToken(next echo.HandlerFunc) echo.HandlerFunc {

	// echoに書き換え
	return func(c echo.Context) error {
		// 普通にechoでjwtのトークン抜き出すコード
		authorization := c.Request().Header.Get("Authorization")

		authHeaders := strings.Split(authorization, " ")
		if len(authHeaders) != 2 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid authorization header, should format `Bearer code`, but got %s\n", authHeaders))
		}

		ctx := c.Request().Context()

		// 取得したtokenをvalidatorに渡す
		// 本当は外部から注入したい
		token, err := a.jwtValidator.ValidateToken(ctx, authHeaders[1])
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("Invalid jwt token. %s", err.Error()))
		}

		log.Printf("parsed token: %s\n", token)

		// with value の keyは空構造体が良いらしい
		ctx = context.WithValue(ctx, TokenContextKey{}, token)
		c.SetRequest(c.Request().WithContext(ctx))

		log.Println("context.value: ", c.Request().Context().Value(TokenContextKey{}))

		return next(c)
	}
}

type TokenContextKey struct{}
