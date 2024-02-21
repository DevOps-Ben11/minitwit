package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/DevOps-Ben11/minitwit/api/model"
	"github.com/DevOps-Ben11/minitwit/api/repository"
	"github.com/DevOps-Ben11/minitwit/api/util"
)

type Register struct {
	repo repository.Repository
}

type RegisterSimulator struct {
	Username string
	Email    string
	PWD      string
}

func CreateRegisterHandler(repo repository.Repository) *Register {
	return &Register{repo: repo}
}

func (h *Register) RegisterHandler(w http.ResponseWriter, r *http.Request) {
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
		user, _ := h.repo.GetUser(vals.Get("username"))

		if !vals.Has("username") || len(vals.Get("username")) == 0 {
			s := "You have to enter a username"
			error = &s
		} else if !vals.Has("email") || !strings.Contains(vals.Get("email"), "@") {
			s := "You have to enter a valid email address"
			error = &s
		} else if !vals.Has("password") || len(vals.Get("password")) == 0 {
			s := "You have to enter a password"
			error = &s
		} else if vals.Get("password") != vals.Get("password2") {
			s := "The two passwords do not match"
			error = &s
		} else if user != nil {
			s := "The username is already taken"
			error = &s
		} else {
			err := h.repo.InsertUser(vals.Get("username"), vals.Get("email"), vals.Get("password"))

			if err != nil {
				log.Println("Error creating user:", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			util.PushFlashMessage(w, r, "You were successfully registered and can login now")
			http.Redirect(w, r, util.UrlFor("/login", ""), http.StatusFound)
			return
		}
	}

	data := model.Template{
		Error:   error,
		Request: model.RenderRequest{Endpoint: "register"},
		Flashes: util.GetFlashedMessages(w, r),
	}

	t, err := template.New("layout.html").Funcs(util.GetFuncMap()).ParseFiles("../web/templates/layout.html", "../web/templates/register.html")

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

/** <------------- SIMULATOR HANDLER -------------> **/
func (h *Register) RegisterSimHandler(w http.ResponseWriter, r *http.Request) {
	// TO DO: Implement latest counter ** update_latest_sim(r) **

	var body RegisterSimulator
	json.NewDecoder(r.Body).Decode(&body)

	username := body.Username
	email := body.Email
	pwd := body.PWD

	var error *string = nil

	if username == "" {
		s := "You have to enter a username"
		error = &s
	} else if email == "" || !strings.Contains(email, "@") {
		s := "You have to enter a valid email address"
		error = &s
	} else if pwd == "" {
		s := "You have to enter a password"
		error = &s
	}

	fmt.Println(username, email, pwd)
	fmt.Println(*error)

	// else if user != nil {
	// 	s := "The username is already taken"
	// 	error = &s
	// } else {
	// 	err := h.repo.InsertUser(vals.Get("username"), vals.Get("email"), vals.Get("password"))

	// 	if err != nil {
	// 		log.Println("Error creating user:", err)
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 		return
	// 	}

	// 	util.PushFlashMessage(w, r, "You were successfully registered and can login now")
	// 	http.Redirect(w, r, util.UrlFor("/login", ""), http.StatusFound)
	// 	return
	// }
}
