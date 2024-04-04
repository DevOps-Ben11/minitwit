package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

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
}
