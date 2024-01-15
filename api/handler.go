package api

import (
	"encoding/json"
	"net/http"

	"github.com/garcia-jc/notes-apiserver/notes"
)

func (n NotesAPI) AddNote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ex := newError(http.StatusMethodNotAllowed, "invalid request",
			"only POST request are allowed")
		ex.Write(w)
		return
	}
	var input notes.Note
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		ex := newError(http.StatusBadGateway, "invalid request",
			"JSON input could not be decoded")
		ex.Write(w)
		return
	}
}

func (n NotesAPI) GetNotes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ex := newError(http.StatusMethodNotAllowed, "invalid request",
			"only GET request are allowed")
		ex.Write(w)
		return
	}
	notes, err := n.notesSrv.GetNotes(r.Context())
	if err != nil {
		ex := newError(http.StatusInternalServerError, "internal error",
			"internal failure while retrieving existing notes")
		ex.Write(w)
		return
	}
	err = json.NewEncoder(w).Encode(notes)
	if err != nil {
		ex := newError(http.StatusInternalServerError, "internal error",
			"unable to properly encode response object")
		ex.Write(w)
		return
	}
}
