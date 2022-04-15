package main

import (
	"embed"
	"net/http"

	"github.com/JonaVDM/wake-the-wizzard/server"
	"github.com/gorilla/mux"
)

//go:embed frontend/*
var static embed.FS

func main() {
	srv := server.Server{
		Router:      mux.NewRouter(),
		StaticFiles: CustomFS{static},
	}

	srv.Init()

	http.ListenAndServe(":3000", srv.Router)
}
