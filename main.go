package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"
)

// Quote Struct (Model)
type Quote struct{
	ID	string `json:"id"`	
	Quote	string `json:"quote"`
	Reference string `json:"reference"`
	Owner string `json:"owner"`
}

// Init quotes var as slice Quote Struct
var mockQuotes []Quote

// Get all quotes
func getQuotes(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mockQuotes)
}

// Get single quote
func getQuote(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// Loop throug books and find with id
	for _, item := range mockQuotes{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Quote{})
}

// Create a quote
func createQuotes(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var quote Quote
	_ = json.NewDecoder(r.Body).Decode(&quote)
	quote.ID = strconv.Itoa(len(mockQuotes)+1)
	mockQuotes = append(mockQuotes, quote)
	json.NewEncoder(w).Encode(quote)

}

// Update a quote
func updateQuotes(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range mockQuotes{
		if item.ID == params["id"]{
			var quote Quote
			_ = json.NewDecoder(r.Body).Decode(&quote)
			quote.ID = item.ID
			mockQuotes[index] = quote;
			json.NewEncoder(w).Encode(quote)
			return
		}
	}
	json.NewEncoder(w).Encode(mockQuotes)
}

//Delete a quote
func deleteQuotes(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range mockQuotes{
		if item.ID == params["id"]{
			mockQuotes = append(mockQuotes[:index], mockQuotes[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(mockQuotes)
}




func main() {
	// Print console when server starting
	fmt.Println("Go Server Starting...") 

	// Init Router
	router := mux.NewRouter().StrictSlash(true)

	// Mock Data
	mockQuotes = append(mockQuotes, Quote{ID: "1", Quote: "May the Force be with you", Reference: "Starwars", Owner:"Jedi"})
	mockQuotes = append(mockQuotes, Quote{ID: "2", Quote: "Great leaders inspire greatness in others.", Reference: "Star Wars: The Clone Wars", Owner:"01×01 – Ambush -- opening quote"})
	mockQuotes = append(mockQuotes, Quote{ID: "3", Quote: "Belief is not a matter of choice, but of conviction.", Reference: "Star Wars: The Clone Wars", Owner:"01×02 – Rising Malevolence -- opening quote"})


	// Route Handler / Endpoints
	router.HandleFunc("/api/quotes", getQuotes).Methods("GET")
	router.HandleFunc("/api/quotes/{id}", getQuote).Methods("GET")
	router.HandleFunc("/api/quotes", createQuotes).Methods("POST")
	router.HandleFunc("/api/quotes/{id}", updateQuotes).Methods("PUT")
	router.HandleFunc("/api/quotes/{id}", deleteQuotes).Methods("DELETE")

	// Run the server and log.Fatal if it fail
	log.Fatal(http.ListenAndServe(":8080", router))
}