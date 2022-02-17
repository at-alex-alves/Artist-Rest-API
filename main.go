package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Artist struct {
	ID             string      `json:"id,omitempty"`
	Firstname      string      `json:"firstname,omitempty"`
	Lastname       string      `json:"lastname,omitempty"`
	MostFamousWork *ArtWork    `json:"mostfamouswork,omitempty"`
	Birthplace     *Birthplace `json:"birthplace,omitempty"`
}

type ArtWork struct {
	Name string `json:"name,omitempty"`
	Year string `json:"country,omitempty"`
}

type Birthplace struct {
	City    string `json:"city,omitempty"`
	Country string `json:"country,omitempty"`
}

var allArtists []Artist

func Artists(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {

	// Gets all artists stored in the variable "allArtists"
	case "GET":
		json.NewEncoder(writer).Encode(allArtists)

	// Creates a new artist
	case "POST":
		var artist Artist

		err := json.NewDecoder(request.Body).Decode(&artist)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		allArtists = append(allArtists, artist)

		json.NewEncoder(writer).Encode(allArtists)

	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
		writer.Write([]byte("Method not allowed!"))
		return
	}
}

func FindSpecificArtist(writer http.ResponseWriter, request *http.Request) {
	urlParts := strings.Split(request.URL.String(), "/")

	// Returns an error for wrong URLs
	if len(urlParts) != 3 {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	switch request.Method {

	// Gets the specified artist
	case "GET":
		for _, item := range allArtists {
			if item.ID == urlParts[2] {
				json.NewEncoder(writer).Encode(item)
				return
			}
		}

		json.NewEncoder(writer).Encode(&Artist{})

	// Deletes the specified artist
	case "DELETE":
		for index, item := range allArtists {
			if item.ID == urlParts[2] {
				allArtists = append(allArtists[:index], allArtists[index+1:]...)
				break
			}

			json.NewEncoder(writer).Encode(allArtists)
		}

	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
		writer.Write([]byte("Method not allowed!"))
		return
	}
}

func main() {
	allArtists = append(allArtists, Artist{
		ID:        "1",
		Firstname: "David",
		Lastname:  "Gilmour",
		MostFamousWork: &ArtWork{
			Name: "Comfortably Numb - The Wall",
			Year: "1979",
		},
		Birthplace: &Birthplace{
			City:    "Cambridge",
			Country: "England",
		},
	})

	allArtists = append(allArtists, Artist{
		ID:        "2",
		Firstname: "Antonio",
		Lastname:  "Vivaldi",
		MostFamousWork: &ArtWork{
			Name: "The Four Seasons",
			Year: "1723",
		},
		Birthplace: &Birthplace{
			City:    "Venice",
			Country: "Italy",
		},
	})

	http.HandleFunc("/artist", Artists)
	http.HandleFunc("/artist/", FindSpecificArtist)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
