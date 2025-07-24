package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		// TODO: add 404 page template
	}

	template, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(writer, "Error parsing index.html template", http.StatusInternalServerError)
		return
	}

	artistsData, err := Artists()
	if err != nil {
		http.Error(writer, "Trouble loading artists", http.StatusInternalServerError)
	}

	artistNames := make([]string, len(artistsData))

	for i, artist := range artistsData {
		artistNames[i] = artist.Name
	}

	template.Execute(writer, artistNames)

}
