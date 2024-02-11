package main

import (
	"fmt"
	"log"
	"net/http"

	"html/template"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// configuration
const DATABASE = "./tmp/minitwit.db"
const PER_PAGE = 30
const DEBUG = true
const SECRET_KEY = "development key"

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
	r.HandleFunc("/{username}/follow", s.userFollowHanlder)
	r.HandleFunc("/{username}/unfollow", s.userUnfollowHandler)
	r.HandleFunc("/{username}", s.userHandler)
	r.HandleFunc("/add_message", s.addMessageHandler)
	r.HandleFunc("/register", s.registerHandler)
	r.HandleFunc("/logout", s.logoutHandler)
	r.HandleFunc("/", s.timelineHandler)

	// https://stackoverflow.com/questions/43601359/how-do-i-serve-css-and-js-in-go
	// https://medium.com/ducktypd/serving-static-files-with-golang-or-gorilla-mux-b6bf8fa2e5e
	// for css stuff
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Open browser on localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func (s *Server) timelineHandler(w http.ResponseWriter, r *http.Request) {
	// used for testing the connection to the database.
	var response Follower
	// https://gowebexamples.com/sessions/ for later
	s.db.Raw("SELECT * FROM follower WHERE who_id = ?;", 191).Scan(&response)
	log.Println("query: ", response)
	msg := fmt.Sprintf("%v", response)
	w.Write([]byte(msg))
}

func (s *Server) logoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World\n This is the Logout"))
}
func (s *Server) registerHandler(w http.ResponseWriter, r *http.Request) {
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
	w.Write([]byte("Hello World\n This is the user timeline"))
}
func (s *Server) userUnfollowHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World\n This is the user unfollow"))
}
func (s *Server) userFollowHanlder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World\n This is the user follow"))
}
func (s *Server) publicHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("layout.html").Funcs(s.funcMap).ParseFiles("gotemplates/layout.html", "gotemplates/timeline.html")
	if err != nil {
		log.Println("Error creating template:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var messages []RenderMessage
	err = s.db.Raw(
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

	if err = t.Execute(w, data); err != nil {
		log.Println("Error rendering frontend:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func UrlFor(path string, arg string) string {
	return ""
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
	return "16 Nov 22:24"
}

type RenderMessage struct {
	Message
	User
}

type RenderRequest struct {
	Endpoint string
}

type Data struct {
	User     *User
	Messages []RenderMessage
	Request  RenderRequest
}
