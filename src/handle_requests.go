package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/at-alex-alves/Artist-Rest-API/src/models"
)

// handleRequests executes tasks based on the type of request and the passed data.
func handleRequests(allArtists []models.Artist) {

	// Handles requests that do not need specific artist Id.
	http.HandleFunc("/artist", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		// Gets all artists stored in the variable "allArtists".
		case "GET":
			json.NewEncoder(w).Encode(allArtists)

		// Creates a new artist.
		case "POST":
			var artist models.Artist

			if err := json.NewDecoder(r.Body).Decode(&artist); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			allArtists = append(allArtists, artist)

			json.NewEncoder(w).Encode(allArtists)

		// Requests different than GET and POST are not allowed.
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

			json.NewEncoder(w).Encode(&models.Artist{})

		// Deletes the specified artist
		case "DELETE":
			for index, item := range allArtists {
				if item.Id == urlParts[2] {
					allArtists = append(allArtists[:index], allArtists[index+1:]...)
					break
				}

				json.NewEncoder(w).Encode(allArtists)
			}

		// Requests different than GET and DELETE are not allowed.
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed!"))
			return
		}
	})
}
