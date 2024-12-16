package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Request struct represents the structure of the incoming JSON request
type Request struct {
	Message string `json:"message"`
}

// Response struct represents the structure of the outgoing JSON response
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	var req Request

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		sendErrorResponse(w, "Invalid JSON format")
		return
	}

	if req.Message == "" {
		sendErrorResponse(w, "Invalid JSON message")
		return
	}

	fmt.Printf("Received message: %s\n", req.Message)

	sendSuccessResponse(w, "Data successfully received")
}

func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Status:  "success",
		Message: "GET request successful",
	}
	sendJSONResponse(w, response)
}

func sendSuccessResponse(w http.ResponseWriter, message string) {
	response := Response{
		Status:  "success",
		Message: message,
	}
	sendJSONResponse(w, response)
}

func sendErrorResponse(w http.ResponseWriter, message string) {
	response := Response{
		Status:  "fail",
		Message: message,
	}
	sendJSONResponse(w, response)
}

func sendJSONResponse(w http.ResponseWriter, response Response) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main1() {
	http.HandleFunc("/post", handlePostRequest)
	http.HandleFunc("/get", handleGetRequest)

	fmt.Println("Server is listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
