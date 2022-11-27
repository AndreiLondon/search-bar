package main

/*The "Groupies" struct contains "Artists", "Location", "Dates", and "Relation".
For the "Artists" , "ArtistNames" field, we use a field tag `json:"artists"` to help Go examine which field in the JSON is mapped to this field in the struct.
*/

type Groupies struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

// declare a custom type using struct
