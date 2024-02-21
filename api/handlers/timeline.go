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
	db repository.DB
}

func CreateTimelineHandler(db repository.DB) *Timeline {
	return &Timeline{db: db}
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
	user, _ := h.db.GetUserById(user_id)
	fmt.Println(user)
	var messages []model.RenderMessage
	messages, ok := h.db.GetUserTimeline(user.User_id)
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
