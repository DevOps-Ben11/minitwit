package util

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/DevOps-Ben11/minitwit/api/model"
	"github.com/eefret/gravatar"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

// configuration
const DATABASE = "./tmp/minitwit.db"
const PER_PAGE = 30
const DEBUG = true
const SECRET_KEY = "development key"

var store = sessions.NewCookieStore([]byte(SECRET_KEY))

func GetStore() *sessions.CookieStore {
	return store
}

func PushFlashMessage(w http.ResponseWriter, r *http.Request, message string) {
	session, err := store.Get(r, "auth")
	if err != nil {
		return
	}

	session.AddFlash(message)
	session.Save(r, w)
}

func GetFlashedMessages(w http.ResponseWriter, r *http.Request) []model.FlashMessage {
	session, err := store.Get(r, "auth")

	if err != nil {
		return []model.FlashMessage{}
	}

	flashes := session.Flashes()
	messages := []model.FlashMessage{}

	for _, v := range flashes {
		messages = append(messages, model.FlashMessage{Message: v.(string)})
	}

	err = session.Save(r, w)
	if err != nil {
		log.Println("Error flash:", err)
	}

	return messages
}

func GeneratePasswordHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func CheckPassword(password string, password_hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password_hash), []byte(password))
	return err == nil
}

func GetCurrentUser(r *http.Request) (id *uint, ok bool) {
	session, err := store.Get(r, "auth")
	if err != nil {
		return nil, false
	}

	u, ok := session.Values["user"]
	if !ok {
		return nil, false
	}

	var user_id uint

	// Attempt to assert the type of u to uint
	user_id, ok = u.(uint)

	if !ok {
		// If it's not of type uint, return nil and false
		return nil, false
	}

	return &user_id, true
}

func UrlFor(path string, arg string) string {
	switch path {
	case "add_message":
		return "/add_message"
	case "user_timeline":
		return fmt.Sprintf("/%s", arg)
	case "unfollow_user":
		return fmt.Sprintf("/%s/unfollow", arg)
	case "follow_user":
		return fmt.Sprintf("/%s/follow", arg)
	case "timeline":
		return "/"
	case "public_timeline":
		return "/public"
	case "register":
		return "/register"
	case "login":
		return "/login"
	case "logout":
		return "/logout"
	default:
		return "/"
	}
}

func Gravatar(size int, email string) string {
	g, err := gravatar.New()

	if err != nil {
		log.Print("Error in fetching gravatar", err)
		return fmt.Sprintf("https://www.gravatar.com/avatar?s=%d", size)
	}

	g.SetSize(uint(size))
	log.Println(email)

	return g.URLParse(strings.TrimSpace(strings.ToLower(email))) + "&d=identicon"
}

func Datetimeformat(date int64) string {
	t := time.Unix(date, 0)

	// Format time as "Month Date HH:MM"
	formattedTime := t.Format("January 02 2006 15:04")

	return formattedTime
}

func GetFuncMap() template.FuncMap {
	return template.FuncMap{
		"UrlFor":             UrlFor,
		"GetFlashedMessages": GetFlashedMessages,
		"Gravatar":           Gravatar,
		"Datetimeformat":     Datetimeformat,
	}
}
