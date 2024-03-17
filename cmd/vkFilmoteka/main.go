package main

import (
	"github.com/localpurpose/vk-filmoteka/handlers"
	"github.com/localpurpose/vk-filmoteka/middleware"
	"github.com/localpurpose/vk-filmoteka/pkg/database/postgres"
	"log"
	"net/http"
)

func main() {

	// TODO Methods Patch, Get, Delete should have an ID in URL-PATH
	// TODO On-change data should return a new model as a JSON object
	// TODO Handlers refactoring
	// TODO Project refactoring after implementing all API methods
	// TODO Add salt when hashing user password
	// TODO implement method to get person by NAME
	// TODO implement method to get movie by NAME

	mux := http.NewServeMux()

	postgres.ConnectDB()

	// Person
	mux.HandleFunc("/person/create", middleware.LogHandler(middleware.Protected(handlers.CreatePerson)))
	mux.HandleFunc("/person/update", middleware.LogHandler(middleware.Protected(handlers.UpdatePerson)))
	mux.HandleFunc("/person/delete", middleware.LogHandler(middleware.Protected(handlers.DeletePerson)))
	mux.HandleFunc("/person", middleware.LogHandler(middleware.Protected(handlers.GetPerson)))

	// Movie
	mux.HandleFunc("/movie/create", middleware.LogHandler(middleware.Protected(handlers.CreateMovie)))
	mux.HandleFunc("/movie/update", middleware.LogHandler(middleware.Protected(handlers.CreateMovie)))
	mux.HandleFunc("/movie/delete", middleware.LogHandler(middleware.Protected(handlers.DeleteMovie)))
	mux.HandleFunc("/movie", middleware.LogHandler(middleware.Protected(handlers.GetMovieByName)))

	// Actors - Relations between movies and persons
	// TODO Implementation actors API (See Task Description)
	mux.HandleFunc("/", middleware.LogHandler(handleHome))

	// Users (Administrator and default user)
	// Roles deference implements with JWT claim Role: user, admin
	mux.HandleFunc("/user/sign-up", middleware.LogHandler(handlers.UserSignUp))
	mux.HandleFunc("/user/sign-in", middleware.LogHandler(handlers.UserSignIn))

	log.Println("Server started!")
	log.Fatal(http.ListenAndServe(":3000", mux))
}

func handleHome(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello from main page"))
}
