package handlers

import (
	"encoding/json"
	models2 "github.com/localpurpose/vk-filmoteka/internal/models"
	"github.com/localpurpose/vk-filmoteka/pkg/database/postgres"
	"io"
	"log"
	"net/http"
	"strings"
)

// CreatePerson godoc
//
//	@Security	ApiKeyAuth
//
//	@Summary	Creates person from request body
//	@Tags		persons
//	@Accept		json
//	@Produce	json
//	@Param		person	body		models.Person	true	"Create Person"
//	@Success	200		{object}	models.Person
//	@Router		/person/create [post]
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

	var person models2.Person

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

// UpdatePerson godoc
//
//	@Security	ApiKeyAuth
//
//	@Summary	Updates person from request URL id
//	@Tags		persons
//	@Accept		json
//	@Produce	json
//	@Param		id		path		int				true	"Update Person"
//	@Param		person	body		models.Person	true	"Update Person BODY"
//	@Success	200		{object}	models.Person
//	@Router		/person/update [patch]
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

	var person models2.Person

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

	var rPerson models2.Person

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

// DeletePerson godoc
//
//	@Security	ApiKeyAuth
//
//	@Summary	Deletes person from request URL id
//	@Tags		persons
//	@Accept		json
//	@Produce	json
//	@Param		id	path	int	true	"Delete Person"
//	@Success	200
//	@Router		/person/delete [delete]
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		newErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var person models2.Person

	personID := r.URL.Query()["id"]

	s := postgres.DB.DB.Delete(&person, personID)
	if s.Error != nil {
		newErrorResponse(w, http.StatusInternalServerError, s.Error.Error())
		return
	}
	newJsonResponse(w, http.StatusOK, map[string]interface{}{"message": "User deleted successfully"})
}

// GetPersonByName godoc
//
//	@Security	ApiKeyAuth
//
//	@Summary	Gets person from request URL name
//	@Tags		persons
//	@Accept		json
//	@Produce	json
//	@Param		name	query		string	false	"Get person by name"
//	@Success	200		{object}	models.Person
//	@Router		/person [get]
func GetPersonByName(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		newErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed. Only GET requests.")
		return
	}

	personName := r.URL.Query()["name"]

	var person []models2.Person

	err := postgres.DB.DB.Select("name").Find(&person).Error
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	for i := 0; i < len(person); i++ {
		if strings.Contains(person[i].Name, personName[0]) {
			log.Println("MATCH:", person[i].Name)
		}
	}

	//newJsonResponse(w, http.StatusOK, map[string]interface{}{
	//	"ID":     person.ID,
	//	"Name":   person.Name,
	//	"Gender": person.Gender,
	//	"Birth":  person.Birth,
	//})
}

// GetAllPersons godoc
//
//	@Security	ApiKeyAuth
//
//	@Summary	Gets all persons
//	@Tags		persons
//	@Accept		json
//	@Produce	json
//	@Success	200
//	@Router		/persons [get]
func GetAllPersons(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		newErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed. Only GET requests.")
		return
	}

	type PersonsDB struct {
		ID     uint            `json:"ID"`
		Name   string          `json:"name"`
		Gender string          `json:"gender"`
		Birth  string          `json:"birth"`
		Movies []models2.Movie `json:"movies"`
	}

	var persons []models2.Person
	err := postgres.DB.DB.Find(&persons).Error
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	var actorsRels []models2.Actor
	err = postgres.DB.DB.Find(&actorsRels).Error
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
	}

	var jsonPersons []PersonsDB

	for i := 0; i < len(actorsRels); i++ {
		for j := 0; j < len(persons); j++ {
			if actorsRels[i].PersonID == persons[j].ID {

				var cur_m []models2.Movie
				log.Println(cur_m, persons[j].ID)
				err = postgres.DB.DB.Where("id = ?", actorsRels[i].MovieId).Find(&cur_m).Error
				if err != nil {
					newErrorResponse(w, http.StatusInternalServerError, err.Error())
				}

				jsonPersons = append(jsonPersons, PersonsDB{
					ID:     persons[j].ID,
					Name:   persons[j].Name,
					Gender: "",
					Birth:  "",
					Movies: cur_m,
				})
				break
			}
		}

	}

	for i := 0; i < len(jsonPersons); i++ {
		log.Println(jsonPersons[i])
	}
	log.Println(jsonPersons)

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write([]byte("")); err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

}
