package server

import (
	"net/http"

	"github.com/JonaVDM/wake-the-wizzard/wol"
	"github.com/gorilla/mux"
)

func (s *Server) routeWake() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		address := vars["address"]

		wol.SendWol(address)
		w.Write([]byte("ok"))
	}
}
