package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func runServer() {
	http.HandleFunc("/browse", browseHandler)
	http.HandleFunc("/artists/", artistHandler)

	// link css
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

func artistHandler(writer http.ResponseWriter, request *http.Request) {
	name := strings.TrimPrefix(request.URL.Path, "/artists/")

	template, err := template.ParseFiles("templates/artists/artist.html")
	if err != nil {
		http.Error(writer, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// an array of artist data for each artist struct
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

			data := ArtistPageData{
				Artist:   artist,
				Relation: relation,
			}

			log.Printf("Relation data for %s: %+v\n", artist.Name, relation)
			log.Printf("DatesLocations length: %d\n", len(relation.DatesLocations))

			if err := template.Execute(writer, data); err != nil {
				http.Error(writer, "Template execution error: "+err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}

	http.NotFound(writer, request)
}

type ArtistPageData struct {
	Artist   Artist
	Relation Relation
}
