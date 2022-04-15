package server

import (
	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

func (s *Server) Init() {
	s.Router.HandleFunc("/api/list", s.routeList())
	s.Router.HandleFunc("/api/wake/{address}", s.routeWake())
}
