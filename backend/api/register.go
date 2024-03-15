package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type RegisterSimulator struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	PWD      string `json:"password"`
}

func (s *Server) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	user_id, ok := s.GetCurrentUser(r)

	fmt.Println(user_id, ok)
	// If the user is authenticated, we don't want to register a new user
	if ok || user_id != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var body RegisterSimulator
	json.NewDecoder(r.Body).Decode(&body)

	username := body.Username
	email := body.Email
	pwd := body.PWD

	var errorStr *string = nil

	if username == "" {
		s := "You have to enter a username"
		errorStr = &s
	} else if email == "" || !strings.Contains(email, "@") {
		s := "You have to enter a valid email address"
		errorStr = &s
	} else if pwd == "" {
		s := "You have to enter a password"
		errorStr = &s
	} else if _, ok := s.userRepo.GetUser(username); ok {
		s := "The username is already taken"
		errorStr = &s
	} else {
		err := s.userRepo.InsertUser(username, email, pwd)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	if errorStr != nil {
		t := ErrReturn{Status: http.StatusBadRequest, ErrorMsg: *errorStr}
		m, err := json.Marshal(t)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(m)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

/** <------------- SIMULATOR HANDLER -------------> **/
func (s *Server) RegisterSimHandler(w http.ResponseWriter, r *http.Request) {
	var body RegisterSimulator
	json.NewDecoder(r.Body).Decode(&body)

	username := body.Username
	email := body.Email
	pwd := body.PWD

	var errorStr *string = nil

	if username == "" {
		s := "You have to enter a username"
		errorStr = &s
	} else if email == "" || !strings.Contains(email, "@") {
		s := "You have to enter a valid email address"
		errorStr = &s
	} else if pwd == "" {
		s := "You have to enter a password"
		errorStr = &s
	} else if _, ok := s.userRepo.GetUser(username); ok {
		s := "The username is already taken"
		errorStr = &s
	} else {

		err := s.userRepo.InsertUser(username, email, pwd)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	if errorStr != nil {
		t := ErrReturn{Status: http.StatusBadRequest, ErrorMsg: *errorStr}
		m, err := json.Marshal(t)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(m)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
