package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	groupies  Groupies
	bands     []Artists
	locations Locations
	dates     Dates
	relations Relations
)

func main() {
	//API endpoint that returns a JSON string our program performs a GET request.
	//GET request
	resApi, err := http.Get(apiURL)

	/*
		Within our main function we first query our API endpoint,
		we map the results of this into either response or err
		and then check to see if err is nil. If it is we exit.
	*/

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	/*
		Then perform the conversion of our response’s bodyresponse, err := http.Get(apiURL)
		from bytes into something meaningful that can be printed out in the console.
		First we use ioutil.ReadAll(response.Body) to read in data from the incoming byte stream.
		Then convert this []byte response into a string using string(responseData) within our print statement.
	*/
	apiData, err := io.ReadAll(resApi.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(apiData, &groupies)
	// defer
	// defer apiData.Close()

	resArt, err := http.Get(groupies.Artists)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	artData, err := io.ReadAll(resArt.Body)
	if err != nil {
		log.Fatal(err)
	}
	// defer
	json.Unmarshal(artData, &bands)

	resLoc, err := http.Get(groupies.Locations)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	locData, err := io.ReadAll(resLoc.Body)
	if err != nil {
		log.Fatal(err)
	}
	// defer

	resDates, err := http.Get(groupies.Dates)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	datesData, err := io.ReadAll(resDates.Body)
	if err != nil {
		log.Fatal(err)
	}
	// defer

	resRel, err := http.Get(groupies.Relation)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	relData, err := io.ReadAll(resRel.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(locData, &locations)
	json.Unmarshal(datesData, &dates)
	json.Unmarshal(relData, &relations)

	/*
		"/" root element
		If we receive a request for a URL ending - "/", "/artist", "/error"
		Home, Artist, Error names of the func that generate a response.
	*/
	http.HandleFunc("/", getHome)
	http.HandleFunc("/artist", getArtist)
	http.HandleFunc("/error", getError)

	// Fixing Static data
	staticHandler := http.StripPrefix(
		"/static/",
		http.FileServer(http.Dir("./static")),
	)
	http.Handle("/static/", staticHandler)
	// _ = http.ListenAndServe(portNumber, nil)

	//Listen for request
	fmt.Println((fmt.Sprintf("Starting application on port %s", portNumber)))
	http.ListenAndServe(portNumber, nil)

}
