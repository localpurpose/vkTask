package handlers

import (
	"encoding/json"
	"github.com/localpurpose/vk-filmoteka/models"
	"github.com/localpurpose/vk-filmoteka/pkg/database/postgres"
	"io"
	"log"
	"net/http"
)

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		if _, err := w.Write([]byte("Method not allowed. Only POST requests.")); err != nil {
			log.Printf("error while writting response body %s", err)
		}
		return
	}

	// Create Person logic implementation
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading body", err)
		return
	}

	var movie models.Movie

	if err = json.Unmarshal(body, &movie); err != nil {
		log.Println("error unmarshalling body")
		return
	}

	if err = postgres.DB.DB.Create(&movie).Error; err != nil {
		log.Println("error while inserting to DB", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	// TODO Implement json returns

}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		w.WriteHeader(http.StatusMethodNotAllowed)
		if _, err := w.Write([]byte("Method not allowed. Only PATCH requests.")); err != nil {
			log.Printf("error while writting response body %s", err)
		}
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading body", err)
		return
	}

	var movie models.Movie
	movieID := r.URL.Query()["id"]

	if err = json.Unmarshal(body, &movie); err != nil {
		log.Println("error unmarshalling body")
		return
	}

	if err = postgres.DB.DB.Where("id = ?", movieID).Updates(&movie).Error; err != nil {
		log.Println("error while updating row in DB", err)
		return
	}
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		if _, err := w.Write([]byte("Method not allowed. Only DELETE requests.")); err != nil {
			log.Printf("error while writting response body %s", err)
		}
		return
	}

	var movie models.Movie
	movieID := r.URL.Query()["id"]

	// TODO check if such user does not exists

	s := postgres.DB.DB.Delete(&movie, movieID)
	if s.Error != nil {
		log.Println("error while deleting row from DB", s.Error)
		return
	}
	w.Write([]byte("OK. Movie deleted"))
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		if _, err := w.Write([]byte("Method not allowed. Only DELETE requests.")); err != nil {
			log.Printf("error while writting response body %s", err)
		}
		return
	}

	movieID := r.URL.Query()["id"]

	var movie models.Movie
	err := postgres.DB.DB.Where("id = ?", movieID).First(&movie).Error
	if err != nil {
		log.Println("Some error while getting movie from DB", err)
		return
	}

	b, err := json.Marshal(movie)
	if err != nil {
		log.Println("some error while unmarshalling", err)
	}
	w.Write([]byte(b))
}
