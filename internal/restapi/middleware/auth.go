package middleware

import (
	"context"
	"fmt"
	"github.com/dainoguchi/bstodo-api/internal/infra/auth0"
	"github.com/labstack/echo/v4"
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

func (a *authMiddleware) EnsureValidToken(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		// 普通にechoでjwtのトークン抜き出すコード
		authorization := c.Request().Header.Get("Authorization")

		authHeaders := strings.Split(authorization, " ")
		if len(authHeaders) != 2 {
			return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("Invalid authorization header, should format `Bearer code`"))
		}

		ctx := c.Request().Context()

		// jwtトークンをパースし、auth0.JwtToken型の変数を返す
		token, err := a.jwtValidator.ValidateToken(ctx, authHeaders[1])
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("Invalid jwt token. %s", err.Error()))
		}

		ctx = context.WithValue(ctx, TokenContextKey{}, token)
		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}

// with value の keyは空構造体が良いらしい
type TokenContextKey struct{}
