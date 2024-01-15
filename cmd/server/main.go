package main

import (
	"log/slog"
	"net/http"

	"github.com/garcia-jc/notes-apiserver/api"
	"github.com/garcia-jc/notes-apiserver/notes"
)

func main() {
	var notes notes.NotesService
	srv := api.NewServer(notes)
	addr := "localhost:9000"
	slog.Info("listening and serving", "address", addr)
	err := http.ListenAndServe(addr, srv)
	slog.Info("server is shutting down", "reason", err)
}
