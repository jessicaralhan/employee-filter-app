package main

import (
	"log"
	"net/http"
)

func main() {
	// Initialize database
	InitDB()
	defer db.Close()

	// Register the handler
	http.HandleFunc("/employees", GetEmployees)

	// Start the server
	log.Println("Server is running on http://localhost:8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
