package middleware

import (
	"context"
	"github.com/google/uuid"
	"net/http"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

type Client struct {
	Id string
}

// Middleware decodes the share session cookie and packs the session into context
func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			c, err := r.Cookie("clientId-cookie")
			client := &Client{Id: c.Value}

			if err == nil && c != nil {
				client.Id = c.Value
				ctx := context.WithValue(r.Context(), userCtxKey, client)
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
				return
			}

			clientId := uuid.NewString()
			client.Id = clientId
			// get the user from the database

			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, client)
			cookie := http.Cookie{
				Name:     "clientId-cookie",
				Value:    clientId,
				Path:     "/",
				MaxAge:   60 * 60 * 72,
				HttpOnly: false,
				Secure:   false,
				SameSite: http.SameSiteLaxMode,
			}
			http.SetCookie(w, &cookie)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *Client {
	raw, _ := ctx.Value(userCtxKey).(*Client)
	return raw
}
