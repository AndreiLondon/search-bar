package main

import (
	"log"
	"net/http"
	"text/template"
)

const portNumber = ":8080"

const apiURL = "https://groupietrackers.herokuapp.com/api"
const artURL = "https://groupietrackers.herokuapp.com/api/artists"
const locURL = "https://groupietrackers.herokuapp.com/api/locations"
const datesURL = "https://groupietrackers.herokuapp.com/api/dates"
const relURL = "https://groupietrackers.herokuapp.com/api/relation"

/*
w - writer - a value for updating the response that will be sent to the browser
where we about to write the result

r - request - a value representing the request from the browser
where all parameters are existing
*/

func getHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	/* Use the template.ParseFiles() function to read the template file into a
	   template set. If there's an error, we log the detailed error message and use // the http.Error() function to send a generic 500 Internal Server Error
	   response to the user.
	*/

	temp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	/* We then use the Execute() method on the template set to write the
	   template content as the response body. The last parameter to Execute()
	   represents any dynamic data that we want to pass in, which for now we'll
	   leave as nil.
	*/
	temp.Execute(w, bands[0])
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

	//renderTemplate(w, "index.html")
}

func getError(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/error" {
		http.NotFound(w, r)
		return

	}
	temp, err := template.ParseFiles("templates/error.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	temp.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func getArtist(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artist" {
		http.NotFound(w, r)
		return

	}
	temp, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	temp.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
