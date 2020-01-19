package main

import (
	// "fmt"
	"encoding/json"
	"log"
	"net/http"
	"io/ioutil"
	"strings"

	// "os"
	// "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/BurntSushi/toml"
	// "github.com/GeertJohan/go.rice"
)

var config = Config{}
// var mgcli = mgc.MGC{}

// Represents database server and credentials
type Config struct {
	Server   string
	Database string
}

// Read and parse the configuration file
func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}

/* func init() {
	config.Read()
	mgcli.Server = config.Server
	mgcli.Database = config.Database
	mgcli.Connect()
} */

/* Responding */

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	//Allow CORS here By * or specific origin
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithHTML(w http.ResponseWriter, code int, payload []byte) {
	// response, _ := json.Marshal(payload)
	//Allow CORS here By * or specific origin
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(code)
	w.Write(payload)
}

func serveWebApp(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("Serving Web App")
	
	webapp, err := ioutil.ReadFile("web-app/index.html")	
	if err != nil {
		log.Println("Error: Reading Web App %d", err)
	}
	log.Println("%s", webapp)
	respondWithHTML(w, http.StatusCreated, webapp)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
			// Do stuff here
			log.Println(r.RequestURI)
			// Call the next handler, which can be another middleware in the chain, or the final handler.
			next.ServeHTTP(w, r)
	})
}

func main() {
	c := cors.New(cors.Options{
		AllowedMethods:     []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}, //
		AllowedOrigins:     []string{"*"},
		AllowCredentials:   true,
		AllowedHeaders:     []string{"Content-Type", "Bearer", "Bearer ", "content-type", "Origin", "Accept"},
		OptionsPassthrough: false,
	})
	
	/* // return 405 for PUT, PATCH and DELETE
	r.HandleFunc("/users", status(405, "GET", "POST")).Methods("PUT", "PATCH", "DELETE")
	*/
	
	r := mux.NewRouter()
	
	/* Serving Web App Folder */
	// box := rice.MustFindBox("web-app")
	// webappFileServer := http.StripPrefix("/css/", http.FileServer(box.HTTPBox()))
	
	r.HandleFunc("/", serveWebApp).Methods("GET")
	
	// r.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("web-app").HTTPBox()))
	// r.HandleFunc("/", http.FileServer(rice.MustFindBox("web-app").HTTPBox())).Methods("GET")
	
	// r.HandleFunc("/", webappFileServer).Methods("GET")
	
	// serving api
	apiPrefix := "/api/"
	devicePrefix := apiPrefix + "device/"
	// fieldPrefix := apiPrefix + "field/"
	// plantPrefix := apiPrefix + "plant/"
	// camPrefix := apiPrefix + "cam/"
	
	/* Mapper */
	
	/* Device */
	// r.HandleFunc(devicePrefix + "list", DeviceList).Methods("GET")
	// r.HandleFunc(camPrefix + "list", DeviceList).Methods("GET")
	// r.HandleFunc(fieldPrefix + "list", FieldList).Methods("GET")
	// r.HandleFunc(plantPrefix + "list", PlantList).Methods("GET")
	
	// implement DEVICE-ID /status /detail /get-sensor /set-control
	r.HandleFunc(devicePrefix + "{device_id}/status", DeviceStatus).Methods("GET")
	// r.HandleFunc(devicePrefix + "{device_id}/detail", DeviceDetail).Methods("GET")
	// r.HandleFunc(devicePrefix + "{device_id}/get-sensor", DeviceGetSensor).Methods("GET")
	// r.HandleFunc(devicePrefix + "{device_id}/set-control", DeviceSetControl).Methods("POSt")
	
	// // implement by-factor
	// r.HandleFunc(devicePrefix + "by-factor/air-temperature", DeviceFactorAirTemp).Methods("GET")
	// r.HandleFunc(devicePrefix + "by-factor/air-humidity", DeviceFactorAirHumi).Methods("GET")
	// r.HandleFunc(devicePrefix + "by-factor/soil-humidity", DeviceFactorSoilHumi).Methods("GET")
	// r.HandleFunc(devicePrefix + "by-factor/light-intensity", DeviceFactorLightInten).Methods("GET")
	
	// implementing CAMERA-ID + /picture /info /status
	/* r.HandleFunc(apiPrefix + camPrefix + "/camId", DeviceList).Methods("GET")
	r.HandleFunc(apiPrefix + "desk/v1/delete", DeleteDesk).Methods("DELETE") */
	
	/* Automator API */
	/* 
	// implementing FIELD-ID /plant-list /device-list
	r.HandleFunc(fieldPrefix + "{field_id}", FieldData).Methods("GET")
	r.HandleFunc(fieldPrefix + "{field_id}/plant-list", FieldPlantList).Methods("GET")
	r.HandleFunc(fieldPrefix + "{field_id}/device-list", FieldDeviceList).Methods("GET")
	
	// implementing PLANT-ID /device-list
	r.HandleFunc(plantPrefix + "{plant_id}", PlantData).Methods("GET")
	r.HandleFunc(plantPrefix + "{plant_id}/device-list", PlantDeviceList).Methods("GET")
	 */
	if err := http.ListenAndServe(":80", loggingMiddleware(c.Handler(r))); err != nil {
		log.Fatal(err)
	}
	log.Println("Serving NAAS Web Application + API")
	
}