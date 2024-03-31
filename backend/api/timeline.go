package api

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/DevOps-Ben11/minitwit/backend/model"
	"github.com/DevOps-Ben11/minitwit/backend/util"
)

func (s *Server) RenderTimeline(w http.ResponseWriter, data model.Template) {
	t, err := template.New("layout.html").Funcs(s.GetFuncMap()).ParseFiles("../web/templates/layout.html", "../web/templates/timeline.html")

	if err != nil {
		log.Println("Error creating template:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = t.Execute(w, data); err != nil {
		log.Println("Error rendering frontend:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

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
		Request:  model.RenderRequest{Endpoint: "timeline"},
		Flashes:  s.GetFlashedMessages(w, r),
	}

	s.RenderTimeline(w, data)
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
		Request:  model.RenderRequest{Endpoint: "public"},
		Messages: messages,
		User:     user,
		Flashes:  s.GetFlashedMessages(w, r),
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
		Request:  model.RenderRequest{Endpoint: "user_timeline"},
		Followed: followed,
		Flashes:  s.GetFlashedMessages(w, r),
	}
	s.RenderTimeline(w, data)
}
