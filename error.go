package kraken

import "fmt"

// Error represents a Kraken API error.
type Error struct {
	errors []string
}

// Error builds a Kraken API error.
func (e *Error) Error() string {
	// TODO: improve error message
	return fmt.Sprintf("%v", e.errors)
}
