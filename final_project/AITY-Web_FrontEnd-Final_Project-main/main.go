package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/handlers" // Import gorilla handlers for CORS
)

// ResponseData represents the structure of the server response
type ResponseData struct {
	Message string `json:"message"`
}

// Handler function to handle requests
func handler(w http.ResponseWriter, r *http.Request) {
	// Set response type to JSON
	w.Header().Set("Content-Type", "application/json")

	if r.URL.Path == "/success" {
		// Simulate a successful response
		response := ResponseData{Message: "Success!"}
		json.NewEncoder(w).Encode(response)
	} else if r.URL.Path == "/error" {
		// Simulate an error response
		http.Error(w, "Something went wrong!", http.StatusInternalServerError)
	} else {
		http.NotFound(w, r)
	}
}

func main3() {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Register the handler function for specific paths
	mux.HandleFunc("/success", handler)
	mux.HandleFunc("/error", handler)

	// Enable CORS for all origins (or you can restrict to specific origins)
	muxWithCors := handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(mux)

	// Start the server
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", muxWithCors)
}
