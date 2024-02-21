package handler

import (
	"log"
	"net/http"
	"text/template"

	"github.com/DevOps-Ben11/minitwit/api/model"
	"github.com/DevOps-Ben11/minitwit/api/repository"
	"github.com/DevOps-Ben11/minitwit/api/util"
)

type Login struct {
	repo repository.Repository
}

func CreateLoginHandler(repo repository.Repository) *Login {
	return &Login{repo: repo}
}

func (h *Login) LoginHandler(w http.ResponseWriter, r *http.Request) {
	user_id, ok := util.GetCurrentUser(r)

	// If the user is authenticated, redirect to the home page
	if ok || user_id != nil {
		http.Redirect(w, r, util.UrlFor("timeline", ""), http.StatusFound)
		return
	}

	var error *string = nil

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		vals := r.PostForm
		user, ok := h.repo.GetUser(vals.Get("username"))

		if !ok || user == nil {
			s := "Invalid username"
			error = &s
		} else if !util.CheckPassword(vals.Get("password"), user.Pw_hash) {
			s := "Invalid password"
			error = &s
		} else {
			session, err := util.GetStore().Get(r, "auth")

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

			util.PushFlashMessage(w, r, "You were logged in")
			http.Redirect(w, r, util.UrlFor("timeline", ""), http.StatusFound)
		}
	}

	t, err := template.New("layout.html").Funcs(util.GetFuncMap()).ParseFiles("../web/templates/layout.html", "../web/templates/login.html")

	if err != nil {
		log.Println("Error creating template:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := model.Template{
		Request: model.RenderRequest{Endpoint: "login"},
		Error:   error,
		Flashes: util.GetFlashedMessages(w, r),
	}

	if err = t.Execute(w, data); err != nil {
		log.Println("Error rendering frontend:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
