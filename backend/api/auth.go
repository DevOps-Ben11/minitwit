package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/DevOps-Ben11/minitwit/backend/model"
)

type AuthHandlerFunc func(user *model.User, w http.ResponseWriter, r *http.Request)

func (s *Server) protect(next AuthHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := s.GetCurrentUser(r)

		if !ok || user == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next(user, w, r)
	}
}

func (s *Server) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Authenticating request: %s %s\n", r.Method, r.URL)

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

func (s *Server) simProtect(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fromSim := r.Header.Get("Authorization")
		if fromSim != "Basic c2ltdWxhdG9yOnN1cGVyX3NhZmUh" {
			e := ErrReturn{Status: http.StatusForbidden, ErrorMsg: "You are not authorized to use this resource!"}
			t, err := json.Marshal(e)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			w.Write(t)
			return
		}
		next(w, r)
	}
}
