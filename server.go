package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func runServer() {
	http.HandleFunc("/", homeHandler)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

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

	// artistNames := make(map[int]string)

	// for id, artist := range artistsData {
	// 	artistNames[id] = artist.Name
	// }

	template.Execute(writer, artistsData)

}
