package main

import (
	_ "github.com/localpurpose/vk-filmoteka/docs"
	handlers2 "github.com/localpurpose/vk-filmoteka/internal/handlers"
	"github.com/localpurpose/vk-filmoteka/pkg/database/postgres"
	middleware2 "github.com/localpurpose/vk-filmoteka/pkg/middleware"
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
	mux.HandleFunc("/person/create", middleware2.LogHandler(middleware2.Protected(handlers2.CreatePerson)))
	mux.HandleFunc("/person/update", middleware2.LogHandler(middleware2.Protected(handlers2.UpdatePerson)))
	mux.HandleFunc("/person/delete", middleware2.LogHandler(middleware2.Protected(handlers2.DeletePerson)))
	mux.HandleFunc("/person", middleware2.LogHandler(middleware2.Protected(handlers2.GetPersonByName)))
	mux.HandleFunc("/persons", middleware2.LogHandler(middleware2.Protected(handlers2.GetAllPersons)))
	// Movie
	mux.HandleFunc("/movie/create", middleware2.LogHandler(middleware2.Protected(handlers2.CreateMovie)))
	mux.HandleFunc("/movie/update", middleware2.LogHandler(middleware2.Protected(handlers2.UpdateMovie)))
	mux.HandleFunc("/movie/delete", middleware2.LogHandler(middleware2.Protected(handlers2.DeleteMovie)))
	mux.HandleFunc("/movie", middleware2.LogHandler(middleware2.Protected(handlers2.GetMovieByName)))

	// Users (Administrator and default user)
	// Roles deference implements with JWT claim Role: user, admin
	// After creating user its default role is user, to change - USE DB (See Task Description...)
	mux.HandleFunc("/user/sign-up", middleware2.LogHandler(handlers2.UserSignUp))
	mux.HandleFunc("/user/sign-in", middleware2.LogHandler(handlers2.UserSignIn))

	mux.HandleFunc("/", middleware2.LogHandler(handleHome))

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
