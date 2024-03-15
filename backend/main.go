package main

import (
	"fmt"
	"log"

	"github.com/DevOps-Ben11/minitwit/backend/api"
	"github.com/DevOps-Ben11/minitwit/backend/repository"
	"github.com/gorilla/sessions"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const port = ":5000"
const DEBUG = true
const SECRET_KEY = "development key"

func main() {
	db, err := gorm.Open(sqlite.Open("../tmp/minitwit.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

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
