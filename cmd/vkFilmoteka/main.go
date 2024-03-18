package main

import (
	_ "github.com/localpurpose/vk-filmoteka/docs"
	"github.com/localpurpose/vk-filmoteka/handlers"
	"github.com/localpurpose/vk-filmoteka/middleware"
	"github.com/localpurpose/vk-filmoteka/pkg/database/postgres"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

// @title						Filmoteka API
// @version					1.0
// @description				VK TEST TASK - FILMOTEKA BACKEND REST API
// @host						localhost:3000
//
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {

	// TODO Person takes part movie inside his struct and DB model *gorm;
	// TODO Validate movie rating (1 - 10);

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
	mux.HandleFunc("/person", middleware.LogHandler(middleware.Protected(handlers.GetPersonByName)))
	mux.HandleFunc("/persons", middleware.LogHandler(middleware.Protected(handlers.GetAllPersons)))
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

	// Swagger 2.0 Specification
	mux.HandleFunc("/doc/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/swagger.json")
	})
	mux.HandleFunc("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3000/swagger/doc.json"),
	))

	log.Println("Server started!")
	log.Fatal(http.ListenAndServe(":3000", mux))

}

func handleHome(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello from main page"))
}
