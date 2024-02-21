package main

import (
	"fmt"
	"log"
	"net/http"

	handler "github.com/DevOps-Ben11/minitwit/api/handlers"
	mw "github.com/DevOps-Ben11/minitwit/api/middlewares"
	"github.com/DevOps-Ben11/minitwit/api/model"
	"github.com/DevOps-Ben11/minitwit/api/repository"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const port = ":5001"

type Server struct {
	r  *mux.Router
	db *gorm.DB
}

func main() {
	s := NewServer()

	err := s.InitRoutes()
	if err != nil {
		log.Fatalln(fmt.Sprintf("Error when adding routes: %v", err))
	}

	err = s.InitDB()
	if err != nil {
		log.Fatalln(fmt.Sprintf("Error when initializing database: %v", err))
	}

	s.StartServer()
}

func NewServer() Server {
	db, err := gorm.Open(sqlite.Open("../tmp/minitwit.db"), &gorm.Config{})

	if err != nil {
		log.Fatalln("Could not open Database", err)
	}

	s := Server{
		r:  mux.NewRouter(),
		db: db,
	}

	return s
}

func (s *Server) StartServer() {
	log.Println("Starting server on port", port)
	log.Fatal(http.ListenAndServe(port, s.r))
}

func (s *Server) InitRoutes() error {
	repo := repository.CreateRepository(s.db)
	rH := handler.CreateRegisterHandler(repo)
	lH := handler.CreateLoginHandler(repo)
	tH := handler.CreateTimelineHandler(repo)

	s.r.Handle("/register", http.HandlerFunc(rH.RegisterHandler))
	s.r.Handle("/sim/register", http.HandlerFunc(rH.RegisterSimHandler))

	s.r.Handle("/login", http.HandlerFunc(lH.LoginHandler))

	s.r.Handle("/", mw.Auth(http.HandlerFunc(tH.TimelineHandler)))

	s.r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../web/static"))))

	// TODO + /sim/...
	// s.Get("/latest", s.LatestHandler)
	// s.Get("/msgs/{username}", s.GetUserMsgsHandler)
	// s.Post("/msgs/{username}", s.PostUserMsgsHandler)
	// s.Get("/msgs", s.MsgsHandler)
	// s.Get("/fllws/{username}", s.GetUserFollowsHandler)
	// s.Post("/fllws/{username}", s.PostUserFollowsHandler)

	return nil
}

func (s *Server) InitDB() error {
	err := s.db.AutoMigrate(
		&model.User{},
		&model.Follower{},
		&model.Message{},
	)
	return err
}
