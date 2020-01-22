package main

import (
	// "encoding/json"
	// "fmt"
	"log"
	"net/http"
	// "gopkg.in/mgo.v2/bson"
)

func PlantList(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("PlantList")

	respondWithJson(w, http.StatusCreated, "PlantList")
}

func PlantData(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("PlantData")

	respondWithJson(w, http.StatusCreated, "PlantData")
}
