package api

import (
	"encoding/json"
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

	w.WriteHeader(http.StatusOK)
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

	w.WriteHeader(http.StatusOK)
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
		log.Println("FollowPostSimHandler: Error getting user:", username)
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
		log.Println("FollowPostSimHandler: Error decoding body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if body.Follow != "" {
		follows, ok := s.userRepo.GetUser(body.Follow)
		if !ok {
			log.Println("FollowPostSimHandler: Error getting user to follow:", body.Follow)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err := s.userRepo.SetFollow(user.User_id, follows.User_id)
		if err != nil {
			log.Println("FollowPostSimHandler: Error setting follow", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		return
	} else if body.Unfollow != "" {
		unfollows, ok := s.userRepo.GetUser(body.Unfollow)
		if !ok {
			log.Println("FollowPostSimHandler: Error getting user to unfollow:", body.Unfollow)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err := s.userRepo.SetUnfollow(user.User_id, unfollows.User_id)
		if err != nil {
			log.Println("FollowPostSimHandler: Error setting unfollow", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		return
	}
}
