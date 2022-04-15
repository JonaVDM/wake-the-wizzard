package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) routeList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode([]string{"test", "mac"})
	}
}
