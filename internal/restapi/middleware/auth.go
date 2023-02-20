package middleware

import (
	"context"
	"fmt"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type AuthMiddleware interface {
	EnsureValidToken(next echo.HandlerFunc) echo.HandlerFunc
}

type authMiddleware struct{}

func NewAuthMiddleware() AuthMiddleware {
	return &authMiddleware{}
}

// CustomClaims contains custom data we want from the token.
type CustomClaims struct {
	Scope string `json:"scope"`
}

// Validate does nothing for this example, but we need
// it to satisfy validator.CustomClaims interface.
func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

// EnsureValidToken is a middleware that will check the validity of our JWT.
func (a *authMiddleware) EnsureValidToken(next echo.HandlerFunc) echo.HandlerFunc {
	issuerURL, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/")
	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	// parserの本体。jwtValidator.ValidateTokenにtoken渡すとauth0の公開鍵拾ってきて検証するっぽい
	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{os.Getenv("AUTH0_AUDIENCE")},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			},
		),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		log.Fatalf("Failed to set up the jwt validator")
	}

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
		token, err := jwtValidator.ValidateToken(ctx, authHeaders[1])
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
