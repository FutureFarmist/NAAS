package main

import (
	// "fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"errors"
	// "os/exec"

	// "math"
	// "sync"
	// "os"
	// "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/BurntSushi/toml"

	"github.com/stianeikeland/go-rpio"
	
	badger "github.com/dgraph-io/badger"
	// "github.com/dgraph-io/badger/v2/options"
	cron "github.com/robfig/cron/v3"
	// "github.com/GeertJohan/go.rice"
	
	// "github.com/kidoman/embd"
	// _ "github.com/kidoman/embd/host/all"
	
	// "gobot.io/x/gobot"
	// "gobot.io/x/gobot/drivers/spi"
	
	// "gobot.io/x/gobot/drivers/gpio"
	// "gobot.io/x/gobot/platforms/firmata"
	
	// "github.com/d2r2/go-dht"
	// "github.com/MichaelS11/go-dht"
	
	// "periph.io/x/periph/host"
	// "periph.io/x/periph/conn/gpio"
	// "periph.io/x/periph/conn/gpio/gpioreg"
	// "periph.io/x/periph/host/rpi"
)

var (
	config = Config{}
	bgdb *badger.DB
	automator *Automator

	DEVICE_KEY = []byte("pins_setup")
	CTLS_KEY = []byte("ctls")
	
	// CTL_KEY = []byte("ctl")
	// ctl1, ctl2, ctl3, ...
	// CTLS_COUNT_KEY = []byte("ctls_count")
	// ctls_update_DATE
	// CTLS_UPDATE_KEY = []byte("ctls_update")

	// every update controllers would give it controller 
	// ctls_ID
	// active_controllers_id_key = []byte("active_controllers_id")

)
var secondParser = cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
type Automator struct {
	cron *cron.Cron
	ctls map[uint16]Controller
	devices map[string]Device
	ctl_entry_map map[uint16]cron.EntryID

	// sensor_values[DEVICE_ID][FACTOR]string
	sensor_values []SensorValue

	device_info map[string]DeviceInfo
}

func (auto *Automator) setup_cron() error {
	
	println("setup_cron")
	auto.cron = cron.New(cron.WithParser(secondParser), cron.WithChain())
	if auto.cron != nil {
		return nil	
	}
	auto.cron.Start()
	defer auto.cron.Stop()

	println("run run_controllers everyday midnight")
	auto.cron.AddFunc("0 0 0 * * *", func() { auto.run_controllers()})
	
	auto.run_controllers()
	
	auto.cron = cron.New(cron.WithSeconds())
	if auto.cron != nil {
		return nil	
	}
	return errors.New("new cron fail")
}

type DeviceInfo struct {
	DeviceId 	string 	`bson:"DeviceId,omitempty" json:"DeviceId,omitempty"`
	Name 			string 	`bson:"Name,omitempty" json:"Name,omitempty"`
	PinMode 	uint8 	`bson:"PinMode,omitempty" json:"PinMode,omitempty"`
	PinType 	uint8 	`bson:"PinType,omitempty" json:"PinType,omitempty"`
	ControlScheme uint8 `bson:"ControlScheme,omitempty" json:"ControlScheme,omitempty"`
	Factors 	string `bson:"Factors,omitempty" json:"Factors,omitempty"`
}

type SensorValue struct {
	Device_id string `bson:"Device_id,omitempty" json:"Device_id,omitempty"`
	Factor	uint8 `bson:"Factor,omitempty" json:"Factor,omitempty"`
	Value	string `bson:"Value,omitempty" json:"Value,omitempty"`
	Is_boolean string `bson:"Is_boolean,omitempty" json:"Is_boolean,omitempty"`
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

/* Responding */

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Println("marshalling respondWithJson Error")
		log.Println(err)
	}
	log.Println(payload)
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

