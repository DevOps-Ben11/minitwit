package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/DevOps-Ben11/minitwit/backend/model"
	"github.com/gorilla/mux"
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

func (s *Server) FollowGetSimHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	profilName := vars["username"]
	profil, ok := s.userRepo.GetUser(profilName)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	no := r.URL.Query().Get("no")
	numUsrs := 100
	if no != "" {
		num, err := strconv.Atoi(no)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		numUsrs = num
	}

	usernames, err := s.userRepo.GetUsersFollowing(profil.User_id, numUsrs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ret := map[string][]string{}
	ret["follows"] = usernames
	RetJson(w, ret)
}

func (s *Server) FollowPostSimHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	user, ok := s.userRepo.GetUser(username)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	type tmp struct {
		Follow   string `json:"follow"`
		Unfollow string `json:"unfollow"`
	}
	var body tmp
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if body.Follow != "" {
		follows, ok := s.userRepo.GetUser(body.Follow)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err := s.userRepo.SetFollow(user.User_id, follows.User_id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		return
	} else if body.Unfollow != "" {
		unfollows, ok := s.userRepo.GetUser(body.Unfollow)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err := s.userRepo.SetUnfollow(user.User_id, unfollows.User_id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		return

	}

}
