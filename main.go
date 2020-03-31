package main

import (
	// "fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	// "math"
	// "sync"
	// "os"
	// "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/BurntSushi/toml"
	_ "github.com/stianeikeland/go-rpio"
	
	badger "github.com/dgraph-io/badger"
	// "github.com/dgraph-io/badger/v2/options"
	cron "github.com/robfig/cron/v3"	
	// "github.com/GeertJohan/go.rice"
	
	/* "gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata" */
)

var (
	config = Config{}
	bgdb *badger.DB

	pins_setup_key = []byte("pins_setup")
	
)
var secondParser = cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
type Automator struct {
	cron *cron.Cron
	ctls map[uint16]Controller
	devices map[string]Device
}

// Represents database server and credentials
type Config struct {
	Server   string
	Database string
}

type FileSystem struct {
	fs http.FileSystem
}

// Open opens file
func (fs FileSystem) Open(path string) (http.File, error) {
	f, err := fs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := strings.TrimSuffix(path, "/") + "/index.html"
		if _, err := fs.fs.Open(index); err != nil {
			return nil, err
		}
	}

	return f, nil
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
	infinite_wait := make(chan string)
	
	/* if err := rpio.Open(); err != nil {
		panic(err)
	}
	defer rpio.Close() */

	// dbOpts := badger.DefaultOptions("naas-db").
	// 	WithSyncWrites(false).
	// 	WithNumVersionsToKeep(math.MaxInt64).
	// 	// WithLogger(&x.ToGlog{}).
	// 	WithCompression(options.None).
	// 	WithEventLogging(false).
	// 	WithLogRotatesToFlush(10).
	// 	WithMaxCacheSize(50) // TODO(Aman): Disable cache altogether
		
	// db, err := badger.Open(dbOpts)
	db, err := badger.Open(badger.DefaultOptions("naas-db"))
  if err != nil {
	  log.Fatal(err)
  }
	defer bgdb.Close()
	
	bgdb = db
	
	err = bgdb.Update(func(txn *badger.Txn) error {
	
		// err := txn.Set([]byte("test"), []byte("test"))
		// if err != nil {
		// 	log.Println("err 1 ", err)
		// 	return err
		// }
		// log.Println("set")
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
	
	if err != nil {
		log.Println("err 3", err)
	}
	
	auto := Automator { 
		cron: cron.New(cron.WithParser(secondParser), cron.WithChain()),
		ctls: make(map[uint16]Controller), 
		devices: make(map[string]Device), 
	}
	// cron := cron.New(WithParser(secondParser), WithChain())
	println("new cron")
	
	auto.cron.Start()
	defer auto.cron.Stop()
	
	println("run run_controllers everyday midnight")
	auto.cron.AddFunc("0 0 0 * * *", func() { auto.run_controllers()})
	
	auto.run_controllers()
	
	
	// clear_ctls_key()
	r := mux.NewRouter()
	
	// serving api
	apiPrefix := "/api/"
	devicePrefix := apiPrefix + "device/"
	controllerPrefix := apiPrefix + "controller/"
	fieldPrefix := apiPrefix + "field/"
	plantPrefix := apiPrefix + "plant/"
	camPrefix := apiPrefix + "cam/"

	r.HandleFunc(apiPrefix+"setup-pins", SetupPins).Methods("GET", "POST")
	
	/* Mapper API */

	/* Device */
	r.HandleFunc(devicePrefix+"list", DeviceList).Methods("GET", "POST")
	r.HandleFunc(camPrefix+"list", CamList).Methods("GET", "POST")
	r.HandleFunc(fieldPrefix+"list", FieldList).Methods("GET", "POST")
	r.HandleFunc(plantPrefix+"list", PlantList).Methods("GET", "POST")

	r.HandleFunc(devicePrefix+"on/{pin}", GPIO_on).Methods("GET", "POST")
	r.HandleFunc(devicePrefix+"off/{pin}", GPIO_off).Methods("GET", "POST")
	
	r.HandleFunc(controllerPrefix+"list", ControllerList).Methods("GET", "POST")
	r.HandleFunc(controllerPrefix+"update", UpdateControllers).Methods("GET", "POST")

	// implement DEVICE-ID /status /detail /get-sensor /set-control
	r.HandleFunc(devicePrefix+"{device_id}/status", DeviceStatusHdr).Methods("GET", "POST")
	// r.HandleFunc(devicePrefix + "{device_id}/detail", DeviceDetail).Methods("POST")
	// r.HandleFunc(devicePrefix + "{device_id}/get-sensor", DeviceGetSensor).Methods("POST")
	// r.HandleFunc(devicePrefix + "{device_id}/set-control", DeviceSetControl).Methods("POSt")

	// // implement by-factor
	r.HandleFunc(devicePrefix+"by-factor/air-temperature", DeviceFactorAirTemp).Methods("GET", "POST")
	// r.HandleFunc(devicePrefix + "by-factor/air-humidity", DeviceFactorAirHumi).Methods("POST")
	// r.HandleFunc(devicePrefix + "by-factor/soil-humidity", DeviceFactorSoilHumi).Methods("POST")
	// r.HandleFunc(devicePrefix + "by-factor/light-intensity", DeviceFactorLightInten).Methods("POST")

	// implementing CAMERA-ID + /picture /info /status
	r.HandleFunc(apiPrefix+camPrefix+"/camId", CamList).Methods("GET", "POST")
	// r.HandleFunc(apiPrefix + "desk/v1/delete", DeleteDesk).Methods("POST")

	/* Automator API */

	// implementing FIELD-ID /plant-list /device-list
	r.HandleFunc(fieldPrefix+"{field_id}", FieldData).Methods("GET", "POST")
	// r.HandleFunc(fieldPrefix + "{field_id}/plant-list", FieldPlantList).Methods("POST")
	// r.HandleFunc(fieldPrefix + "{field_id}/device-list", FieldDeviceList).Methods("POST")

	// implementing PLANT-ID /device-list
	r.HandleFunc(plantPrefix+"{plant_id}", PlantData).Methods("GET", "POST")
	// r.HandleFunc(plantPrefix + "{plant_id}/device-list", PlantDeviceList).Methods("POST")
	
	
	c := cors.New(cors.Options{
		AllowedMethods:     []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}, //
		AllowedOrigins:     []string{"*"},
		AllowCredentials:   true,
		AllowedHeaders:     []string{"Content-Type", "Bearer", "Bearer ", "content-type", "Origin", "Accept"},
		OptionsPassthrough: false,
	})	
	go http.ListenAndServe(":3030", loggingMiddleware(c.Handler(r)))
	
	// box := rice.MustFindBox("web-app")
	// webappFileServer := http.StripPrefix("/", http.FileServer(box.HTTPBox()))
	// r.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("web-app").HTTPBox()))
	// fs := http.FileServer(http.Dir("web-app"))
	// webapp := http.NewServeMux()
	webapp := mux.NewRouter()
	// webapp.Handle("/", fs)
	webapp.PathPrefix("/").Handler(http.FileServer(http.Dir("./web-app/")))
	
	cwa := cors.New(cors.Options{
		AllowedMethods:     []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}, //
		AllowedOrigins:     []string{"*"},
		AllowCredentials:   true,
		AllowedHeaders:     []string{"Content-Type", "Bearer", "Bearer ", "content-type", "Origin", "Accept"},
		OptionsPassthrough: false,
	})	
	go http.ListenAndServe(":80", cwa.Handler(webapp))
	
	log.Println("Serving NAAS Web Application + API")
	
	<-infinite_wait
	
}

