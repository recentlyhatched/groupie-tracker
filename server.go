package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func runServer() {
	http.HandleFunc("/browse", browseHandler)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func browseHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		// TODO: add 404 page template
	}

	template, err := template.ParseFiles("templates/browse.html")
	if err != nil {
		http.Error(writer, "Error parsing browse.html template", http.StatusInternalServerError)
		return
	}

	// an array of artist data for each artist struct
	artistsData, err := Artists()
	if err != nil {
		http.Error(writer, "Trouble loading artists", http.StatusInternalServerError)
	}

	template.Execute(writer, artistsData)

}
