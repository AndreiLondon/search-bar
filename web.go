package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

const issuesURL = "https://groupietrackers.herokuapp.com/api"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/artist", Artist)
	http.HandleFunc("/error", Error)

	//Listen for request
	fmt.Println((fmt.Sprintf("Starting application on port %s", portNumber)))
	_ = http.ListenAndServe(portNumber, nil)
}
