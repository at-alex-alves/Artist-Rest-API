package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type artWork struct {
	Name string `json:"name,omitempty"`
	Year string `json:"country,omitempty"`
}

type birthplace struct {
	City    string `json:"city,omitempty"`
	Country string `json:"country,omitempty"`
}

type artist struct {
	Id             string      `json:"id,omitempty"`
	FirstName      string      `json:"firstName,omitempty"`
	LastName       string      `json:"lastName,omitempty"`
	MostFamousWork artWork    `json:"mostFamousWork,omitempty"`
	Birthplace     birthplace `json:"birthplace,omitempty"`
}

var allArtists []artist

func main() {
	allArtists = append(allArtists, artist{
		Id:        "1",
		FirstName: "David",
		LastName:  "Gilmour",
		MostFamousWork: artWork{
			Name: "Comfortably Numb - The Wall",
			Year: "1979",
		},
		Birthplace: birthplace{
			City:    "Cambridge",
			Country: "England",
		},
	})

	allArtists = append(allArtists, artist{
		Id:        "2",
		FirstName: "Antonio",
		LastName:  "Vivaldi",
		MostFamousWork: artWork{
			Name: "The Four Seasons",
			Year: "1723",
		},
		Birthplace: birthplace{
			City:    "Venice",
			Country: "Italy",
		},
	})

	// Handles requests that do not need specific artist Id.
	http.HandleFunc("/artist", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		// Gets all artists stored in the variable "allArtists"
		case "GET":
			json.NewEncoder(w).Encode(allArtists)

		// Creates a new artist
		case "POST":
			var artist artist

			if err := json.NewDecoder(r.Body).Decode(&artist); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			allArtists = append(allArtists, artist)

			json.NewEncoder(w).Encode(allArtists)

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed!"))
			return
		}
	})

	// Handles requests that need specific artist Id.
	http.HandleFunc("/artist/", func(w http.ResponseWriter, r *http.Request) {
		urlParts := strings.Split(r.URL.String(), "/")

		// Returns an error for wrong URLs
		if len(urlParts) != 3 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		switch r.Method {

		// Gets the specified artist
		case "GET":
			for _, item := range allArtists {
				if item.Id == urlParts[2] {
					json.NewEncoder(w).Encode(item)
					return
				}
			}

			json.NewEncoder(w).Encode(&artist{})

		// Deletes the specified artist
		case "DELETE":
			for index, item := range allArtists {
				if item.Id == urlParts[2] {
					allArtists = append(allArtists[:index], allArtists[index+1:]...)
					break
				}

				json.NewEncoder(w).Encode(allArtists)
			}

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed!"))
			return
		}
	})

	// Starts the server.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
