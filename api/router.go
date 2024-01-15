package api

import (
	"net/http"

	"github.com/garcia-jc/notes-apiserver/notes"
)

type NotesAPI struct {
	notesSrv notes.NotesService
}

func NewServer(notes notes.NotesService) http.Handler {
	api := NotesAPI{notesSrv: notes}
	mux := http.NewServeMux()
	mux.HandleFunc("/note", api.AddNote)
	mux.HandleFunc("/notes", api.GetNotes)
	return mux
}
