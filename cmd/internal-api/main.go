package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Define your expected request structure
type RequestData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Define your response structure
type ResponseData struct {
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data RequestData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	response := ResponseData{
		Message: "Hello, " + data.Name + "! You are " + string(rune(data.Age)) + " years old.",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/api/greet", handler)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
