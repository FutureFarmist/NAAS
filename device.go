package main

import (
	// "encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/stianeikeland/go-rpio"
	"strconv"
	"io/ioutil"
	
	badger "github.com/dgraph-io/badger"
		
	// "gopkg.in/mgo.v2/bson"
)

type Device struct {
		
		Id string `bson:"Id" json:"Id"`
		Name string `bson:"Name,omitempty" json:"Name,omitempty"`
		DeviceId string `bson:"DeviceId,omitempty" json:"DeviceId,omitempty"`
		Desc string `bson:"Desc,omitempty" json:"Desc,omitempty"`

		/* Associated pin */
		Pin	uint8 `bson:"Pin,omitempty" json:"Pin,omitempty"`
		
		GPIO uint8 `bson:"GPIO,omitempty" json:"GPIO,omitempty"`
		
		// ENUM ACTIVE, INACTIVE, ERROR
		Status uint8 `bson:"Status,omitempty" json:"Status,omitempty"`
		
		PinType uint8 `bson:"PinType,omitempty" json:"PinType,omitempty"`
		PinMode uint8 `bson:"PinMode,omitempty" json:"PinMode,omitempty"`

		/* These are for sensor */
		Factor uint16 `bson:"Factor,omitempty" json:"Factor,omitempty"`
		Unit uint16 `bson:"Unit,omitempty" json:"Unit,omitempty"`
		
		/* These are for sensor */
		Value1 uint16 `bson:"Value1,omitempty" json:"Value1,omitempty"`
		Value2 uint16 `bson:"Value2,omitempty" json:"Value2,omitempty"`
		Value3 uint16 `bson:"Value3,omitempty" json:"Value3,omitempty"`
		BoolState bool `bson:"BoolState,omitempty" json:"BoolState,omitempty"`
	
	}

// List Family
func DeviceList(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("DeviceList")
	
	// var valCopy []byte
	err := bgdb.View(func(txn *badger.Txn) error {
	
		item, err := txn.Get(pins_setup_key)
		if err == nil {
			err = item.Value(func(val []byte) error {
				// This func with val would only be called if item.Value encounters no error.
				// log.Println("item: ", string(val))
				
				respondWithJson(w, http.StatusCreated, string(val) )	
				// Copying or parsing val is valid.
				// valCopy = append([]byte{}, val...)
				
				return nil
			})
			if err != nil {
				respondWithJson(w, http.StatusCreated, "{result: 'no value'}")
				log.Println("no value")
			}
		} 
		return nil
	})
	
	if err != nil {
		log.Println("err 3", err)
	}

	// respondWithJson(w, http.StatusCreated, "{ device_list: }")
}


func SetupPins(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("r.Body: ", r.Body)
	
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("converting error")
		respondWithJson(w, http.StatusCreated, "{status: 'fail'}")
	}

	
	err = bgdb.Update(func(txn *badger.Txn) error {
		
		if err := txn.Set(pins_setup_key, body); err != nil {
			respondWithJson(w, http.StatusCreated, "{status: 'setting fail'}")
			return nil
		}
		log.Println("pins setup successful")
		
		item, err := txn.Get(pins_setup_key)
		if err == nil {
			_ = item.Value(func(val []byte) error {
			// This func with val would only be called if item.Value encounters no error.
				
			// Accessing val here is valid.
			log.Println("chaged item: ", string(val))

			// Copying or parsing val is valid.
			// valCopy = append([]byte{}, val...)

			// Assigning val slice to another variable is NOT OK.
			// valNot = val // Do not do this.
			return nil
			})
		} 
	
		return nil
	})
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

func GPIO_on(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	pin := vars["GPIO"]
	
	pinNum, err := strconv.Atoi(pin)
	if err != nil {
		fmt.Println(err, pin)
	}
	pinHdl := rpio.Pin(pinNum)
	pinHdl.Output()
	pinHdl.High()
	
	log.Println("GPIO_on", pin)
	
	respondWithJson(w, http.StatusCreated, "{status: 'ok'}")
}

func GPIO_off(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	
	vars := mux.Vars(r)
	pin := vars["GPIO"]
	
	pinNum, err := strconv.Atoi(pin)
	if err != nil {
		fmt.Println(err, pin)
	}
	
	pinHdl := rpio.Pin(pinNum)
	pinHdl.Output()
	pinHdl.Low()
	
	log.Println("GPIO_off", pin)

	respondWithJson(w, http.StatusCreated, "{status: 'ok'}")
}


