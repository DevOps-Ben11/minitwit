package api

import (
	"log"
	"net/http"
	"time"
)

func (s *Server) addMessageHandler(w http.ResponseWriter, r *http.Request) {
	// Register a new message from a user
	user, ok := s.GetCurrentUser(r)

	if !ok {
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	vals := r.PostForm

	// early return if empty form
	if vals.Get("text") == "" {
		http.Redirect(w, r, UrlFor("timeline", ""), http.StatusFound)

		return
	}

	err = s.db.Exec("insert into message (author_id, text, pub_date, flagged) values (?, ?, ?, 0)",
		user.User_id,
		vals.Get("text"),
		time.Now().Unix(),
	).Error

	if err != nil {
		log.Println("Error loggin in: ", err)
	}

	s.PushFlashMessage(w, r, "Your message was recorded")
	http.Redirect(w, r, UrlFor("timeline", ""), http.StatusFound)
	return
}
