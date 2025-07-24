package main

import (
	"fmt"
	"net/http"
)

func main() {
	runServer()

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
