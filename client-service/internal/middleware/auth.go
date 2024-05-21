package middleware

import (
	"client-service/internal/service/auth"
	"context"
	"net/http"
	"strings"
)

// HTTP middleware setting a value on the request context
func Auth(authService *auth.AuthService) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authorizationHeader := r.Header.Get("Authorization")
			idToken := strings.TrimLeft(authorizationHeader, "Bearer ")

			token, err := authService.VerifyToken(r.Context(), idToken)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			uid := token.UID

			// create new context from `r` request context, and assign key `"user"`
			// to value of `"123"`
			//ctx := context.WithValue(r.Context(), "user", "123")

			ctx := context.WithValue(r.Context(), "uid", uid)

			// call the next handler in the chain, passing the response writer and
			// the updated request object with the new context value.
			//
			// note: context.Context values are nested, so any previously set
			// values will be accessible as well, and the new `"user"` key
			// will be accessible from this point forward.
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
