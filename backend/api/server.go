package api

import (
	"github.com/DevOps-Ben11/minitwit/backend/model"
	"github.com/DevOps-Ben11/minitwit/backend/repository"
	"github.com/DevOps-Ben11/minitwit/backend/util"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"
	"html/template"
	"log"
	"net/http"
)

type Server struct {
	r        *mux.Router
	db       *gorm.DB
	store    *sessions.CookieStore
	userRepo repository.IUserRepository
	msgRepo  repository.IMessageRepository
}

func NewServer(db *gorm.DB, store *sessions.CookieStore, userRepo repository.IUserRepository, msgRepo repository.IMessageRepository) Server {

	s := Server{
		r:        mux.NewRouter(),
		db:       db,
		store:    store,
		userRepo: userRepo,
		msgRepo:  msgRepo,
	}

	return s
}

func (s *Server) StartServer(port string) {
	log.Println("Starting server on port", port)
	log.Fatal(http.ListenAndServe(port, s.r))
}

func (s *Server) GetStore() *sessions.CookieStore {
	return s.store
}

func (s *Server) InitRoutes() error {
	s.r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../web/static"))))

	s.r.Use(s.Auth)

	s.r.HandleFunc("/register", s.RegisterHandler)
	s.r.HandleFunc("/sim/register", s.RegisterSimHandler)

	s.r.HandleFunc("/login", s.LoginHandler)
	s.r.HandleFunc("/logout", s.LogoutHandler)

	s.r.HandleFunc("/public", s.PublicTimelineHandler)
	s.r.HandleFunc("/add_message", s.protect(s.AddMessageHandler)).Methods("POST")

	s.r.HandleFunc("/{username}/follow", s.protect(s.FollowHandler))
	s.r.HandleFunc("/{username}/unfollow", s.protect(s.UnfollowHandler))
	s.r.HandleFunc("/{username}", s.UserHandler)

	s.r.HandleFunc("/", s.protect(s.TimelineHandler))
	// TODO + /sim/...
	// s.Get("/latest", s.LatestHandler)
	// s.Get("/msgs/{username}", s.GetUserMsgsHandler)
	// s.Post("/msgs/{username}", s.PostUserMsgsHandler)
	// s.Get("/msgs", s.MsgsHandler)
	// s.Get("/fllws/{username}", s.GetUserFollowsHandler)
	// s.Post("/fllws/{username}", s.PostUserFollowsHandler)

	return nil
}

func (s *Server) InitDB() error {
	err := s.db.AutoMigrate(
		&model.User{},
		&model.Follower{},
		&model.Message{},
	)
	return err
}
func (s *Server) PushFlashMessage(w http.ResponseWriter, r *http.Request, message string) {
	session, err := s.store.Get(r, "auth")
	if err != nil {
		return
	}

	session.AddFlash(message)
	session.Save(r, w)
}

func (s *Server) GetFlashedMessages(w http.ResponseWriter, r *http.Request) []model.FlashMessage {
	session, err := s.store.Get(r, "auth")

	if err != nil {
		return []model.FlashMessage{}
	}

	flashes := session.Flashes()
	var messages []model.FlashMessage

	for _, v := range flashes {
		messages = append(messages, model.FlashMessage{Message: v.(string)})
	}

	err = session.Save(r, w)
	if err != nil {
		log.Println("Error flash:", err)
	}

	return messages
}
func (s *Server) GetCurrentUser(r *http.Request) (user *model.User, ok bool) {
	ctx := r.Context()
	user, ok = ctx.Value(UserKey).(*model.User)
	return user, ok
}

func (s *Server) GetFuncMap() template.FuncMap {
	return template.FuncMap{
		"UrlFor":             UrlFor,
		"GetFlashedMessages": s.GetFlashedMessages,
		"Gravatar":           util.Gravatar,
		"Datetimeformat":     util.Datetimeformat,
	}
}
