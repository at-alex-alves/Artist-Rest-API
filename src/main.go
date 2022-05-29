package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/at-alex-alves/Artist-Rest-API/src/models"
)

var allArtists []models.Artist

func main() {
	// Loads the data that is going to be used in the example.
	file, err := ioutil.ReadFile("./src/example/example_data.json")
	if err != nil {
		panic(err)
	}

	json_data := []models.Artist{}

	// Parses the JSON data and stores the result in the json_data variable.
	if err := json.Unmarshal([]byte(file), &json_data); err != nil {
		panic(err)
	}

	allArtists = append(allArtists, json_data...)

	handleRequests(allArtists)

	// Starts the server.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
