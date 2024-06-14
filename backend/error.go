package backend

import (
	"fmt"
	"net/http"
	"text/template"
)

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseGlob("frontend/html/*.html"))
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		err := "404 Page not found"
		fmt.Println(err)
		ErrorPage(w, err, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		err := "405 Method is not allowed"
		fmt.Println(err)
		ErrorPage(w, err, http.StatusMethodNotAllowed)
		return
	}

	artists, err := GetAllArtists()
	if err != nil {
		err := "500 Internal Server Error"
		fmt.Println(err)
		ErrorPage(w, err, http.StatusInternalServerError)
		return
	}
	location, err := GetAllLocations()
	if err != nil {
		err := "500 Internal Server Error"
		fmt.Println(err)
		ErrorPage(w, err, http.StatusInternalServerError)
		return
	}
	allInfo := Everything{artists, location}

	err = tmp.ExecuteTemplate(w, "index.html", allInfo)
	if err != nil {
		err := "500 Internal Server Error"
		fmt.Println(err)
		ErrorPage(w, err, http.StatusInternalServerError)
		return
	}
}

func ErrorPage(w http.ResponseWriter, errors string, code int) {
	w.WriteHeader(code)
	err:=tmp.ExecuteTemplate(w, "error.html", errors)
	if err != nil {
		err := "500 Internal Server Error"
		fmt.Println(err)
		http.Error(w, err, 500)
		return
	}
}
