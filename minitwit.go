package main

import (
	"fmt"
	"log"
	"net/http"

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
	db *gorm.DB
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

	s := Server{db: db}
	r.HandleFunc("/public", s.publicHandler)
	r.HandleFunc("/{username}/follow", s.userFollowHanlder)
	r.HandleFunc("/{username}/unfollow", s.userUnfollowHandler)
	r.HandleFunc("/{username}", s.userHandler)
	r.HandleFunc("/login", s.loginHandler)
	r.HandleFunc("/add_message", s.addMessageHandler)
	r.HandleFunc("/register", s.registerHandler)
	r.HandleFunc("/logout", s.logoutHandler)
	r.HandleFunc("/", s.timelineHandler)

	log.Println("Open browser on localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func (s *Server) timelineHandler(w http.ResponseWriter, r *http.Request) {
	// used for testing the connection to the database.
	var response Follower
	s.db.Raw("SELECT * FROM follower WHERE who_id = ?;", 191).Scan(&response)
	log.Println("query: ", response)
	msg := fmt.Sprintf("%v", response)
	w.Write([]byte(msg))
}

func (s *Server) logoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World\n This is the timeline"))
}
func (s *Server) registerHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World\n This is the timeline"))
}
func (s *Server) addMessageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World\n This is the timeline"))
}
func (s *Server) loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World\n This is the timeline"))
}
func (s *Server) userHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World\n This is the timeline"))
}
func (s *Server) userUnfollowHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World\n This is the timeline"))
}
func (s *Server) userFollowHanlder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World\n This is the timeline"))
}
func (s *Server) publicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World\n This is the timeline"))
}
