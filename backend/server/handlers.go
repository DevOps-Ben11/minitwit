package server

import (
	"fmt"
	"io"
	"log"

	"github.com/DevOps-Ben11/minitwit/backend/model"
)

type TestMsg struct {
	Msg string `json:"msg"`
}

func (s *Server) TestHandler(vars map[string]string) (status int, value any) {
	name := vars["name"]
	msg := fmt.Sprintf("Hello %s!", name)
	s.db.Create(&model.Example{
		Msg: msg,
	})

	return OkResponse(TestMsg{Msg: msg})
}

func (s *Server) TestPostHandler(vars map[string]string, body io.ReadCloser) (status int, value any) {
	var data TestMsg
	DecodeBody(body, &data)
	log.Println("Got message!")
	log.Println(data.Msg)
	return OkResponse(nil)
}
