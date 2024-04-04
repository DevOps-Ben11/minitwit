package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DevOps-Ben11/minitwit/backend/model"
	"github.com/DevOps-Ben11/minitwit/backend/util"
	"github.com/gorilla/mux"
)

func (s *Server) TimelineHandler(user *model.User, w http.ResponseWriter, r *http.Request) {
	var messages []model.RenderMessage
	messages, err := s.userRepo.GetUserTimeline(user.User_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error getting timeline for user", user.User_id, err)
		return
	}

	data := model.Template{
		User:     user,
		Messages: messages,
	}

	m, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(m)
}

func (s *Server) PublicTimelineHandler(w http.ResponseWriter, r *http.Request) {
	messages, err := s.msgRepo.GetPublicMessages(util.PER_PAGE)
	if err != nil {
		log.Println("Error getting messages:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, _ := s.GetCurrentUser(r)

	data := model.Template{
		Messages: messages,
		User:     user,
	}

	m, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(m)
}

func (s *Server) UserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username, ok := vars["username"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	profile, ok := s.userRepo.GetUser(username)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	messages, err := s.msgRepo.GetUserMessages(profile.User_id, util.PER_PAGE)
	if err != nil {
		log.Println("Error getting messages:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	followed := false
	user, ok := s.GetCurrentUser(r)
	if ok {
		followed = s.userRepo.GetIsFollowing(user.User_id, profile.User_id)
	}

	data := model.Template{
		User:     user,
		Profile:  profile,
		Messages: messages,
		Followed: followed,
	}

	m, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(m)
}
