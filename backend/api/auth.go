package api

import (
	"context"
	"fmt"
	"github.com/DevOps-Ben11/minitwit/backend/model"
	"log"
	"net/http"
	"strings"
)

type AuthHandlerFunc func(user *model.User, w http.ResponseWriter, r *http.Request)

func (s *Server) protect(next AuthHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := s.GetCurrentUser(r)
		if !ok || user == nil {
			http.Redirect(w, r, UrlFor("public_timeline", ""), http.StatusFound)
			return
		}
		next(user, w, r)
	}
}

func (s *Server) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/static/") {
			next.ServeHTTP(w, r)
			return
		}

		log.Println(fmt.Sprintf("Authenticating request: %s %s\n", r.Method, r.URL))

		_, ok := s.GetCurrentUser(r)
		session, err := s.store.Get(r, "auth")
		if err != nil {
			log.Println("Error getting session:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		u, ok := session.Values["user"]
		if !ok {
			next.ServeHTTP(w, r)
			return
		}

		userId, ok := u.(uint)
		if !ok {
			next.ServeHTTP(w, r)
			return
		}

		user, ok := s.userRepo.GetUserById(userId)
		if !ok {
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), UserKey, user)
		newR := r.WithContext(ctx)

		// Continue with the request
		next.ServeHTTP(w, newR)
	})
}

// AuthCtxKey As per https://pkg.go.dev/context#WithValue
type AuthCtxKey string

var UserKey = AuthCtxKey("user")
