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

	mux := http.NewServeMux()

	postgres.ConnectDB()
	// Person
	mux.HandleFunc("/person/create", handlers.CreatePerson)

	mux.HandleFunc("/person/update", handlers.UpdatePerson)
	mux.HandleFunc("/person/delete", handlers.DeletePerson)
	mux.HandleFunc("/person", handlers.GetPerson)

	// Movie
	mux.HandleFunc("/movie/create", handlers.CreateMovie)
	mux.HandleFunc("/movie/update", handlers.UpdateMovie)
	mux.HandleFunc("/movie/delete", handlers.DeleteMovie)
	mux.HandleFunc("/movie", handlers.GetMovie)

	// Actors - Relations between movies and persons
	// TODO Implementation actors API (See Task Description)
	mux.HandleFunc("/", handleHome)

	// Users (Administrator and default user)
	// Roles deference implements with JWT claim Role: user, admin
	mux.HandleFunc("/user/sign-up", handlers.UserSignUp)
	mux.HandleFunc("/user/sign-in", handlers.UserSignIn)

	log.Println("Server started!")
	log.Fatal(http.ListenAndServe(":3000", mux))
}

func handleHome(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello from main page"))
}
