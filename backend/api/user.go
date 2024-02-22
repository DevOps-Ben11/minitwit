package api

import (
	"fmt"
	"github.com/DevOps-Ben11/minitwit/backend/model"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (s *Server) FollowHandler(user *model.User, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	profilName := vars["username"]
	profil, ok := s.userRepo.GetUser(profilName)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err := s.userRepo.SetFollow(user.User_id, profil.User_id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error following user", err)
		return
	}

	s.PushFlashMessage(w, r, fmt.Sprintf("You are now following \"%s\"", profil.Username))
	http.Redirect(w, r, UrlFor("user_timeline", profil.Username), http.StatusFound)

}

func (s *Server) UnfollowHandler(user *model.User, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	profilName := vars["username"]
	profil, ok := s.userRepo.GetUser(profilName)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err := s.userRepo.SetUnfollow(user.User_id, profil.User_id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error following user", err)
		return
	}

	s.PushFlashMessage(w, r, fmt.Sprintf("You are no longer following \"%s\"", profil.Username))
	http.Redirect(w, r, UrlFor("user_timeline", profil.Username), http.StatusFound)

}
