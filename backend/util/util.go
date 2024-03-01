package util

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/eefret/gravatar"
	"golang.org/x/crypto/bcrypt"
)

const PER_PAGE = 30

func GeneratePasswordHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 3)
	return string(bytes)
}

func CheckPassword(password string, password_hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password_hash), []byte(password))
	return err == nil
}

func Gravatar(size int, email string) string {
	g, err := gravatar.New()

	if err != nil {
		log.Print("Error in fetching gravatar", err)
		return fmt.Sprintf("https://www.gravatar.com/avatar?s=%d", size)
	}

	g.SetSize(uint(size))

	return g.URLParse(strings.TrimSpace(strings.ToLower(email))) + "&d=identicon"
}

func Datetimeformat(date int64) string {
	t := time.Unix(date, 0)

	// Format time as "Month Date HH:MM"
	formattedTime := t.Format("January 02 2006 15:04")

	return formattedTime
}
