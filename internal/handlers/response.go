package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func newErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	log.Println("[Error]", message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err := w.Write([]byte("{\n\t\"message\":\"" + message + "\"\n}"))
	if err != nil {
		log.Println("error while writing body", err)
		return
	}
}

func newJsonResponse(w http.ResponseWriter, statusCode int, jsn map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	jsonResp, err := json.Marshal(jsn)
	if err != nil {
		log.Println("error while marshalling json")
		return
	}
	if _, err = w.Write(jsonResp); err != nil {
		log.Println("some error while writing response", err)
		return
	}
}
