package main

import (
	"net/http"

	"github.com/JonaVDM/wake-the-wizzard/server"
	"github.com/gorilla/mux"
)

func main() {
	srv := server.Server{
		Router: mux.NewRouter(),
	}

	srv.Init()

	http.ListenAndServe(":3000", srv.Router)
}
