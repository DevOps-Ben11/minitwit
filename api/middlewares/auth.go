package mw

import (
	"fmt"
	"net/http"

	"github.com/DevOps-Ben11/minitwit/api/util"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Authenticating request: %s %s\n", r.Method, r.URL)

		_, ok := util.GetCurrentUser(r)
		if !ok {
			http.Redirect(w, r, util.UrlFor("public_timeline", ""), http.StatusFound)
			return
		}

		// Continue with the request
		next.ServeHTTP(w, r)
	})
}
