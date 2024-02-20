package server

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) InitRoutes() error {

	s.Get("/latest", s.LatestHandler)
	return nil
}

type Handler func(vars map[string]string, r *http.Request) (status int, value any)

func (s *Server) Get(route string, handler Handler) {
	f := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		status, value := handler(vars, r)

		returnValue, err := json.Marshal(value)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		w.Write(returnValue)
	}

	s.r.HandleFunc(route, f).Methods("GET")
}

func (s *Server) Post(route string, handler Handler) {
	f := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		status, value := handler(vars, r)

		returnValue, err := json.Marshal(value)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		w.Write(returnValue)
	}

	s.r.HandleFunc(route, f).Methods("POST")
}

func DecodeBody(body io.ReadCloser, v any) error {
	return json.NewDecoder(body).Decode(v)
}

func OkResponse(value any) (int, any) {
	return http.StatusOK, value
}
