package main

import (
	"log"
	"os"

	"github.com/DevOps-Ben11/minitwit/backend/api"
	"github.com/DevOps-Ben11/minitwit/backend/db"
	"github.com/DevOps-Ben11/minitwit/backend/repository"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

const port = ":5000"
const DEBUG = true

func main() {
	dotErr := godotenv.Load()
	if dotErr != nil {
		log.Println("No .env found, continuing without")
	}
	db, err := db.GetDB()

	if err != nil {
		log.Fatalln("Could not open Database", err)
	}

	secretKey, ok := os.LookupEnv("SECRET_KEY")

	if !ok {
		panic("SECRET_KEY not found in env")
	}

	var store = sessions.NewCookieStore([]byte(secretKey))

	userRepo := repository.CreateUserRepository(db)
	msgRepo := repository.CreateMessageRepository(db)

	s := api.NewServer(db, store, userRepo, msgRepo)

	err = s.InitRoutes()
	if err != nil {
		log.Fatalf("Error when adding routes: %v\n", err)
	}

	err = s.InitDB()
	if err != nil {
		log.Fatalf("Error when initializing database: %v\n", err)
	}

	s.StartServer(port)
}
