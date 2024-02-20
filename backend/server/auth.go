package server

import (
	"net/http"
)

func (s *Server) auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ok := s.GetCurrentUser(r)

		// if !ok {
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	return
		// }

		next.ServeHTTP(w, r)
	})
}
