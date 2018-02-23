package chat

// LocationAlreadyExistsError ...
type LocationAlreadyExistsError struct{}

func (e *LocationAlreadyExistsError) Error() string {
	return "Location already exists!"
}
