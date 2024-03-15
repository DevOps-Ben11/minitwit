package main

import (
	"fmt"
	"log"

	"github.com/DevOps-Ben11/minitwit/backend/api"
	"github.com/DevOps-Ben11/minitwit/backend/db"
	"github.com/DevOps-Ben11/minitwit/backend/repository"
	"github.com/gorilla/sessions"
)

const port = ":5000"
const DEBUG = true
const SECRET_KEY = "development key"

func main() {
	db, err := db.GetDB()

	if err != nil {
		log.Fatalln("Could not open Database", err)
	}

	var store = sessions.NewCookieStore([]byte(SECRET_KEY))

	userRepo := repository.CreateUserRepository(db)
	msgRepo := repository.CreateMessageRepository(db)

	s := api.NewServer(db, store, userRepo, msgRepo)

	err = s.InitRoutes()
	if err != nil {
		log.Fatalln(fmt.Sprintf("Error when adding routes: %v", err))
	}

	err = s.InitDB()
	if err != nil {
		log.Fatalln(fmt.Sprintf("Error when initializing database: %v", err))
	}

	s.StartServer(port)
}
