package backend

import "net/http"

// FilterhHandler is an HTTP handler function that handles the "/filter/" endpoint
func FilterhHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the requested URL path is "/filter/"
	if r.URL.Path != "/filter/" {
		err := "404 Page not found"
		ErrorPage(w, err, http.StatusNotFound)
		return
	}

	// Check if the request method is GET
	if r.Method != http.MethodGet {
		err := "405 Method is not allowed"
		ErrorPage(w, err, http.StatusMethodNotAllowed)
		return
	}

	// Retrieve all the artists from the data source
	artists, err := GetAllArtists()
	if err != nil {
		err := "500 Internal Server Error"
		ErrorPage(w, err, http.StatusInternalServerError)
		return
	}

	// Retrieve all the locations from the data source
	location, err := GetAllLocations()
	if err != nil {
		err := "500 Internal Server Error"
		ErrorPage(w, err, http.StatusInternalServerError)
		return
	}

	// Get the query parameters from the request
	range1 := r.FormValue("range1")
	range2 := r.FormValue("range2")
	locations := r.FormValue("Locations")
	firstalbum1 := r.FormValue("range3")
	firstalbum2 := r.FormValue("range4")
	member := r.Form["interest"]

	// Create a searchedArtists object that contains the artists and locations
	searchedArtists := Everything{artists, location}

	// Filter the artists by location
	if locations != "" {
		searchedArtists, err = LocationSearch(searchedArtists, locations)
		if err != nil {
			err := "500 Internal Server Error"
			ErrorPage(w, err, http.StatusInternalServerError)
			return
		}
	}

	// Filter the artists by creation date
	if range1 != "" && range2 != "" {
		searchedArtists, err = CreationDate(searchedArtists, range1, range2)
		if err != nil {
			err := "500 Internal Server Error"
			ErrorPage(w, err, http.StatusInternalServerError)
			return
		}
	}

	// Filter the artists by first album date
	if firstalbum1 != "" && firstalbum2 != "" {
		searchedArtists, err = FirstAlbumDates(searchedArtists, firstalbum1, firstalbum2)
		if err != nil {
			err := "500 Internal Server Error"
			ErrorPage(w, err, http.StatusInternalServerError)
			return
		}
	}

	// Filter the artists by number of members
	if len(member) != 0 {
		searchedArtists, err = MembersNumber(searchedArtists, member)
		if err != nil {
			err := "500 Internal Server Error"
			ErrorPage(w, err, http.StatusInternalServerError)
			return
		}
	}

	// Render the search.html template with the filtered searchedArtists object
	err = tmp.ExecuteTemplate(w, "search.html", searchedArtists)
	if err != nil {
		err := "500 Internal Server Error"
		ErrorPage(w, err, http.StatusInternalServerError)
		return
	}
}