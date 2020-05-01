package main

import (
	"os"
	"log"
	"net/http"
	"encoding/json"
	"context"
	"time"
	"cloud.google.com/go/datastore"
	"github.com/gorilla/mux"
	"github.com/emre/react-golang-web-app/repository"
	"github.com/emre/react-golang-web-app/model"
)


// Init quotes var as slice Quote Struct
var datastoreClient *datastore.Client

// Get all quotes
func getQuotes(w http.ResponseWriter,r *http.Request){
	ctx := context.Background()

	quotes, err := repository.GetAllQuotes(ctx, datastoreClient)
	if err != nil {
		log.Printf("Failed to fetch quote list: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quotes)
}

// Get single quote
func getQuote(w http.ResponseWriter,r *http.Request){	
	ctx := context.Background()
	params := mux.Vars(r)

	quote, dbErr := repository.GetSingleQuote(ctx, datastoreClient, params["id"])
	if dbErr != nil {
		log.Printf("Failed to fetch the quote: %v", dbErr)
		http.Error(w, dbErr.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&quote)
}

// Create a quote
func createQuotes(w http.ResponseWriter,r *http.Request){
	ctx := context.Background()

	var quote model.Quote	
	if decodeErr := json.NewDecoder(r.Body).Decode(&quote) ; decodeErr != nil {
		log.Printf("Failed to decode params: %v", decodeErr)
		http.Error(w, decodeErr.Error(), http.StatusBadRequest)
		
	}
	quote.Created = time.Now();

	savedQuote, dbErr := repository.CreateQuote(ctx, datastoreClient, quote)
	if dbErr != nil {
		log.Printf("Failed to write the quote: %v", dbErr)
		http.Error(w, dbErr.Error(), http.StatusInternalServerError)
	}

	 w.Header().Set("Content-Type", "application/json")
	 json.NewEncoder(w).Encode(savedQuote)

}

//Delete a quote
func deleteQuotes(w http.ResponseWriter,r *http.Request){
	ctx := context.Background()
	params := mux.Vars(r)	

	dbErr := repository.DeleteTask(ctx, datastoreClient, params["id"])
	if dbErr != nil {
		log.Printf("Failed to delete the quote: %v", dbErr)
		http.Error(w, dbErr.Error(), http.StatusInternalServerError)
		
	}
}

//Update a quote
func updateQuotes(w http.ResponseWriter,r *http.Request){
	ctx := context.Background()
	var quote model.Quote
	//params := mux.Vars(r)
	decodeErr := json.NewDecoder(r.Body).Decode(&quote)
	if(decodeErr != nil){
		log.Println(decodeErr)
		http.Error(w, decodeErr.Error(), http.StatusBadRequest)
	}

	//quote.Key = params["id"]
	savedQuote, dbErr := repository.UpdateQuote(ctx, datastoreClient, quote)
	if dbErr != nil {
		log.Printf("Failed to write quote: %v", dbErr)
		http.Error(w, dbErr.Error(), http.StatusInternalServerError)
	}

	 w.Header().Set("Content-Type", "application/json")
	 json.NewEncoder(w).Encode(savedQuote)

}

// Create mock data if DB empty
func dbInitializer(){
	ctx := context.Background()

	quotes, dbErr := repository.GetAllQuotes(ctx, datastoreClient)
	if dbErr != nil {
		log.Printf("Failed to fetch quote list: %v", dbErr)
	}

	// I try to putAll batch operation but keep getting error 
	// therefor create each item by iteration but this is not perefered
	if len(quotes) <= 0 {
		log.Printf("db initialized")
		if _, dbErr := repository.CreateQuote(ctx, datastoreClient, model.Quote{
			Quote:"May the Force be with you",
			Reference:"Starwars",
			Owner:"Jedi",
			Created: time.Now(),
		}); dbErr != nil {
			log.Printf("Failed to write quote: %v", dbErr)
		}

		if _, dbErr := repository.CreateQuote(ctx, datastoreClient, model.Quote{
			Quote:"Great leaders inspire greatness in others.",
			Reference:"Star Wars: The Clone Wars",
			Owner:"01×01 – Ambush -- opening quote",
			Created: time.Now(),
		}); dbErr != nil {
			log.Printf("Failed to write quote: %v", dbErr)
		}

		if _, dbErr := repository.CreateQuote(ctx, datastoreClient, model.Quote{
			Quote:"Belief is not a matter of choice, but of conviction.",
			Reference:"Star Wars: The Clone Wars",
			Owner:"01×02 – Rising Malevolence -- opening quote",
			Created: time.Now(),
		}); dbErr != nil {
			log.Printf("Failed to write quote: %v", dbErr)
		}
	}
}

func main() {
	// Describe the project
	projID := os.Getenv("DATASTORE_PROJECT_ID")
	if projID == "" {
		log.Fatal(`You need to set the environment variable "DATASTORE_PROJECT_ID"`)
	}

	//datastore_build_service
	ctx := context.Background()
	var err error
	datastoreClient,err = datastore.NewClient(ctx, projID)
	if err != nil {
		log.Fatalf("Could not create datastore client: %v", err)
	}

	// Create mock data if DB empty
	dbInitializer()


	// Init Router
	router := mux.NewRouter().StrictSlash(true)

	// Route Handler & Endpoints
	router.HandleFunc("/api/quotes", getQuotes).Methods("GET")
	router.HandleFunc("/api/quotes/{id}", getQuote).Methods("GET")
	router.HandleFunc("/api/quotes", createQuotes).Methods("POST")
	router.HandleFunc("/api/quotes", updateQuotes).Methods("PUT")
	router.HandleFunc("/api/quotes/{id}", deleteQuotes).Methods("DELETE")

	// Print console when server starting
	log.Print("Listening on :8080")
	// Run the server and log.Fatal if it fail
	log.Fatal(http.ListenAndServe(":8080", router))
}