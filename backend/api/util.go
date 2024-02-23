package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// configuration

func UrlFor(path string, arg string) string {
	switch path {
	case "add_message":
		return "/add_message"
	case "user_timeline":
		return fmt.Sprintf("/%s", arg)
	case "unfollow_user":
		return fmt.Sprintf("/%s/unfollow", arg)
	case "follow_user":
		return fmt.Sprintf("/%s/follow", arg)
	case "timeline":
		return "/"
	case "public_timeline":
		return "/public"
	case "register":
		return "/register"
	case "login":
		return "/login"
	case "logout":
		return "/logout"
	default:
		return "/"
	}
}

func (s *Server) LatestHandler(w http.ResponseWriter, r *http.Request) {
	var id int
	kv, err := s.GetKeyVal("latest")
	if err != nil {
		id = -1
	} else {
		i, err := strconv.Atoi(kv.Value)
		if err != nil {
			id = -1
		} else {
			id = i
		}
	}
	type tmp struct {
		Latest int `json:"latest"`
	}

	t := tmp{Latest: id}
	m, err := json.Marshal(t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(m)
}

func (s *Server) LatestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		latestStr := r.URL.Query().Get("latest")
		if latestStr != "" {
			log.Println("Updating latest value:", latestStr)
			s.SetKeyVal("latest", latestStr)
		}
		next.ServeHTTP(w, r)
	})
}

type ErrReturn struct {
	Status   int    `json:"status"`
	ErrorMsg string `json:"error_msg"`
}

func RetJson(w http.ResponseWriter, value any) {
	t, err := json.Marshal(value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(t)
	return
}
