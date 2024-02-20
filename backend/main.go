package main

import (
	"fmt"
	"log"

	"github.com/DevOps-Ben11/minitwit/backend/server"
)

func main() {

	s := server.NewServer()
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
