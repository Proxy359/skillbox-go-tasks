package main

import (
	"cde/service"
	"cde/storage"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	store := &storage.Storage{Store: make(map[int]*storage.User)}
	service := service.Service{Storage: store}
	mux.HandleFunc("/create", service.Create)
	mux.HandleFunc("/make_friends", service.MakeFriends)
	mux.HandleFunc("/deliete", service.DelieteUser)
	mux.HandleFunc("/friends/", service.GetFriends)
	mux.HandleFunc("/newAge/", service.NewAge)
	mux.HandleFunc("/allUsers", service.GetUsers)

	http.ListenAndServe("localhost:8082", mux)
}
