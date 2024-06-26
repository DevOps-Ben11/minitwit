package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/DevOps-Ben11/minitwit/backend/model"
	"github.com/DevOps-Ben11/minitwit/backend/repository"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gorm.io/gorm"
)

var (
	responseCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "minitwit_http_response_total",
		},
		[]string{"handler", "status", "method"},
	)
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
	// <------------- Simulation Routes ------------->
	simR := s.r.PathPrefix("/sim").Subrouter()
	simR.Use(s.StatusMonitoring)
	simR.Use(s.LatestMiddleware)
	simR.HandleFunc("/register", s.RegisterSimHandler).Methods("POST").Name("Sim Register")
	simR.HandleFunc("/latest", s.LatestHandler).Methods("GET").Name("Get Latest")
	simR.HandleFunc("/msgs/{username}", s.simProtect(s.MessageGetSimUserHandler)).Methods("GET").Name("Sim Get User Messages")
	simR.HandleFunc("/msgs/{username}", s.simProtect(s.MessagePostSimUserHandler)).Methods("POST").Name("Sim Post Message")
	simR.HandleFunc("/msgs", s.simProtect(s.MessagesSimHandler)).Methods("GET").Name("Sim Get Public Messages")
	simR.HandleFunc("/fllws/{username}", s.simProtect(s.FollowGetSimHandler)).Methods("GET").Name("Get Follows")
	simR.HandleFunc("/fllws/{username}", s.simProtect(s.FollowPostSimHandler)).Methods("POST").Name("Post Follows")

	// <------------- API Routes ------------->
	apiR := s.r.PathPrefix("/api").Subrouter()
	apiR.Use(s.Auth)
	apiR.HandleFunc("/register", s.RegisterHandler).Methods("POST")
	apiR.HandleFunc("/login", s.LoginHandler).Methods("POST")
	apiR.HandleFunc("/logout", s.LogoutHandler).Methods("POST")

	apiR.HandleFunc("/public", s.PublicTimelineHandler).Methods("GET")
	apiR.HandleFunc("/add_message", s.protect(s.AddMessageHandler)).Methods("POST")
	apiR.HandleFunc("/timeline", s.protect(s.TimelineHandler)).Methods("GET")

	apiR.HandleFunc("/{username}/follow", s.protect(s.FollowHandler)).Methods("POST")
	apiR.HandleFunc("/{username}/unfollow", s.protect(s.UnfollowHandler)).Methods("POST")
	apiR.HandleFunc("/timeline/{username}", s.UserHandler).Methods("GET")

	// <------------- Static Routes ------------->
	s.r.Handle("/metrics", promhttp.Handler())

	// Serve static files
	s.r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("../client/dist/assets/"))))

	s.r.HandleFunc("/icon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../client/dist/icon.ico")
	})

	s.r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../client/dist/index.html")
	})

	return nil
}

func (s *Server) InitDB() error {
	log.Println("Initializing DB")
	err := s.db.AutoMigrate(
		&model.User{},
		&model.Follower{},
		&model.Message{},
		&model.KeyVal{},
	)
	return err
}

func (s *Server) GetCurrentUser(r *http.Request) (user *model.User, ok bool) {
	ctx := r.Context()
	user, ok = ctx.Value(UserKey).(*model.User)
	return user, ok
}

func (s *Server) GetKeyVal(key string) (model.KeyVal, error) {
	ret := model.KeyVal{Key: key}
	err := s.db.First(&ret).Error
	return ret, err
}

func (s *Server) SetKeyVal(key string, value string) error {
	err := s.db.Save(&model.KeyVal{Key: key, Value: value}).Error
	return err
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func (s *Server) StatusMonitoring(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lrw := &responseWriter{ResponseWriter: w, status: 200}
		next.ServeHTTP(lrw, r)
		var handlerLabel string
		route := mux.CurrentRoute(r)
		if route != nil {
			name := route.GetName()
			if name != "" {
				handlerLabel = name
			}
		}

		status := strconv.Itoa(lrw.status)

		responseCounter.WithLabelValues(handlerLabel, status, r.Method).Inc()
	})
}
