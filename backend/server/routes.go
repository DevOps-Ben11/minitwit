package server

// func (s *main.Server) InitRoutes() error {

// 	// s.Get("/latest", s.LatestHandler)
// 	s.r.Handle("/register", s.auth(http.HandlerFunc(s.RegisterHandler))).Methods("GET")
// 	// s.Post("/sim/register", s.RegisterSimHandler)

// 	// s.Get("/msgs/{username}", s.GetUserMsgsHandler)
// 	// s.Post("/msgs/{username}", s.PostUserMsgsHandler)
// 	// s.Get("/msgs", s.MsgsHandler)
// 	// s.Get("/fllws/{username}", s.GetUserFollowsHandler)
// 	// s.Post("/fllws/{username}", s.PostUserFollowsHandler)

// 	return nil
// }

// type Handler func(vars map[string]string, r *http.Request) (status int, value any)

// func (s *Server) Handle(route string, handler Handler, method string) {
// 	f := func(w http.ResponseWriter, r *http.Request) {
// 		vars := mux.Vars(r)

// 		status, value := handler(vars, r)

// 		returnValue, err := json.Marshal(value)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(status)
// 		w.Write(returnValue)
// 	}

// 	s.r.HandleFunc(route, f).Methods(method)
// }

// // authentication middleware

// func (s *Server) Get(route string, handler Handler) {
// 	s.Handle(route, handler, "GET")
// }

// func (s *Server) Post(route string, handler Handler) {
// 	s.Handle(route, handler, "POST")
// }

// func DecodeBody(body io.ReadCloser, v any) error {
// 	return json.NewDecoder(body).Decode(v)
// }

// func OkResponse(value any) (int, any) {
// 	return http.StatusOK, value
// }