func get_device_info() ([]byte, error) {

	// get device_info from device-list.json
	device_info_json, err := ioutil.ReadFile("device_info.json")
	if err != nil {
		return nil, errors.New("error reading device_info.json")
	}
	return device_info_json, nil
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
	if err := rpio.Open(); err != nil {
		panic(err)
	}
	defer rpio.Close()

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

	// TODO: become utility functions on NAAS and Seed
	// on-demand commands
	// clear_clts()
	// clear_devices()

	err = bgdb.Update(func(txn *badger.Txn) error {
	
		// err := txn.Set([]byte("test"), []byte("test"))
		// if err != nil {
		// 	log.Println("err 1 ", err)
		// 	return err
		// }
		// log.Println("set")
		item, err := txn.Get(DEVICE_KEY)
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
		// cron: cron.New(cron.WithParser(secondParser), cron.WithChain()),
		cron: cron.New(cron.WithSeconds()),
		ctls: make(map[uint16]Controller), 
		devices: make(map[string]Device), 
		ctl_entry_map: make(map[uint16]cron.EntryID), 
		sensor_values: []SensorValue{},
		device_info: make(map[string]DeviceInfo),
	}

	automator = &auto
	device_info_json, err := get_device_info()
	var device_info []DeviceInfo = []DeviceInfo{}
	err = json.Unmarshal(device_info_json, &device_info)
	if err != nil {
		log.Println(err)
	}

	for _, di := range device_info {
		auto.device_info[di.DeviceId] = di
	}
	
	// cron := cron.New(WithParser(secondParser), WithChain())
	println("new cron")
	
	auto.setup_cron()
	
	// mocking device_value
	auto.sensor_values = append([]SensorValue{}, SensorValue{ Device_id: "7", Factor: 3, Value: "59.3", Is_boolean: "false"})
	auto.sensor_values = append(auto.sensor_values, SensorValue{ Device_id: "7", Factor: 4, Value: "37.7", Is_boolean: "false"})
	auto.sensor_values = append(auto.sensor_values, SensorValue{ Device_id: "40", Factor: 1, Value: "78.9", Is_boolean: "false"})
	/* auto.sensor_values["dv2"] = append(
		auto.sensor_values["dv2"], 
		SensorValue{ factor: 2, value: 5.3}) */
	
	r := mux.NewRouter()
	
	// serving api
	apiPrefix := "/api/"
	devicePrefix := apiPrefix + "device/"
	controllerPrefix := apiPrefix + "controller/"
	// fieldPrefix_ := apiPrefix + "field/"
	// plantPrefix := apiPrefix + "plant/"
	// camPrefix := apiPrefix + "cam/"

	r.HandleFunc(apiPrefix+"setup-pins", SetupPins).Methods("POST")
	
	/* Mapper API */

	/* Device */
	r.HandleFunc(devicePrefix+"list", Device_List).Methods("GET", "POST")
	// r.HandleFunc(camPrefix+"list", CamList).Methods("GET", "POST")
	// r.HandleFunc(fieldPrefix+"list", FieldList).Methods("GET", "POST")
	// r.HandleFunc(plantPrefix+"list", PlantList).Methods("GET", "POST")
	r.HandleFunc(devicePrefix+"info", GetDeviceInfo).Methods("GET", "POST")
	r.HandleFunc(devicePrefix+"values", auto.DeviceValues).Methods("GET", "POST")

	r.HandleFunc(devicePrefix+"on/{pin}", GPIO_on).Methods("GET", "POST")
	r.HandleFunc(devicePrefix+"off/{pin}", GPIO_off).Methods("GET", "POST")
	
	r.HandleFunc(controllerPrefix+"list", ControllerList).Methods("GET", "POST")
	r.HandleFunc(controllerPrefix+"update", auto.UpdateControllers).Methods("POST")

	// implement DEVICE-ID /status /detail /get-sensor /set-control
	r.HandleFunc(devicePrefix+"{device_id}/status", DeviceStatusHdr).Methods("GET", "POST")
	// r.HandleFunc(devicePrefix + "{device_id}/detail", DeviceDetail).Methods("POST")
	// r.HandleFunc(devicePrefix + "{device_id}/get-sensor", DeviceGetSensor).Methods("POST")
	// r.HandleFunc(devicePrefix + "{device_id}/set-control", DeviceSetControl).Methods("POSt")

	// // implement by-factor
	// r.HandleFunc(devicePrefix+"by-factor/air-temperature", DeviceFactorAirTemp).Methods("GET", "POST")
	// r.HandleFunc(devicePrefix + "by-factor/air-humidity", DeviceFactorAirHumi).Methods("POST")
	// r.HandleFunc(devicePrefix + "by-factor/soil-humidity", DeviceFactorSoilHumi).Methods("POST")
	// r.HandleFunc(devicePrefix + "by-factor/light-intensity", DeviceFactorLightInten).Methods("POST")

	// implementing CAMERA-ID + /picture /info /status
	// r.HandleFunc(apiPrefix+camPrefix+"/camId", CamList).Methods("GET", "POST")
	// r.HandleFunc(apiPrefix + "desk/v1/delete", DeleteDesk).Methods("POST")


	// implementing FIELD-ID /plant-list /device-list
	// r.HandleFunc(fieldPrefix+"{field_id}", FieldData).Methods("GET", "POST")
	// r.HandleFunc(fieldPrefix + "{field_id}/plant-list", FieldPlantList).Methods("POST")
	// r.HandleFunc(fieldPrefix + "{field_id}/device-list", FieldDeviceList).Methods("POST")

	// implementing PLANT-ID /device-list
	// r.HandleFunc(plantPrefix+"{plant_id}", PlantData).Methods("GET", "POST")
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
	go http.ListenAndServe(":8080", cwa.Handler(webapp))
	
	log.Println("Serving NAAS Web Application + API")
	

	/* temperature, humidity, retried, err :=
		dht.ReadDHTxxWithRetry(dht.DHT11, 37, false, 10)
	if err != nil {
		log.Fatal(err)
	}
	// Print temperature and humidity
	fmt.Printf("Temperature = %v*C, Humidity = %v%% (retried %d times)\n",
		temperature, humidity, retried) */
		
		
	/* err = dht.HostInit()
	if err != nil {
		fmt.Println("HostInit error:", err)
	}

	dht, err := dht.NewDHT(rpi.P1_7, dht.Fahrenheit, "dht11")
	if err != nil {
		fmt.Println("NewDHT error:", err)
	}

	humidity, temperature, err := dht.ReadRetry(11)
	if err != nil {
		fmt.Println("Read error:", err)
	}

	fmt.Printf("humidity: %v\n", humidity)
	fmt.Printf("temperature: %v\n", temperature) */
	
	
		
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
	