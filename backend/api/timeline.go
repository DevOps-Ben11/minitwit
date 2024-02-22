package api

import (
	"github.com/DevOps-Ben11/minitwit/backend/util"
	"html/template"
	"log"
	"net/http"

	"github.com/DevOps-Ben11/minitwit/backend/model"
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
	var messages []model.RenderMessage
	err := s.db.Raw(
		"select message.*, user.* from message, user where message.flagged = 0 and message.author_id = user.user_id order by message.pub_date desc limit ?",
		util.PER_PAGE).Scan(&messages).Error

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

	s.RenderTimeline(w, data)
}
