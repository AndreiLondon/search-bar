package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Get the JSON response from the URL.
func getData(apiURL string) (string, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	// Read the body of the response into []byte.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Declare a variable of type .
	var art Groupies
	// Unmarshal the JSON data into the variable.
	if err := json.Unmarshal(body, &art); err != nil {
		log.Fatal(err)
	}
	// Print the first station.
	fmt.Printf("%+v\n\n", art.Artists[0])
	return "", err
}
