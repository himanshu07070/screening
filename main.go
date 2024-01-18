package main

import (
	"log"
	"net/http"
)

func main() {
	//Renamed setupJsonApi to setupJSONAPI for consistency and adherence to Go naming conventions (camelCase)
	setupJSONAPI()

	// handle errors
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
