package server

type TestMsg struct {
	Msg string `json:"msg"`
}

// func (s *Server) TestHandler(vars map[string]string, r *http.Request) (status int, value any) {
// 	name := vars["name"]
// 	msg := fmt.Sprintf("Hello %s!", name)
// 	s.db.Create(&model.Example{
// 		Msg: msg,
// 	})

// 	return OkResponse(TestMsg{Msg: msg})
// }

// func (s *Server) TestPostHandler(vars map[string]string, r *http.Request) (status int, value any) {
// 	var data TestMsg
// 	DecodeBody(r.Body, &data)
// 	log.Println("Got message!")
// 	log.Println(data.Msg)
// 	return OkResponse(nil)
// }

// func (s *Server) LatestHandler(vars map[string]string, r *http.Request) (status int, value any) {

// 	return 404, nil
// }

// func (s *Server) RegisterSimHandler(vars map[string]string, r *http.Request) (status int, value any) {

// 	return 404, nil
// }

// func (s *Server) MsgsHandler(vars map[string]string, r *http.Request) (status int, value any) {

// 	return 404, nil
// }

// func (s *Server) GetUserMsgsHandler(vars map[string]string, r *http.Request) (status int, value any) {
// 	username := vars["username"]

// 	return 404, username
// }

// func (s *Server) PostUserMsgsHandler(vars map[string]string, r *http.Request) (status int, value any) {
// 	username := vars["username"]

// 	return 404, username
// }

// func (s *Server) GetUserFollowsHandler(vars map[string]string, r *http.Request) (status int, value any) {
// 	username := vars["username"]

// 	return 404, username
// }

// func (s *Server) PostUserFollowsHandler(vars map[string]string, r *http.Request) (status int, value any) {
// 	username := vars["username"]

// 	return 404, username
// }
