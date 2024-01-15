package main

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/garcia-jc/notes-apiserver/api"
	"github.com/garcia-jc/notes-apiserver/notes"
	"github.com/google/uuid"
)

func main() {
	notes := make(map[string]notes.Note)
	srv := api.NewServer(ephimerNotes(notes))
	addr := "localhost:9000"
	slog.Info("listening and serving", "address", addr)
	err := http.ListenAndServe(addr, srv)
	slog.Info("server is shutting down", "reason", err)
}

// in-memory implementation for demo purposes
type ephimerNotes map[string]notes.Note

var _ notes.NotesService = ephimerNotes{}

func (e ephimerNotes) AddNote(ctx context.Context, n *notes.Note) error {
	n.ID = uuid.NewString()
	e[n.ID] = *n
	return nil
}

func (e ephimerNotes) GetNotes(ctx context.Context) ([]notes.Note, error) {
	notes := make([]notes.Note, 0, len(e))
	for _, note := range e {
		notes = append(notes, note)
	}
	return notes, nil
}
