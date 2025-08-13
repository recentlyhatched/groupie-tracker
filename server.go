package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func runServer() {
	http.HandleFunc("/browse", browseHandler)
	http.HandleFunc("/artists/", artistHandler)
	http.HandleFunc("/", homeHandler)

	// link css
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		http.Error(writer, "Error parsing form data", http.StatusBadRequest)
		return
	}
	searchYearString := request.FormValue("year-input")

	searchYear, err := strconv.Atoi(searchYearString)

	artistsByYear, err := GetGroupByCreationYear(searchYear)
	if err != nil {
		http.Error(writer, "Error getting artists by year", http.StatusInternalServerError)
		return
	}

	template, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(writer, "Error parsing index.html template", http.StatusInternalServerError)
		return
	}

	template.Execute(writer, artistsByYear)
}

func browseHandler(writer http.ResponseWriter, request *http.Request) {
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

	if request.URL.Path != "/browse" {
		// 404 page not found
		http.NotFound(writer, request)
	}

}

func artistHandler(writer http.ResponseWriter, request *http.Request) {
	name := strings.TrimPrefix(request.URL.Path, "/artists/")

	template, err := template.ParseFiles("templates/artists/artist.html")
	if err != nil {
		http.Error(writer, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// an array of artist data for each instance of artist struct
	artistsData, err := Artists()
	if err != nil {
		http.Error(writer, "Trouble loading artists", http.StatusInternalServerError)
		return
	}

	// find artist using name
	for _, artist := range artistsData {
		if artist.Name == name {
			relation, err := GetRelation(artist.Relations)
			if err != nil {
				http.Error(writer, "Trouble loading relation", http.StatusInternalServerError)
				return
			}

			// add instances of Artist and Relation struct to ArtistPageData
			data := ArtistPageData{
				Artist:   artist,
				Relation: relation,
			}

			// render html template
			if err := template.Execute(writer, data); err != nil {
				http.Error(writer, "Template execution error: "+err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}
	// 404 artist page not found
	http.NotFound(writer, request)
}

type ArtistPageData struct {
	Artist   Artist
	Relation Relation
}
