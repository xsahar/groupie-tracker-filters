package backend

import (
	// Import the necessary packages
	"fmt"
	"net/http"
	"text/template"
)

// Initialize a template variable to be used for rendering HTML templates
var tmp *template.Template

// The init function is called when the package is imported
func init() {
	// Parse all the HTML templates in the "frontend/html/" directory
	// and store them in the tmp variable
	tmp = template.Must(template.ParseGlob("frontend/html/*.html"))
}

// MainPage is an HTTP handler function that handles the "/" endpoint
func MainPage(w http.ResponseWriter, r *http.Request) {
	// Check if the requested URL path is "/"
	if r.URL.Path != "/" {
		// If the path is not "/", return a 404 error page
		err := "404 Page not found"
		fmt.Println(err)
		ErrorPage(w, err, http.StatusNotFound)
		return
	}

	// Check if the request method is GET
	if r.Method != http.MethodGet {
		// If the request method is not GET, return a 405 error page
		err := "405 Method is not allowed"
		fmt.Println(err)
		ErrorPage(w, err, http.StatusMethodNotAllowed)
		return
	}

	// Retrieve all the artists from the data source
	artists, err := GetAllArtists()
	if err != nil {
		// If there's an error retrieving the artists, return a 500 error page
		err := "500 Internal Server Error"
		fmt.Println(err)
		ErrorPage(w, err, http.StatusInternalServerError)
		return
	}

	// Retrieve all the locations from the data source
	location, err := GetAllLocations()
	if err != nil {
		// If there's an error retrieving the locations, return a 500 error page
		err := "500 Internal Server Error"
		fmt.Println(err)
		ErrorPage(w, err, http.StatusInternalServerError)
		return
	}

	// Create an Everything object that contains the artists and locations
	allInfo := Everything{artists, location}

	// Render the index.html template with the allInfo object
	err = tmp.ExecuteTemplate(w, "index.html", allInfo)
	if err != nil {
		// If there's an error rendering the template, return a 500 error page
		err := "500 Internal Server Error"
		fmt.Println(err)
		ErrorPage(w, err, http.StatusInternalServerError)
		return
	}
}

// ErrorPage is a helper function that renders an error page
func ErrorPage(w http.ResponseWriter, errors string, code int) {
	// Set the HTTP status code for the response
	w.WriteHeader(code)

	// Render the error.html template with the error message
	err := tmp.ExecuteTemplate(w, "error.html", errors)
	if err != nil {
		// If there's an error rendering the template, return a 500 error
		err := "500 Internal Server Error"
		fmt.Println(err)
		http.Error(w, err, 500)
		return
	}
}