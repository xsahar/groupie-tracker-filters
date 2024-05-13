package handler

import (
	"encoding/json"
	"groupie-tracker/model"
	"html/template"
	"log"
	"net/http"
	"strconv"
	// "fmt"
	"strings"
)

// Global Variable to hold parsed html template
var templates *template.Template
var artists []model.Artist // Define artists as a global variable

// Data struct now includes dates
type Data struct {
	Artist    model.Artist
	Relations model.DatesLocations
	Locations model.Location
	Dates     model.ConcertDate
}

func init() {
	// Parse the HTML templates during initialization
	templates = template.Must(template.ParseGlob("templates/*.html"))

	// Fetch data during initialization
	var err error
	artists, err = fetchData()
	if err != nil {
		log.Fatal("Error fetching artists data:", err)
	}
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		if r.Method == http.MethodGet {
			// Get the search query from the URL parameters
			query := r.URL.Query().Get("search")

			// Filter the artists based on the search query
			filteredArtists := filterArtists(artists, query)

			// Create TemplateData struct with filtered artists and query
			data := TemplateData{
				Artists: filteredArtists,
				Query:   query,
			}

			if err := renderTemplate(w, r, "index.html", data); err != nil {
				log.Println("Error rendering template:", err)
				errorHandler(w, r, http.StatusInternalServerError)
				return
			}
		} else {
			errorHandler(w, r, http.StatusMethodNotAllowed)
			return
		}
	} else {
		errorHandler(w, r, http.StatusNotFound)
	}
}

// Function to filter artists based on the search query
func filterArtists(artists []model.Artist, query string) []model.Artist {
	if query == "" {
		return artists
	}

	firstLetter := strings.ToLower(query[:1]) // Get the first letter of the query

	var filteredArtists []model.Artist
	for _, artist := range artists {
		// Perform case-insensitive search on the first letter of the artist name
		if strings.HasPrefix(strings.ToLower(artist.Name), firstLetter) {
			filteredArtists = append(filteredArtists, artist)
		}
	}

	return filteredArtists
}

type TemplateData struct {
	Artists []model.Artist
	Query   string
}

func renderTemplate(w http.ResponseWriter, r *http.Request, tmplName string, data interface{}) error {
	tmpl, err := templates.Clone()
	if err != nil {
		return err
	}
	if err := tmpl.ExecuteTemplate(w, tmplName, data); err != nil {
		return err
	}
	return nil
}

func fetchData() ([]model.Artist, error) {
	// Fetch data from artist API
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Parse the JSON response into a slice of Artist structs
	var artists []model.Artist
	if err := json.NewDecoder(response.Body).Decode(&artists); err != nil {
		return nil, err
	}
	return artists, nil
}

func fetchRelations(id string) (model.DatesLocations, error) {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + id)
	if err != nil {
		return model.DatesLocations{}, err
	}
	defer res.Body.Close()

	var relations model.DatesLocations
	if err := json.NewDecoder(res.Body).Decode(&relations); err != nil {
		return model.DatesLocations{}, err
	}

	return relations, nil
}

func fetchLocations(id string) (model.Location, error) {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/locations/" + id)
	if err != nil {
		return model.Location{}, err
	}
	defer res.Body.Close()

	var locations model.Location
	if err := json.NewDecoder(res.Body).Decode(&locations); err != nil {
		return model.Location{}, err
	}

	return locations, nil
}

func fetchDates(id string) (model.ConcertDate, error) {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/dates/" + id)
	if err != nil {
		return model.ConcertDate{}, err
	}
	defer res.Body.Close()

	var dates model.ConcertDate
	if err := json.NewDecoder(res.Body).Decode(&dates); err != nil {
		return model.ConcertDate{}, err
	}

	return dates, nil
}

func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/details/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	// Find the artist by ID
	var chosenArtist model.Artist
	found := false
	for _, artist := range artists {
		if artist.ID == id {
			chosenArtist = artist
			found = true
			break
		}
	}

	if !found {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	// Fetch relations
	relations, err := fetchRelations(idStr)
	if err != nil {
		log.Println("Error fetching relations:", err)
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Fetch locations
	locations, err := fetchLocations(idStr)
	if err != nil {
		log.Println("Error fetching locations:", err)
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Fetch dates
	dates, err := fetchDates(idStr)
	if err != nil {
		log.Println("Error fetching dates:", err)
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}

	data := Data{
		Artist:    chosenArtist,
		Relations: relations,
		Locations: locations,
		Dates:     dates,
	}

	// Attempt to render the template
	err = renderTemplate(w, r, "details.html", data)
	if err != nil {
		log.Println("Error rendering template:", err)
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}
}

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	errorHandler(w, r, http.StatusInternalServerError)
}
