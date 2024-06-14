package main

import (
    "fmt"
    "groupie-tracker-filters/backend"
    "log"
    "net/http"
)

func main() {
    fs := http.FileServer(http.Dir("./frontend/css_style"))
    http.Handle("/css_style/", http.StripPrefix("/css_style", fs))
    http.HandleFunc("/", backend.MainPage)
    http.HandleFunc("/artists/", backend.InfoAboutArtist)
    http.HandleFunc("/filter/", backend.FilterhHandler)
    fmt.Println("http://localhost:8000/")
    err := http.ListenAndServe(":8000", nil)
    if err != nil {
        log.Fatal(err)
    }
}