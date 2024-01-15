package notes

import "context"

// using the model defined in types.go, define the high-level service interface
// that may be implemented in order to fulfil the purpose of this API server

type NotesService interface {
	AddNote(ctx context.Context, n *Note) error
	GetNotes(ctx context.Context) ([]Note, error)
}
