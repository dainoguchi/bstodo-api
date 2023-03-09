package httputil

import "context"

// with value の keyは空構造体が良いらしい
type Auth0IDContextKey struct{}

func SetAuth0ID(ctx context.Context, auth0ID string) context.Context {
	return context.WithValue(ctx, Auth0IDContextKey{}, auth0ID)
}

func GetAuth0ID(ctx context.Context) string {
	v := ctx.Value(Auth0IDContextKey{})
	auth0ID, ok := v.(string)
	if !ok {
		return ""
	}

	return auth0ID
}
