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

	cookieHMAC, okHmac := os.LookupEnv("SECRET_COOKIE_HMAC")
	cookieAES, okAes := os.LookupEnv("SECRET_COOKIE_AES")

	if !okHmac || !okAes {
		panic("SECRET_COOKIE_HMAC or/and SECRET_COOKIE_AES not found in env")
	}

	log.Println("SECRET_COOKIE_HMAC string: ", cookieHMAC)
	log.Println("SECRET_COOKIE_AES string: ", cookieAES)

	var store = sessions.NewCookieStore(
		[]byte(cookieHMAC), []byte(cookieAES),

		// To support legacy cookies
		[]byte("development key"), nil,
	)
	store.Options.HttpOnly = true

	// Must be turned on when HTTPS is made available on production
	// store.Options.Secure = true

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
