package main

/* import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"
) */

/* func CreateFile(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var file models.IFile
	if err := json.NewDecoder(r.Body).Decode(&file); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	log.Println("CreateProject: ")
	file.Id = bson.NewObjectId()
	// file.AppId = bson.ObjectIdHex(file.AppId.(string))
	if file.FileIndexId != "" {
		file.FileIndexId = bson.NewObjectId()
	}
	if err := mgcli.FileInsert(file); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, file)
} */