package main

import (
	"fmt"
	"groupie-tracker/handler"
	"log"
	"net/http"
)

func main() {
	// Serve the CSS file
	fs := http.FileServer(http.Dir("style"))
	http.Handle("/style/", http.StripPrefix("/style/", fs))

	http.HandleFunc("/", handler.FormHandler)
	http.HandleFunc("/details/", handler.DetailsHandler)
	fmt.Println("Server started at http://localhost:3000/")

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}