package model

// An Artist struct to map the artist details
type Artist struct {
	ID             int      `json:"id"`
	Image          string   `json:"image"`
	Name           string   `json:"name"`
	Members        []string `json:"members"`
	CreationDate   int64    `json:"creationDate"`
	FirstAlbum     string   `json:"firstAlbum"`
	LocationsUrl   string   `json:"locationsUrl"`
	ConcertDatesUrl string   `json:"concertDatesUrl"`
	RelationsUrl   string `json:"datesLocation"`
	// Relations      []string  `json:"relations"`
}
