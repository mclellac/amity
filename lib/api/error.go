package api

// Error
type Error struct {
	Error string `json:"error"`
}

// NewError
func NewError(msg string) *Error {
	return &Error{Error: msg}
}
