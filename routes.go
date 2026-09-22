package main

import "net/http"

// SetupRoutes creates and configures the HTTP router for the application.
//
// A router's job is to look at an incoming HTTP request and decide
// which handler should process it.
//
// The handlers themselves contain the application logic, while this
// file is responsible for connecting URL paths to those handlers.
//
// Current routes:
//
// GET / → GetMembers
// /add/{name} → AddMember
// /delete/{id} → DeleteMember
// /points/add/{id} → AddPoints
// /points/delete/{id} → DeletePoints
func SetupRoutes(handler *MemberHandler) *http.ServeMux {
	// Create a new HTTP request multiplexer (router).
	mux := http.NewServeMux()

	// Register the route for retrieving members.
	//
	// A request to "/" will be handled by GetMembers.
	mux.HandleFunc("/", handler.GetMembers)

	// Register the route for creating a member.
	//
	// Example:
	// /add/John
	//
	// The AddMember handler extracts "John" from the URL.
	mux.HandleFunc("/add/", handler.AddMember)

	// Register the route for deleting a member.
	//
	// Example:
	// /delete/3
	//
	// The DeleteMember handler extracts "3" from the URL.
	mux.HandleFunc("/delete/", handler.DeleteMember)

	// Register the route for adding 10 points to a member.
	//
	// Example:
	// /points/add/3
	mux.HandleFunc("/points/add/", handler.AddPoints)

	// Register the route for subtracting 10 points from a member.
	//
	// Example:
	// /points/delete/3
	mux.HandleFunc("/points/delete/", handler.DeletePoints)

	// Return the configured router to main().
	//
	// main() then passes this router to http.ListenAndServe().
	return mux

}
