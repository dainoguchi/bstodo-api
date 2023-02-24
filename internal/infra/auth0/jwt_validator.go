package auth0

import (
	"context"
	"errors"
	"fmt"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"net/url"
	"time"
)

type JwtValidator interface {
	ValidateToken(context.Context, string) (*JwtToken, error)
}

// jwt-validatorのラップ構造体
// tokenの型も独自定義型に詰め替える
type jwtValidator struct {
	auth0domain, auth0audience string
}

func NewJwtValidator(auth0domain string, auth0audience string) JwtValidator {
	return &jwtValidator{auth0domain: auth0domain, auth0audience: auth0audience}
}

// validator.ValidatedTokenに対応
type JwtToken struct {
	RegisteredClaims RegisteredClaims
}

type RegisteredClaims struct {
	Issuer    string   `json:"iss,omitempty"`
	Subject   string   `json:"sub,omitempty"`
	Audience  []string `json:"aud,omitempty"`
	Expiry    int64    `json:"exp,omitempty"`
	NotBefore int64    `json:"nbf,omitempty"`
	IssuedAt  int64    `json:"iat,omitempty"`
	ID        string   `json:"jti,omitempty"`
}

func (jv *jwtValidator) ValidateToken(ctx context.Context, tokenString string) (*JwtToken, error) {
	issuerURL, err := url.Parse("https://" + jv.auth0domain + "/")
	if err != nil {
		return nil, err
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	// parserの本体。jwtValidator.ValidateTokenにtoken渡すとauth0の公開鍵拾ってきて検証するっぽい
	vd, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{jv.auth0audience},
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		return nil, err
	}

	// ライブラリの関数使用
	rawToken, err := vd.ValidateToken(ctx, tokenString)
	if err != nil {
		return nil, err
	}

	token, ok := rawToken.(*validator.ValidatedClaims)
	if !ok {
		return nil, fmt.Errorf("%w", errors.New("validate token error, cant cast to validator.ValidatedClaims"))
	}

	// ライブラリ依存をなくす為独自定義型に詰め替える
	// 中身ほ一緒
	return &JwtToken{
		RegisteredClaims: RegisteredClaims{
			Issuer:    token.RegisteredClaims.Issuer,
			Subject:   token.RegisteredClaims.Subject,
			Audience:  token.RegisteredClaims.Audience,
			Expiry:    token.RegisteredClaims.Expiry,
			NotBefore: token.RegisteredClaims.NotBefore,
			IssuedAt:  token.RegisteredClaims.IssuedAt,
			ID:        token.RegisteredClaims.ID,
		},
	}, nil
}
