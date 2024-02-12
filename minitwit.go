package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"html/template"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// configuration
const DATABASE = "./tmp/minitwit.db"
const PER_PAGE = 30
const DEBUG = true
const SECRET_KEY = "development key"

var store = sessions.NewCookieStore([]byte(SECRET_KEY))

type Server struct {
	db      *gorm.DB
	funcMap template.FuncMap
}
type User struct {
	User_id  uint
	Username string
	Email    string
	Pw_hash  string
}

type Message struct {
	Message_id uint
	Author_id  uint
	Text       string
	Pub_date   int64
	Flagged    bool
}

type Follower struct {
	Who_id  uint
	Whom_id uint
}

func main() {
	r := mux.NewRouter()
	db, err := gorm.Open(sqlite.Open("./tmp/minitwit.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	funcMap := template.FuncMap{
		"UrlFor":             UrlFor,
		"GetFlashedMessages": GetFlashedMessages,
		"Gravatar":           Gravatar,
		"Datetimeformat":     Datetimeformat,
	}

	s := Server{db: db, funcMap: funcMap}
	r.HandleFunc("/public", s.publicHandler)
	r.HandleFunc("/login", s.loginHandler)
	r.HandleFunc("/add_message", s.addMessageHandler)
	r.HandleFunc("/register", s.registerHandlerGet).Methods("Get")
	r.HandleFunc("/logout", s.logoutHandler)
	r.HandleFunc("/{username}/follow", s.userFollowHanlder)
	r.HandleFunc("/{username}/unfollow", s.userUnfollowHandler)
	r.HandleFunc("/{username}", s.userHandler)
	r.HandleFunc("/", s.timelineHandler)

	// https://stackoverflow.com/questions/43601359/how-do-i-serve-css-and-js-in-go
	// https://medium.com/ducktypd/serving-static-files-with-golang-or-gorilla-mux-b6bf8fa2e5e
	// for css stuff
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Open browser on localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func (s *Server) timelineHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := s.GetCurrentUser(r)
	if !ok {
		http.Redirect(w, r, UrlFor("public_timeline", ""), http.StatusFound)
		return
	}
	var messages []RenderMessage
	err := s.db.Raw("select message.*, user.* from message, user where message.flagged = 0 and message.author_id = user.user_id and (user.user_id = ? or user.user_id in (select whom_id from follower where who_id = ?)) order by message.pub_date desc limit ?",
		user.User_id, user.User_id, PER_PAGE).Scan(&messages).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data := Data{
		User:     user,
		Messages: messages,
		Request:  RenderRequest{Endpoint: "timeline"},
	}
	s.RenderTimeline(w, data)
}

func (s *Server) logoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World\n This is the Logout"))
}
func (s *Server) registerHandlerGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World\n This is the register"))
}
func (s *Server) addMessageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World\n This is the add message"))
}

func (s *Server) loginHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("layout.html").Funcs(s.funcMap).ParseFiles("gotemplates/layout.html", "gotemplates/login.html")
	if err != nil {
		log.Println("Error creating template:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = t.Execute(w, nil); err != nil {
		log.Println("Error rendering frontend:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *Server) userHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var profile User
	username, ok := vars["username"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := s.db.Raw("select * from user where username = ?", username).Scan(&profile).Error
	if err != nil {
		log.Println("Error getting user:", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var messages []RenderMessage
	err = s.db.Raw("select message.*, user.* from message, user where user.user_id = message.author_id and user.user_id = ? order by message.pub_date desc limit ?", profile.User_id, PER_PAGE).Scan(&messages).Error
	if err != nil {
		log.Println("Error getting messages:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	followed := false
	user, ok := s.GetCurrentUser(r)
	if ok {
		s.db.Raw("select 1 from follower where follower.who_id = ? and follower.whom_id = ?", user.User_id, profile.User_id).Scan(&followed)
	}

	data := Data{
		User:     user,
		Profile:  &profile,
		Messages: messages,
		Request:  RenderRequest{Endpoint: "user_timeline"},
		Followed: followed,
	}
	s.RenderTimeline(w, data)
}
func (s *Server) userUnfollowHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World\n This is the user unfollow"))
}
func (s *Server) userFollowHanlder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World\n This is the user follow"))
}
func (s *Server) publicHandler(w http.ResponseWriter, r *http.Request) {
	var messages []RenderMessage
	err := s.db.Raw(
		"select message.*, user.* from message, user where message.flagged = 0 and message.author_id = user.user_id order by message.pub_date desc limit ?",
		PER_PAGE).Scan(&messages).Error

	if err != nil {
		log.Println("Error getting messages:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := Data{
		Request:  RenderRequest{Endpoint: "public"},
		Messages: messages,
		User:     nil,
	}

	s.RenderTimeline(w, data)
}

func (s *Server) RenderTimeline(w http.ResponseWriter, data Data) {
	t, err := template.New("layout.html").Funcs(s.funcMap).ParseFiles("gotemplates/layout.html", "gotemplates/timeline.html")
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

func UrlFor(path string, arg string) string {
	switch path {
	case "user_timeline":
		return fmt.Sprintf("/%s", arg)
	case "unfollow_user":
		return fmt.Sprintf("/%s/unfollow", arg)
	case "follow_user":
		return fmt.Sprintf("/%s/follow", arg)
	case "timeline":
		return "/"
	case "public_timeline":
		return "/public"
	case "register":
		return "/register"
	case "login":
		return "/login"
	default:
		return "/"
	}
}

func (s *Server) GetCurrentUser(r *http.Request) (user *User, ok bool) {
	session, err := store.Get(r, "auth")
	if err != nil {
		return nil, false
	}
	u, ok := session.Values["user"]
	if !ok {
		return nil, false
	}
	var user_id uint
	if user_id, ok = u.(uint); !ok {
		// Handle the case that it's not an expected type
		return nil, false
	}
	s.db.Raw("select * from user where user_id = ?", user_id).Scan(&user)
	if user == nil {
		return nil, false
	}
	return user, true
}

type FlashMessage struct {
	Message string
}

func GetFlashedMessages() []FlashMessage {
	return []FlashMessage{FlashMessage{Message: "Hello"}}
}

func Gravatar(size int, name string) string {
	return fmt.Sprintf("https://gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50?s=%d", size)
}

func Datetimeformat(date int64) string {
	t := time.Unix(date, 0)
	// Format time as "Month Date HH:MM"
	formattedTime := t.Format("January 02 2006 15:04")

	return formattedTime
}

type RenderMessage struct {
	Message
	User
}

type RenderRequest struct {
	Endpoint string
}

type Data struct {
	User *User
	// When on another users page
	Profile  *User
	Messages []RenderMessage
	Request  RenderRequest
	Followed bool
	Error    *string
}
