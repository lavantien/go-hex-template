package main

import (
	"fmt"
	"log"
	"net/http"

	"go-hex-template/pkg/adding"
	"go-hex-template/pkg/http/rest"
	"go-hex-template/pkg/listing"
	"go-hex-template/pkg/reviewing"
	"go-hex-template/pkg/storage/json"
	"go-hex-template/pkg/storage/memory"
)

// StorageType defines available storage types
type Type int

const (
	// JSON will store data in JSON files saved on disk
	JSON Type = iota
	// Memory will store data in memory
	Memory
)

func main() {

	// set up storage
	storageType := JSON // this could be a flag; hardcoded here for simplicity

	var adder adding.Service
	var lister listing.Service
	var reviewer reviewing.Service

	switch storageType {
	case Memory:
		s := new(memory.Storage)

		adder = adding.NewService(s)
		lister = listing.NewService(s)
		reviewer = reviewing.NewService(s)

	case JSON:
		// error handling omitted for simplicity
		s, _ := json.NewStorage()

		adder = adding.NewService(s)
		lister = listing.NewService(s)
		reviewer = reviewing.NewService(s)
	}

	// set up the HTTP server
	router := rest.Handler(adder, lister, reviewer)

	fmt.Println("The beer server is on tap now: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
