package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	// Print console when server starting
	fmt.Println("Go Server Starting...") 

	// Init Router
	router := mux.NewRouter().StrictSlash(true)

	// Route Handler / Endpoints
	router.HandleFunc("/", homeLink).Methods("GET")

	// Run the server and log.Fatal if it fail
	log.Fatal(http.ListenAndServe(":8080", router))
}