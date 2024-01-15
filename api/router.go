package api

import (
	"net/http"

	"github.com/garcia-jc/notes-apiserver/notes"
)

type NotesAPI struct {
	srv notes.NotesService
}

func NewServer(notesSrv notes.NotesService) http.Handler {
	api := NotesAPI{srv: notesSrv}
	mux := http.NewServeMux()
	mux.HandleFunc("/note", api.AddNote)
	mux.HandleFunc("/notes", api.GetNotes)
	return mux
}
