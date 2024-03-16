package main

import (
	"github.com/localpurpose/vk-filmoteka/handlers"
	"github.com/localpurpose/vk-filmoteka/pkg/database/postgres"
	"log"
	"net/http"
)

func main() {

	// TODO Methods Patch, Get, Delete should have an ID in URL-PATH - DONE
	// TODO On-change data should return a new model as a JSON object
	// TODO Handlers refactoring
	// TODO Project refactoring after implementing all API methods
	// TODO Add salt when hashing user password
	// TODO implement method to get person by NAME

	postgres.ConnectDB()
	// Person
	http.HandleFunc("/person/create", handlers.CreatePerson)
	http.HandleFunc("/person/update", handlers.UpdatePerson)
	http.HandleFunc("/person/delete", handlers.DeletePerson)
	http.HandleFunc("/person", handlers.GetPerson)

	// Movie
	http.HandleFunc("/movie/create", handlers.CreateMovie)
	http.HandleFunc("/movie/update", handlers.UpdateMovie)
	http.HandleFunc("/movie/delete", handlers.DeleteMovie)
	http.HandleFunc("/movie", handlers.GetMovie)

	// Actors - Relations between movies and persons
	// TODO Implementation actors API (See Task Description)
	http.HandleFunc("/", handleHome)

	// Users (Administrator and default user)
	// Roles deference implements with JWT claim Role: user, admin
	http.HandleFunc("/user/sign-up", handlers.UserSignUp)
	http.HandleFunc("/user/sign-in", handlers.UserSignIn)

	log.Fatal(http.ListenAndServe(":3000", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from main page"))
}
