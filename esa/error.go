package esa

import "fmt"

// Error is struct for esa error
type Error struct {
	StatusCode int
	Err        string `json:"error"`
	Message    string `json:"message"`
}

// Error output error message
func (e Error) Error() string {
	return fmt.Sprintf(
		"failed to esa requrest: statusCode='%d', error='%s', message='%s'",
		e.StatusCode,
		e.Err,
		e.Message,
	)
}
