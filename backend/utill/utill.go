package utill

import (
	"net/http"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

// configuration
const DATABASE = "./tmp/minitwit.db"
const PER_PAGE = 30
const DEBUG = true
const SECRET_KEY = "development key"

var store = sessions.NewCookieStore([]byte(SECRET_KEY))

func PushFlashMessage(w http.ResponseWriter, r *http.Request, message string) {
	session, err := store.Get(r, "auth")
	if err != nil {
		return
	}
	session.AddFlash(message)
	session.Save(r, w)
}

func GeneratePasswordHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}
