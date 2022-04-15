package server

import (
	"io/fs"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Router      *mux.Router
	StaticFiles fs.FS
}

func (s *Server) Init() {
	s.Router.HandleFunc("/api/list", s.routeList())
	s.Router.HandleFunc("/api/wake/{address}", s.routeWake())
	s.Router.PathPrefix("/").Handler(http.FileServer(http.FS(s.StaticFiles)))
}
