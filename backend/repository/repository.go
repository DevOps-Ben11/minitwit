package repository

import (
	"errors"

	"github.com/DevOps-Ben11/minitwit/backend/server"
	"github.com/DevOps-Ben11/minitwit/backend/utill"
)

func registerDB(s server.Server, username string, email string, password string) error {
	if s.GetUserId(username) != 0 {
		errors.New("The username is already taken")

	}
	err := s.db.Exec("insert into user (username, email, pw_hash) values (?, ?, ?)", username, email, utill.GeneratePasswordHash(password)).Error

	if err != nil {
		panic("AAAAAH!!!!")
	}
	return nil
}