/* // return 405 for PUT, PATCH and DELETE
r.HandleFunc("/users", status(405, "GET", "POST")).Methods("PUT", "PATCH", "DELETE")
*/
/* Serving Web App Folder */



	// pi := raspi.NewAdaptor()
	
	// hardcoded
	/* r1 := gpio.NewLedDriver(r, "14")
	r2 := gpio.NewLedDriver(r, "15")
	r3 := gpio.NewLedDriver(r, "16")
	r4 := gpio.NewLedDriver(r, "18") */

	/* devices := []Device{
		{
			DeviceId:	1,
			Name: "Relay 1",
			Pin: 14
		},
		{
			DeviceId:	2,
			Name: "Relay 2",
			Pin: 15
		},
		{
			DeviceId:	3,
			Name: "Relay 3",
			Pin: 16
		},
		{
			DeviceId:	4,
			Name: "Relay 4",
			Pin: 18
		},
	} */
	
	/* sv1 := gpio.NewLedDriver(pi, devices[1].Pin)
	sv2 := gpio.NewLedDriver(pi, devices[2].Pin)
	sv3 := gpio.NewLedDriver(pi, devices[3].Pin)
	sv4 := gpio.NewLedDriver(pi, devices[4].Pin) */
	
	/*sv1 := rpio.Pin(devices[1].Pin)
	sv2 := rpio.Pin(devices[2].Pin)
	sv3 := rpio.Pin(devices[3].Pin)
	sv4 := rpio.Pin(devices[4].Pin)
	
	 work := func() {
		gobot.Every(1*time.Second, func() {
						led.Toggle()
		})
	}

	robot := gobot.NewRobot("raspi",
					[]gobot.Connection{pi},
					[]gobot.Device{sv1, sv2, sv3, sv4},
					work,
	)

	robot.Start() */
	