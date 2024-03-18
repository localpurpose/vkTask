package handlers

import (
	"encoding/json"
	"github.com/localpurpose/vk-filmoteka/models"
	"github.com/localpurpose/vk-filmoteka/pkg/database/postgres"
	"io"
	"log"
	"net/http"
	"strings"
)

// CreateMovie godoc
//
// @Security ApiKeyAuth
//
//	@Summary	Creates movie from request body
//	@Tags		movies
//	@Accept		json
//	@Produce	json
//	@Success	200
//	@Router		/movie/create [post]
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		newErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed. Only POST requests.")
		return
	}

	// Create Person logic implementation
	body, err := io.ReadAll(r.Body)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	var movie models.Movie

	if err = json.Unmarshal(body, &movie); err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err = postgres.DB.DB.Create(&movie).Error; err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newJsonResponse(w, http.StatusOK, map[string]interface{}{
		"ID":          movie.ID,
		"Name":        movie.Name,
		"Description": movie.Description,
		"Date":        movie.Date,
		"Rating":      movie.Rating,
	})

}

// UpdateMovie @Summary	Updates movie from request body by url path id
// @Security ApiKeyAuth
// @Tags		movies
// @Accept		json
// @Produce	json
// @Success	200
// @Router		/movie/update/ [patch]
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		newErrorResponse(w, http.StatusMethodNotAllowed, "Method not Allowed. Only PATCH requests.")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	var movie models.Movie
	movieID := r.URL.Query()["id"]

	if err = json.Unmarshal(body, &movie); err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err = postgres.DB.DB.Where("id = ?", movieID).Updates(&movie).Error; err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newJsonResponse(w, http.StatusOK, map[string]interface{}{
		"ID":          movie.ID,
		"Name":        movie.Name,
		"Description": movie.Description,
		"Date":        movie.Date,
		"Rating":      movie.Rating,
	})
}

// DeleteMovie godoc
//
// @Security ApiKeyAuth
// @Summary	Updates movie from request body by url path id
// @Tags		movies
// @Accept		json
// @Produce	json
// @Success	200
// @Router		/movie/delete/ [delete]
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		newErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed. Only DELETE requests.")
		return
	}

	var movie models.Movie
	movieID := r.URL.Query()["id"]

	// TODO check if movie user does not exists

	s := postgres.DB.DB.Delete(&movie, movieID)
	if s.Error != nil {
		newErrorResponse(w, http.StatusInternalServerError, s.Error.Error())
		return
	}
	newJsonResponse(w, http.StatusOK, map[string]interface{}{
		"message": "user deleted - ok",
	})
}

// GetMovieByName godoc
//
// @Security ApiKeyAuth
// @Summary	Updates movie from request body by Name
// @Tags		movies
// @Accept		json
// @Produce	json
// @Success	200
// @Router		/movie [get]
func GetMovieByName(w http.ResponseWriter, r *http.Request) {

	// TODO Implement sorting by: (ORDER BY): name,rating,date (default:rating)

	if r.Method != http.MethodGet {
		newErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed. Only GET requests.")
		return
	}

	movieReq := r.URL.Query()["name"]
	ssd := r.URL.Query()["ssd"]
	log.Println("ssd--->", ssd)

	var movie []models.Movie

	err := postgres.DB.DB.Select("name").Find(&movie).Error
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	for i := 0; i < len(movie); i++ {
		if strings.Contains(movie[i].Name, movieReq[0]) {
			log.Println("MATCH:", movie[i].Name)
		}
	}

}
