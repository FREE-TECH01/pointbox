package main

import (
	"log"
	"net/http"
)

// main is the entry point of the PointBox application.
//
// This function is responsible for putting all the major pieces of the
// application together and starting the HTTP server.
//
// The application follows a simple flow:
//
// 1. Create the member store (data layer).
// 2. Create the member handler and give it access to the store.
// 3. Create the router and register the application's HTTP routes.
// 4. Start the HTTP server on port 8080.
//
// Keeping this wiring in main makes it easy to see how the different
// parts of the application depend on each other.
func main() {
	// Create the member store.
	//
	// The store is responsible for managing member data. We create it
	// first because the handler will need access to it.
	store := NewMemberStore()

	// Create the HTTP handler and give it the store.
	//
	// The handler is responsible for receiving HTTP requests, processing
	// them, and returning HTTP responses. It uses the store whenever it
	// needs to read or modify member data.
	handler := NewMemberHandler(store)

	// Create the application's router.
	//
	// SetupRoutes connects URL paths and HTTP methods to the appropriate
	// handler functions.
	router := SetupRoutes(handler)

	// Let us know that the server is about to start.
	log.Print("Starting server on :8080")

	// Start the HTTP server on port 8080.
	//
	// The router is passed to the server so that incoming requests are
	// directed to the correct handler based on the routes we configured.
	http.ListenAndServe(":8080", router)

}
