package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DevOps-Ben11/minitwit/backend/util"
)

type Login struct {
	Username string `json:"username"`
	PWD      string `json:"password"`
}

func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var body Login
	json.NewDecoder(r.Body).Decode(&body)

	username := body.Username
	pwd := body.PWD

	user, ok := s.userRepo.GetUser(username)
	var errorStr *string = nil

	if !ok || user == nil {
		s := "Invalid username"
		errorStr = &s
	} else if !util.CheckPassword(pwd, user.Pw_hash) {
		s := "Invalid password"
		errorStr = &s
	} else {
		session, err := s.GetStore().Get(r, "auth")

		if err != nil {
			log.Println("Error getting session:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		session.Values["user"] = user.User_id
		err = session.Save(r, w)

		if err != nil {
			log.Println("Error logging in:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
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
		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := s.GetStore().Get(r, "auth")

	if err != nil {
		log.Println("Error getting session", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for k := range session.Values {
		delete(session.Values, k)
	}
	session.Options.MaxAge = -1

	err = session.Save(r, w)

	if err != nil {
		log.Println("Error loggin out:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
