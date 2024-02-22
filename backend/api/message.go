package api

import (
	"github.com/DevOps-Ben11/minitwit/backend/model"
	"log"
	"net/http"
)

func (s *Server) AddMessageHandler(user *model.User, w http.ResponseWriter, r *http.Request) {
	// Register a new message from a user

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	vals := r.PostForm

	// early return if empty form
	if vals.Get("text") == "" {
		http.Redirect(w, r, UrlFor("timeline", ""), http.StatusFound)

		return
	}

	err = s.msgRepo.AddMessage(user, vals.Get("text"))

	if err != nil {
		log.Println("Error creating message: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	s.PushFlashMessage(w, r, "Your message was recorded")
	http.Redirect(w, r, UrlFor("timeline", ""), http.StatusFound)
	return
}
