package src

import (
	"LFGbackend/types"
	"context"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

var sessionContextKey = &contextKey{"session_token"}

type contextKey struct {
	name string
}

var sessions = map[string]*types.PostSession{}

// Middleware decodes the share session cookie and packs the session into context
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token, err := r.Cookie("session_token")
		// Allow authed users in
		if err == nil {
			_, ok := sessions[token.Value]
			if ok {
				next.ServeHTTP(w, r)
				return
			}
		}

		if err != nil {
			log.Println(err)
		}

		// Create a new random session token
		// we use the "github.com/google/uuid" library to generate UUIDs
		sessionToken := uuid.NewString()
		expiresAt := time.Now().Add(72 * time.Hour)

		http.SetCookie(w, &http.Cookie{
			Name:    "session_token",
			Value:   sessionToken,
			Expires: expiresAt,
		})

		// put it in context
		ctx := context.WithValue(
			r.Context(),
			sessionContextKey,
			types.PostSession{
				ClientId: sessionToken,
			},
		)

		// and call the next with our new context
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

// SessionContext finds the user from the context. REQUIRES Middleware to have run.
func SessionContext(ctx context.Context) *types.PostSession {
	raw, _ := ctx.Value(sessionContextKey).(*types.PostSession)
	return raw
}
