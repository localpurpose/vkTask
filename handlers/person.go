package handlers

import (
	"encoding/json"
	"github.com/localpurpose/vk-filmoteka/models"
	"github.com/localpurpose/vk-filmoteka/pkg/database/postgres"
	"io"
	"log"
	"net/http"
)

func CreatePerson(w http.ResponseWriter, r *http.Request) {
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

	var person models.Person

	if err := json.Unmarshal(body, &person); err != nil {
		log.Println("error unmarshalling body")
		return
	}

	if err = postgres.DB.DB.Create(&person).Error; err != nil {
		log.Println("error while inserting to DB", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	// TODO Implement json returns

}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
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

	var person models.Person

	if err := json.Unmarshal(body, &person); err != nil {
		log.Println("error unmarshalling body")
		return
	}

	if err = postgres.DB.DB.Where("id = ?", person.ID).Updates(&person).Error; err != nil {
		log.Println("error while updating row in DB", err)
		return
	}
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		if _, err := w.Write([]byte("Method not allowed. Only DELETE requests.")); err != nil {
			log.Printf("error while writting response body %s", err)
		}
		return
	}

	var person models.Person
	personID := r.URL.Query()["id"]

	// TODO check if such user does not exists

	s := postgres.DB.DB.Delete(&person, personID)
	if s.Error != nil {
		log.Println("error while deleting row from DB", s.Error)
		return
	}
	w.Write([]byte("OK. User deleted"))
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		if _, err := w.Write([]byte("Method not allowed. Only DELETE requests.")); err != nil {
			log.Printf("error while writting response body %s", err)
		}
		return
	}

	personID := r.URL.Query()["id"]

	var person models.Person
	err := postgres.DB.DB.Where("id = ?", personID).First(&person).Error
	if err != nil {
		log.Println("Some error while getting user from DB", err)
		return
	}

	w.Write([]byte("OK."))
	b, err := json.Marshal(person)
	if err != nil {
		log.Println("some error while unmarshalling", err)
	}
	w.Write([]byte(b))
}
