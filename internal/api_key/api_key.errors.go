package apiKey

// ErrAPIKeyNotFound is returned when an API key is not found
type ErrAPIKeyNotFound struct{}

func (e ErrAPIKeyNotFound) Error() string {
	return "API key not found"
}
