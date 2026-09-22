package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

//to handle the members
type MemberHandler struct {
	store *MemberStore
}

func NewMemberHandler(store *MemberStore) *MemberHandler {
	return &MemberHandler{
		store: store,
	}
}

func (h *MemberHandler) GetMembers(w http.ResponseWriter, r *http.Request) {
	members := h.store.GetMembers()

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


func (h *MemberHandler) AddMember(w http.ResponseWriter, r *http.Request){
	name := strings.TrimPrefix(r.URL.Path, "/add/")

	if name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	if strings.Contains(name, "/"){
		http.Error(w, "invalid Name or url", http.StatusBadRequest)
		return
	}

	h.store.AddMember(name)

	fmt.Fprintf(w, "%s has been added successfully!", name)
}

func (h *MemberHandler) DeleteMember(w http.ResponseWriter, r *http.Request){
	idStr := strings.TrimPrefix(r.URL.Path, "/delete/")

	if idStr == "" {
		http.Error(w, "Id of user is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	del := h.store.DeleteMember(id)

	if !del{
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "member with %d was successfully deleted", id)
}

func (h *MemberHandler) AddPoints(w http.ResponseWriter, r *http.Request){
	idStr := strings.TrimPrefix(r.URL.Path, "/points/add/")

	if idStr == "" {
		http.Error(w, "Id of user is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	added := h.store.AddPoints(id)

	if !added {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "10 point succesfully added to membe %d", id)
}

func (h *MemberHandler) DeletePoints(w http.ResponseWriter, r *http.Request){
	idStr := strings.TrimPrefix(r.URL.Path, "/points/delete/")

	if idStr == "" {
		http.Error(w, "Id of user is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	subtracted := h.store.DeletePoints(id)

	if !subtracted {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "10 point successfully subtracted from member %d", id)
}