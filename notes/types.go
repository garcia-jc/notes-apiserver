package notes

// general-purpose API type definitions, meant to be reused across the
// server impl and the api/client submodule

type Note struct {
	ID      string `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}
