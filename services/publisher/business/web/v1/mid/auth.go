package mid

import (
	"context"
	"net/http"

	"github.com/vikaskumar1187/publisher_saasv2/services/publisher/business/web/v1/auth"
	"github.com/vikaskumar1187/publisher_saasv2/services/publisher/foundation/web"
)

// Authenticate validates a JWT from the `Authorization` header.
func Authenticate(a *auth.Auth) web.Middleware {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			claims, err := a.Authenticate(ctx, r.Header.Get("authorization"))
			if err != nil {
				return auth.NewAuthenticationError("authenticate: failed: %s", err)
			}

			ctx = auth.SetClaims(ctx, claims)

			return handler(ctx, w, r)
		}

		return h
	}

	return m
}

// Authorize validates that an authenticated user has at least one role from a
// specified list. This method constructs the actual function that is used.
func Authorize(a *auth.Auth) web.Middleware {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			claims := auth.GetClaims(ctx)

			if err := a.Authorize(ctx, claims); err != nil {
				return auth.NewAuthorizationError("authorize: you are not authorized for that action, claims[%v] : %s", err)
			}

			return handler(ctx, w, r)
		}

		return h
	}

	return m
}
