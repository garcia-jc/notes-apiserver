package notes

type Note struct {
	ID      string `json:"id"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}
