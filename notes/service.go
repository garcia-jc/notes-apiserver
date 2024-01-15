package notes

import "context"

type NotesService interface {
	AddNote(ctx context.Context, n *Note) error
	GetNotes(ctx context.Context) ([]Note, error)
}
