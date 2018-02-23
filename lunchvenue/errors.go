package lunchvenue

// LocationAlreadyExistsError ...
type LocationAlreadyExistsError struct{}

func (e *LocationAlreadyExistsError) Error() string {
	return "Location already exists!"
}

// LocationNotFoundError ...
type LocationNotFoundError struct{}

func (e *LocationNotFoundError) Error() string {
	return "Location does not exist!"
}
