package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const portNumber = ":8080"
const apiURL = "https://groupietrackers.herokuapp.com/api"

/*
w - writer - a value for updating the response that will be sent to the browser
where we about to write the result

r - request - a value representing the request from the browser
where all parameters are existing
*/

func Home(w http.ResponseWriter, r *http.Request) {
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
	temp.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

	//renderTemplate(w, "index.html")
}

func main() {
	/*
		"/" root element
		If we receive a request for a URL ending - "/", "/artist", "/error"
		Home, Artist, Error names of the func that generate a response.
	*/
	http.HandleFunc("/", Home)

	//Fixing Static data
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
