package main

import (
	"log"
	"net/http"
)

func main() {
	store := NewMemberStore()
	handler := NewMemberHandler(store)
	router := SetupRoutes(handler)

	log.Print("Starting server on :8080")

	http.ListenAndServe(":8080", router)

}
