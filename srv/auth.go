package srv

import (
	"LFGbackend/types"
	"context"
	"github.com/google/uuid"
	"log"
	"net/http"
	"sync"
	"time"
)

var sessionContextKey = &contextKey{"session_token"}

type contextKey struct {
	name string
}

var mutex = sync.Mutex{}

// Middleware decodes the share session cookie and packs the session into context
func Middleware(s *Server, sessions map[string]*types.PostSession, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token, err := r.Cookie("session_token")
		// Allow authed users in
		if err == nil {
			mutex.Lock()
			_, ok := sessions[token.Value]
			mutex.Unlock()
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

		session := &types.PostSession{
			ClientId: sessionToken,
		}

		go func() {
			time.Sleep(time.Hour * 72)
			mutex.Lock()
			delete(sessions, sessionToken)
			mutex.Unlock()
			s.DeletePlayer(session)
		}()

		mutex.Lock()
		sessions[sessionToken] = session
		mutex.Unlock()

		// put it in context
		ctx := context.WithValue(
			r.Context(),
			sessionContextKey,
			session,
		)

		// and call the next with our new context
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

// GetSession  finds the user from the context. REQUIRES Middleware to have run.
func GetSession(ctx context.Context) *types.PostSession {
	raw, _ := ctx.Value(sessionContextKey).(*types.PostSession)
	return raw
}
