package api

import (
	"github.com/DevOps-Ben11/minitwit/backend/util"
	"html/template"
	"log"
	"net/http"

	"github.com/DevOps-Ben11/minitwit/backend/model"
)

func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	user_id, ok := s.GetCurrentUser(r)

	// If the user is authenticated, redirect to the home page
	if ok || user_id != nil {
		http.Redirect(w, r, UrlFor("timeline", ""), http.StatusFound)
		return
	}

	var error *string = nil

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		vals := r.PostForm
		user, ok := s.userRepo.GetUser(vals.Get("username"))

		if !ok || user == nil {
			s := "Invalid username"
			error = &s
		} else if !util.CheckPassword(vals.Get("password"), user.Pw_hash) {
			s := "Invalid password"
			error = &s
		} else {
			session, err := s.GetStore().Get(r, "auth")

			if err != nil {
				log.Println("Error getting session:", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			session.Values["user"] = user.User_id
			err = session.Save(r, w)

			if err != nil {
				log.Println("Error logging in:", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			s.PushFlashMessage(w, r, "You were logged in")
			http.Redirect(w, r, UrlFor("timeline", ""), http.StatusFound)
		}
	}

	t, err := template.New("layout.html").Funcs(s.GetFuncMap()).ParseFiles("../web/templates/layout.html", "../web/templates/login.html")

	if err != nil {
		log.Println("Error creating template:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := model.Template{
		Request: model.RenderRequest{Endpoint: "login"},
		Error:   error,
		Flashes: s.GetFlashedMessages(w, r),
	}

	if err = t.Execute(w, data); err != nil {
		log.Println("Error rendering frontend:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
