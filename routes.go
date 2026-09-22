package main

import "net/http"

func SetupRoutes(handler *MemberHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.GetMembers)
	mux.HandleFunc("/add/", handler.AddMember)
	mux.HandleFunc("/delete/", handler.DeleteMember)
	mux.HandleFunc("/points/add/", handler.AddPoints)
	mux.HandleFunc("/points/delete/", handler.DeletePoints)

	return mux
}