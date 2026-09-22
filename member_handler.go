package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// MemberHandler contains the HTTP handlers responsible for working
// with members.
//
// The handler does not store the members itself. Instead, it receives
// a pointer to MemberStore and asks the store to perform operations
// such as adding, deleting, and updating points.
//
// This creates a separation of responsibilities:
//
// HTTP request
// ↓
// MemberHandler
// ↓
// MemberStore
// ↓
// Member data
type MemberHandler struct {
	store *MemberStore
}

// NewMemberHandler creates a new MemberHandler and connects it to
// the provided MemberStore.
//
// We pass the store into the handler instead of creating a new store
// inside the handler so that both parts of the application use the
// same collection of members.
func NewMemberHandler(store *MemberStore) *MemberHandler {
	return &MemberHandler{
		store: store,
	}
}

// GetMembers handles requests for retrieving all members.
//
// The handler asks the store for all members and then writes their
// information directly into the HTTP response.
//
// Example response:
//
// ID: 1
// Name: John
// Points: 20
// Last Updated: ...
func (h *MemberHandler) GetMembers(w http.ResponseWriter, r *http.Request) {
	// Ask the store for the current list of members.
	members := h.store.GetMembers()

	// Loop through every member returned by the store.
	//
	// "id" is the member's ID/key, while "member" contains
	// the actual Member struct.
	for id, member := range members {
		fmt.Fprintf(
			w,
			"ID: %d\nName: %s\nPoints: %d\nLast Updated: %s\n",
			id,
			member.Name,
			member.Points,
			member.LastUpdated,
		)
	}

}

// AddMember handles requests for creating a new member.
//
// The member's name is currently taken directly from the URL.
//
// For example:
//
// /add/John
//
// would attempt to create a member named "John".
func (h *MemberHandler) AddMember(w http.ResponseWriter, r *http.Request) {
	// Remove "/add/" from the beginning of the URL path.
	//
	// Example:
	// "/add/John" → "John"
	name := strings.TrimPrefix(r.URL.Path, "/add/")

	// A member must have a name.
	// If the URL did not contain a name, return HTTP 400 (Bad Request).
	if name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	// A "/" inside the name would make the URL structure ambiguous.
	//
	// For example:
	// "/add/John/Doe"
	//
	// is not considered a valid member name by this simple routing
	// approach.
	if strings.Contains(name, "/") {
		http.Error(w, "invalid Name or url", http.StatusBadRequest)
		return
	}

	// Pass responsibility for actually creating the member to the store.
	h.store.AddMember(name)

	// Tell the client that the member was successfully created.
	fmt.Fprintf(w, "%s has been added successfully!", name)

}

// DeleteMember handles requests for deleting a member by ID.
//
// Example:
//
// /delete/3
//
// attempts to delete the member whose ID is 3.
func (h *MemberHandler) DeleteMember(w http.ResponseWriter, r *http.Request) {
	// Remove "/delete/" from the URL path.
	//
	// Example:
	// "/delete/3" → "3"
	idStr := strings.TrimPrefix(r.URL.Path, "/delete/")

	// Make sure an ID was actually provided.
	if idStr == "" {
		http.Error(w, "Id of user is required", http.StatusBadRequest)
		return
	}

	// Convert the ID from a string taken from the URL into an integer.
	//
	// HTTP URLs contain text, so we need strconv.Atoi before we can
	// use the value as an integer.
	id, err := strconv.Atoi(idStr)

	// If conversion fails, the URL did not contain a valid integer ID.
	//
	// Example:
	// "/delete/abc" → conversion fails.
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Ask the store to delete the member.
	//
	// DeleteMember returns false if no member with that ID exists.
	del := h.store.DeleteMember(id)

	if !del {
		http.NotFound(w, r)
		return
	}

	// Tell the client that the member was successfully deleted.
	fmt.Fprintf(w, "member with %d was successfully deleted", id)

}

// AddPoints handles requests for adding points to a member.
//
// The current implementation adds a fixed 10 points.
//
// Example:
//
// /points/add/3
//
// adds 10 points to member ID 3.
func (h *MemberHandler) AddPoints(w http.ResponseWriter, r *http.Request) {
	// Remove "/points/add/" from the URL path.
	//
	// Example:
	// "/points/add/3" → "3"
	idStr := strings.TrimPrefix(r.URL.Path, "/points/add/")

	// Make sure an ID was provided.
	if idStr == "" {
		http.Error(w, "Id of user is required", http.StatusBadRequest)
		return
	}

	// Convert the ID from the URL from a string to an integer.
	id, err := strconv.Atoi(idStr)

	// If the ID isn't a valid number, return a 404 response.
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Ask the store to add the points.
	//
	// The store is responsible for actually changing the member's
	// point balance.
	added := h.store.AddPoints(id)

	// If the member doesn't exist, return 404 Not Found.
	if !added {
		http.NotFound(w, r)
		return
	}

	// Tell the client that the points were added successfully.
	fmt.Fprintf(w, "10 points successfully added to member %d", id)

}

// DeletePoints handles requests for subtracting points from a member.
//
// The current implementation subtracts a fixed 10 points.
//
// Example:
//
// /points/delete/3
//
// subtracts 10 points from member ID 3.
func (h *MemberHandler) DeletePoints(w http.ResponseWriter, r *http.Request) {
	// Remove "/points/delete/" from the URL path.
	//
	// Example:
	// "/points/delete/3" → "3"
	idStr := strings.TrimPrefix(r.URL.Path, "/points/delete/")

	// Make sure an ID was provided.
	if idStr == "" {
		http.Error(w, "Id of user is required", http.StatusBadRequest)
		return
	}

	// Convert the ID from the URL from a string to an integer.
	id, err := strconv.Atoi(idStr)

	// If the ID isn't a valid number, return 404 Not Found.
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Ask the store to subtract the points.
	subtracted := h.store.DeletePoints(id)

	// If the member doesn't exist, return 404 Not Found.
	if !subtracted {
		http.NotFound(w, r)
		return
	}

	// Tell the client that the points were successfully subtracted.
	fmt.Fprintf(w, "10 points successfully subtracted from member %d", id)

}
