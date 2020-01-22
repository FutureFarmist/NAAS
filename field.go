package main

import (
	// "encoding/json"
	// "fmt"
	"log"
	"net/http"
	// "gopkg.in/mgo.v2/bson"
)

func FieldList(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("FieldList")

	respondWithJson(w, http.StatusCreated, "FieldList")
}

func FieldData(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("FieldData")

	respondWithJson(w, http.StatusCreated, "FieldData")
}
