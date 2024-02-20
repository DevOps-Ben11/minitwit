package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const port = ":5000"

type Server struct {
	r  *mux.Router
	db *gorm.DB
}

func NewServer() Server {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
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
