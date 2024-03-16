package main

import (
	"github.com/localpurpose/vk-filmoteka/handlers"
	"github.com/localpurpose/vk-filmoteka/pkg/database/postgres"
	"log"
	"net/http"
)

func main() {

	postgres.ConnectDB()
	// Person
	// TODO Methods Patch, Get, Delete should have an ID in URL-PATH
	http.HandleFunc("/person/create", handlers.CreatePerson)
	http.HandleFunc("/person/update", handlers.UpdatePerson)
	http.HandleFunc("/person/delete", handlers.DeletePerson)
	http.HandleFunc("/person", handlers.GetPerson)

	http.HandleFunc("/", handleHome)

	log.Fatal(http.ListenAndServe(":3000", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from main page"))
}
