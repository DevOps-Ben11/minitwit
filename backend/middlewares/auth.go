package mw

import (
	"fmt"
	"net/http"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Authenticating request: %s %s\n", r.Method, r.URL)

		next.ServeHTTP(w, r)
	})
}
