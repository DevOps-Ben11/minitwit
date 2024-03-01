package api

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/DevOps-Ben11/minitwit/backend/model"
)

type RegisterSimulator struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	PWD      string `json:"pwd"`
}

func (s *Server) RegisterHandler(w http.ResponseWriter, r *http.Request) {
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
		user, _ := s.userRepo.GetUser(vals.Get("username"))

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
			err := s.userRepo.InsertUser(vals.Get("username"), vals.Get("email"), vals.Get("password"))

			if err != nil {
				log.Println("Error creating user:", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			s.PushFlashMessage(w, r, "You were successfully registered and can login now")
			http.Redirect(w, r, UrlFor("/login", ""), http.StatusFound)
			return
		}
	}

	data := model.Template{
		Error:   error,
		Request: model.RenderRequest{Endpoint: "register"},
		Flashes: s.GetFlashedMessages(w, r),
	}

	t, err := template.New("layout.html").Funcs(s.GetFuncMap()).ParseFiles("../web/templates/layout.html", "../web/templates/register.html")

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
func (s *Server) RegisterSimHandler(w http.ResponseWriter, r *http.Request) {
	// start := time.Now()
	var body RegisterSimulator
	json.NewDecoder(r.Body).Decode(&body)

	username := body.Username
	email := body.Email
	pwd := body.PWD

	var errorStr *string = nil

	if username == "" {
		s := "You have to enter a username"
		errorStr = &s
	} else if email == "" || !strings.Contains(email, "@") {
		s := "You have to enter a valid email address"
		errorStr = &s
	} else if pwd == "" {
		s := "You have to enter a password"
		errorStr = &s
	} else if _, ok := s.userRepo.GetUser(username); ok {
		s := "The username is already taken"
		errorStr = &s
	} else {
		// t := time.Now()
		err := s.userRepo.InsertUser(username, email, pwd)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
	// elapsed := time.Since(start)
	// log.Printf("RegisterSimHandler took %s", elapsed)

	if errorStr != nil {
		t := ErrReturn{Status: http.StatusBadRequest, ErrorMsg: *errorStr}
		m, err := json.Marshal(t)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(m)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
