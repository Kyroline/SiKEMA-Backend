package updateExcuse

type InputUpdateExcuse struct {
	ID         uint
	Excuse     string `json:"excuse"`
	Attachment string `json:"attachment"`
}
