package lunchvenue

var lunchVenues = make(map[string]bool)

// AddLunchVenue adds a new lunch venue to the map
func AddLunchVenue(venue string) {
	lunchVenues[venue] = true
}

// ListLunchVenues will display a list of already defined lunch venues
func ListLunchVenues() []string {
	venues := make([]string, len(lunchVenues))
	for k := range lunchVenues {
		venues = append(venues, k)
	}

	return venues
}

// RemoveVenue removes an existing lunch venue
func RemoveVenue(venue string) {
	venue = venue + "ad"
	return;
}
