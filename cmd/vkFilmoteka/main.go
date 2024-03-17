package main

import (
	"github.com/localpurpose/vk-filmoteka/handlers"
	"github.com/localpurpose/vk-filmoteka/middleware"
	"github.com/localpurpose/vk-filmoteka/pkg/database/postgres"
	"log"
	"net/http"
)

func main() {

	// TODO implement method to get person by NAME;
	// TODO Person takes part movie inside his struct and DB model *gorm;
	// TODO Validate movie rating (1 - 10);
	// TODO Project refactoring after implementing all API methods;
	// TODO Roles permissions (See Task Description...)
	// TODO Swagger 2.0 specification - /doc folder
	// TODO Unit Testing min. 70% - 100%
	// TODO Logger beautify

	mux := http.NewServeMux()

	postgres.ConnectDB()

	// Person
	mux.HandleFunc("/person/create", middleware.LogHandler(middleware.Protected(handlers.CreatePerson)))
	mux.HandleFunc("/person/update", middleware.LogHandler(middleware.Protected(handlers.UpdatePerson)))
	mux.HandleFunc("/person/delete", middleware.LogHandler(middleware.Protected(handlers.DeletePerson)))
	mux.HandleFunc("/person", middleware.LogHandler(middleware.Protected(handlers.GetPerson)))

	// Movie
	mux.HandleFunc("/movie/create", middleware.LogHandler(middleware.Protected(handlers.CreateMovie)))
	mux.HandleFunc("/movie/update", middleware.LogHandler(middleware.Protected(handlers.UpdateMovie)))
	mux.HandleFunc("/movie/delete", middleware.LogHandler(middleware.Protected(handlers.DeleteMovie)))
	mux.HandleFunc("/movie", middleware.LogHandler(middleware.Protected(handlers.GetMovieByName)))

	// Users (Administrator and default user)
	// Roles deference implements with JWT claim Role: user, admin
	// After creating user its default role is user, to change - USE DB (See Task Description...)
	mux.HandleFunc("/user/sign-up", middleware.LogHandler(handlers.UserSignUp))
	mux.HandleFunc("/user/sign-in", middleware.LogHandler(handlers.UserSignIn))

	mux.HandleFunc("/", middleware.LogHandler(handleHome))

	log.Println("Server started!")
	log.Fatal(http.ListenAndServe(":3000", mux))

}

func handleHome(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello from main page"))
}
