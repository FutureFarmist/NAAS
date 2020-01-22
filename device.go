package main

import (
	// "encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/stianeikeland/go-rpio"
	"strconv"
	// "gopkg.in/mgo.v2/bson"
)

// List Family
func DeviceList(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("DeviceList")

	respondWithJson(w, http.StatusCreated, "device list")
}

func CamList(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("CamListt")

	respondWithJson(w, http.StatusCreated, "CamListt")
}

// Factor Family

func DeviceFactorAirTemp(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("DeviceFactorAirTemp")

	respondWithJson(w, http.StatusCreated, "DeviceFactorAirTemp")
}

// Device Family

func DeviceStatusHdr(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	deviceId := vars["device_id"]

	log.Println("DeviceList", deviceId)

	respondWithJson(w, http.StatusCreated, "DeviceStatus"+deviceId)
}

func PinOn(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	pin := vars["pin"]
	
	pinNum, err := strconv.Atoi(pin)
	if err != nil {
		fmt.Println(err, pin)
	}
	pinHdl := rpio.Pin(pinNum)
	pinHdl.Output()
	pinHdl.High()
	
	log.Println("PinOn", pin)
	
	respondWithJson(w, http.StatusCreated, "{status: 'ok'}")
}

func PinOff(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	
	vars := mux.Vars(r)
	pin := vars["pin"]
	
	pinNum, err := strconv.Atoi(pin)
	if err != nil {
		fmt.Println(err, pin)
	}
	
	pinHdl := rpio.Pin(pinNum)
	pinHdl.Output()
	pinHdl.Low()
	
	log.Println("PinOff", pin)

	respondWithJson(w, http.StatusCreated, "{status: 'ok'}")
}