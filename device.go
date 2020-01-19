package main

import (
	// "encoding/json"
	// "fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	// "gopkg.in/mgo.v2/bson"
)

func DeviceList(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("DeviceList")
	
	respondWithJson(w, http.StatusCreated, "device list")
}

func DeviceStatus(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	
	vars := mux.Vars(r)
  deviceId := vars["device_id"]  
	
	log.Println("DeviceList", deviceId)
	
	respondWithJson(w, http.StatusCreated, "DeviceStatus"+deviceId)
}