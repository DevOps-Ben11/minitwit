package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DevOps-Ben11/minitwit/backend/model"
)

type TestMsg struct {
	Msg string `json:"msg"`
}

func (s *Server) TestHandler(vars map[string]string, r *http.Request) (status int, value any) {
	name := vars["name"]
	msg := fmt.Sprintf("Hello %s!", name)
	s.db.Create(&model.Example{
		Msg: msg,
	})

	return OkResponse(TestMsg{Msg: msg})
}

func (s *Server) TestPostHandler(vars map[string]string, r *http.Request) (status int, value any) {
	var data TestMsg
	DecodeBody(r.Body, &data)
	log.Println("Got message!")
	log.Println(data.Msg)
	return OkResponse(nil)
}

func (s *Server) LatestHandler(vars map[string]string, r *http.Request) (status int, value any) {

	return 404, nil
}

func (s *Server) RegisterHandler(vars map[string]string, r *http.Request) (status int, value any) {

	// // We have to deal with the user here, so we check that data entered are right (We don't have to this for the simulation)
	// var error *string = nil
	// vals := r.PostForm
	// if !vals.Has("username") || len(vals.Get("username")) == 0 {
	// 	s := "You have to enter a username"
	// 	error = &s
	// } else if !vals.Has("email") || !strings.Contains(vals.Get("email"), "@") {
	// 	s := "You have to enter a valid email address"
	// 	error = &s
	// } else if !vals.Has("password") || len(vals.Get("password")) == 0 {
	// 	s := "You have to enter a password"
	// 	error = &s
	// } else if vals.Get("password") != vals.Get("password2") {
	// 	s := "The two passwords do not match"
	// 	error = &s
	// }

	// if error != nil {
	// 	// send flash and stop the process because data is wrong, TODO Matteo
	// }

	// // Now we know that data are right, we can ask to register in the db
	// registerInDBResult := repository.registerDB(s, vals.Get("username"), vals.Get("email"), vals.Get("password"))
	// if registerInDBResult != nil {
	// 	// send flash and stop the process because data is wrong, TODO Matteo
	// }

	// // Now know that user has been registered and we should update/return template
	// utill.PushFlashMessage(w, r, "You were successfully registered and can login now")
	// http.Redirect(w, r, UrlFor("login", ""), http.StatusFound)

	return 404, nil
}

func (s *Server) RegisterSimHandler(vars map[string]string, r *http.Request) (status int, value any) {

	return 404, nil
}

func (s *Server) MsgsHandler(vars map[string]string, r *http.Request) (status int, value any) {

	return 404, nil
}

func (s *Server) GetUserMsgsHandler(vars map[string]string, r *http.Request) (status int, value any) {
	username := vars["username"]

	return 404, username
}

func (s *Server) PostUserMsgsHandler(vars map[string]string, r *http.Request) (status int, value any) {
	username := vars["username"]

	return 404, username
}

func (s *Server) GetUserFollowsHandler(vars map[string]string, r *http.Request) (status int, value any) {
	username := vars["username"]

	return 404, username
}

func (s *Server) PostUserFollowsHandler(vars map[string]string, r *http.Request) (status int, value any) {
	username := vars["username"]

	return 404, username
}
