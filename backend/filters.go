package backend

import "net/http"

func FilterhHandler(w http.ResponseWriter, r *http.Request) {
	
	if r.URL.Path != "/filter/" {
		err := "404 Page not found"
		ErrorPage(w, err, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		err := "405 Method is not allowed"
		ErrorPage(w, err, http.StatusMethodNotAllowed)
		return
	}

	artists, err := GetAllArtists()
	if err != nil {
		err := "500 Internal Server Error"
		ErrorPage(w, err, http.StatusInternalServerError)
		return
	}
	location, err := GetAllLocations()
	if err != nil {
		err := "500 Internal Server Error"
		ErrorPage(w, err, http.StatusInternalServerError)
		return
	}

	range1 := r.FormValue("range1")
	range2 := r.FormValue("range2")
	searchedArtists := Everything{artists, location}
	locations := r.FormValue("Locations")

	if locations != "" {
		searchedArtists, err = LocationSearch(searchedArtists, locations)
		if err != nil {
			err := "500 Internal Server Error"
			ErrorPage(w, err, http.StatusInternalServerError)
			return
		}
	}

	if range1 != "" && range2 != "" {
		searchedArtists, err = CreationDate(searchedArtists, range1, range2)
		if err != nil {
			err := "500 Internal Server Error"
			ErrorPage(w, err, http.StatusInternalServerError)
			return
		}
	}

	firstalbum1 := r.FormValue("range3")
	firstalbum2 := r.FormValue("range4")
	if firstalbum1 != "" && firstalbum2 != "" {
		searchedArtists, err = FirstAlbumDates(searchedArtists, firstalbum1, firstalbum2)
		if err != nil {
			err := "500 Internal Server Error"
			ErrorPage(w, err, http.StatusInternalServerError)
			return
		}
	}
	member := r.Form["interest"]
	if len(member) != 0 {
		searchedArtists, err = MembersNumber(searchedArtists, member)
		if err != nil {
			err := "500 Internal Server Error"
			ErrorPage(w, err, http.StatusInternalServerError)
			return
		}
	}

	err = tmp.ExecuteTemplate(w, "search.html", searchedArtists)
	if err != nil {
		err := "500 Internal Server Error"
		ErrorPage(w, err, http.StatusInternalServerError)
		return
	}
}