package server

import "github.com/DevOps-Ben11/minitwit/backend/model"

func (s *Server) InitDB() error {
	err := s.db.AutoMigrate(
		&model.User{},
		&model.Follower{},
		&model.Message{},
	)
	return err
}
