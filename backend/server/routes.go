package server

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) InitRoutes() error {

	s.Get("/test/{name}", s.TestHandler)
	s.Post("/test", s.TestPostHandler)
	return nil
}

type GetHandler func(vars map[string]string) (status int, value any)
type PostHandler func(vars map[string]string, body io.ReadCloser) (status int, value any)

func (s *Server) Get(route string, handler GetHandler) {
	f := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		status, value := handler(vars)

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

func (s *Server) Post(route string, handler PostHandler) {
	f := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		status, value := handler(vars, r.Body)

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
