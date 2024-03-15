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

		s.PushFlashMessage(w, r, "You were logged in")
		http.Redirect(w, r, UrlFor("timeline", ""), http.StatusFound)
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
	err = session.Save(r, w)
	if err != nil {
		log.Println("Error loggin out:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	s.PushFlashMessage(w, r, "You were logged out")
	http.Redirect(w, r, UrlFor("public_timeline", ""), http.StatusFound)
}
