package handlers

import (
	"encoding/json"
	"github.com/localpurpose/vk-filmoteka/models"
	"github.com/localpurpose/vk-filmoteka/pkg/database/postgres"
	"io"
	"net/http"
	"strconv"
)

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		newErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed. Only POST requests.")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	var person models.Person

	if err = json.Unmarshal(body, &person); err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err = postgres.DB.DB.Create(&person).Error; err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newJsonResponse(w, http.StatusOK, map[string]interface{}{
		"ID":     person.ID,
		"Name":   person.Name,
		"Gender": person.Gender,
		"Birth":  person.Birth,
	})

}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		newErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed. Only PATCH requests.")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	var person models.Person

	personID := r.URL.Query()["id"]

	if err = json.Unmarshal(body, &person); err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	res := postgres.DB.DB.Where("id = ?", personID).Updates(&person)
	if res.Error != nil {
		newErrorResponse(w, http.StatusInternalServerError, res.Error.Error())
		return
	}

	var rPerson models.Person

	if err = postgres.DB.DB.Where("id = ?", personID).First(&rPerson).Error; err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newJsonResponse(w, http.StatusOK, map[string]interface{}{
		"ID":     rPerson.ID,
		"Name":   rPerson.Name,
		"Gender": rPerson.Gender,
		"Birth":  rPerson.Birth,
	})
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		newErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var person models.Person

	personID := r.URL.Query()["id"]

	s := postgres.DB.DB.Delete(&person, personID)
	if s.Error != nil {
		newErrorResponse(w, http.StatusInternalServerError, s.Error.Error())
		return
	}
	newJsonResponse(w, http.StatusOK, map[string]interface{}{"message": "User deleted successfully"})
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		newErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed. Only GET requests.")
		return
	}

	personID := r.URL.Query()["id"]

	var person models.Person

	err := postgres.DB.DB.Where("id = ?", personID).First(&person).Error
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newJsonResponse(w, http.StatusOK, map[string]interface{}{
		"ID":     strconv.Itoa(int(person.ID)),
		"Name":   person.Name,
		"Gender": person.Gender,
		"Birth":  person.Birth,
	})

}
