package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/DevOps-Ben11/minitwit/api/model"
	"github.com/DevOps-Ben11/minitwit/api/repository"
	"github.com/DevOps-Ben11/minitwit/api/util"
)

type Timeline struct {
	repo repository.Repository
}

func CreateTimelineHandler(repo repository.Repository) *Timeline {
	return &Timeline{repo: repo}
}

func (h *Timeline) RenderTimeline(w http.ResponseWriter, data model.Template) {
	t, err := template.New("layout.html").Funcs(util.GetFuncMap()).ParseFiles("../web/templates/layout.html", "../web/templates/timeline.html")

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

func (h *Timeline) TimelineHandler(w http.ResponseWriter, r *http.Request) {
	user_id, _ := util.GetCurrentUser(r)
	user, _ := h.repo.GetUserById(user_id)
	fmt.Println(user)
	var messages []model.RenderMessage
	messages, ok := h.repo.GetUserTimeline(user.User_id)
	fmt.Println(ok)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := model.Template{
		User:     user,
		Messages: messages,
		Request:  model.RenderRequest{Endpoint: "timeline"},
		Flashes:  util.GetFlashedMessages(w, r),
	}

	h.RenderTimeline(w, data)
}
