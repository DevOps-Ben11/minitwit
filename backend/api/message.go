package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/DevOps-Ben11/minitwit/backend/model"
	"github.com/gorilla/mux"
)

type AddMessage struct {
	Message string `json:"message"`
}

func (s *Server) AddMessageHandler(user *model.User, w http.ResponseWriter, r *http.Request) {
	// Register a new message from a user
	var body AddMessage
	json.NewDecoder(r.Body).Decode(&body)

	message := body.Message

	// early return if empty
	if message == "" {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	err := s.msgRepo.AddMessage(user, message)

	if err != nil {
		log.Println("Error creating message: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) MessagesSimHandler(w http.ResponseWriter, r *http.Request) {
	no := r.URL.Query().Get("no")
	numMsgs := 100
	if no != "" {
		num, err := strconv.Atoi(no)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		numMsgs = num
	}

	messages, err := s.msgRepo.GetPublicMessages(numMsgs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	filteredMsgs := []map[string]any{}
	for _, msg := range messages {
		tmpMessage := map[string]any{}
		tmpMessage["content"] = msg.Text
		tmpMessage["pub_date"] = msg.Pub_date
		tmpMessage["user"] = msg.Username
		filteredMsgs = append(filteredMsgs, tmpMessage)
	}

	RetJson(w, filteredMsgs)
}

func (s *Server) MessageGetSimUserHandler(w http.ResponseWriter, r *http.Request) {
	no := r.URL.Query().Get("no")
	numMsgs := 100

	if no != "" {
		num, err := strconv.Atoi(no)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		numMsgs = num
	}

	vars := mux.Vars(r)
	username := vars["username"]
	user, ok := s.userRepo.GetUser(username)

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	messages, err := s.msgRepo.GetUserMessages(user.User_id, numMsgs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	filteredMsgs := []map[string]any{}
	for _, msg := range messages {
		tmpMessage := map[string]any{}
		tmpMessage["content"] = msg.Text
		tmpMessage["pub_date"] = msg.Pub_date
		tmpMessage["user"] = msg.Username
		filteredMsgs = append(filteredMsgs, tmpMessage)
	}

	RetJson(w, filteredMsgs)

}

func (s *Server) MessagePostSimUserHandler(w http.ResponseWriter, r *http.Request) {
	type MessagePost struct {
		Content string `json:"content"`
	}

	var body MessagePost
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println("MessagePostSimUserHandler: Error decoding body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	username := vars["username"]
	user, ok := s.userRepo.GetUser(username)

	if !ok {
		log.Println("MessagePostSimUserHandler: Error getting user:", username, body.Content)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = s.msgRepo.AddMessage(user, body.Content)
	if err != nil {
		log.Println("MessagePostSimUserHandler: Error adding message:", err, body.Content)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
