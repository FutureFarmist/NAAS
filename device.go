package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/stianeikeland/go-rpio"
	"strconv"
	"io/ioutil"
	"errors"
	"os/exec"
	"strings"
	// "reflect"

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
		
		GPIO uint8 `bson:"GPIO" json:"GPIO"`
		
		// ENUM ACTIVE, INACTIVE, ERROR
		Status uint8 `bson:"Status,omitempty" json:"Status,omitempty"`
		
		PinType uint8 `bson:"PinType" json:"PinType"`
		PinMode uint8 `bson:"PinMode" json:"PinMode"`

		/* These are for sensor */
		// DeviceId defines factors
		// Factors []Factor `bson:"Factors,omitempty" json:"Factors,omitempty"`
		Unit uint16 `bson:"Unit,omitempty" json:"Unit,omitempty"`
		
		/* These are for sensor */
		Value1 uint16 `bson:"Value1,omitempty" json:"Value1,omitempty"`
		Value2 uint16 `bson:"Value2,omitempty" json:"Value2,omitempty"`
		Value3 uint16 `bson:"Value3,omitempty" json:"Value3,omitempty"`
		BoolState bool `bson:"BoolState,omitempty" json:"BoolState,omitempty"`
	
	}

// List Family
func Device_List(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("DeviceList")
	
	// var valCopy []byte
	err := bgdb.View(func(txn *badger.Txn) error {
	
		item, err := txn.Get(DEVICE_KEY)
		if err == nil {
			err = item.Value(func(val []byte) error {
				// This func with val would only be called if item.Value encounters no error.
				// log.Println("item: ", string(val))
				
				var devices []Device
				err = json.Unmarshal(val, &devices)
				if err != nil {
					log.Println(err)
				}	
				
				respondWithJson(w, http.StatusCreated, devices)
				// Copying or parsing val is valid.
				// valCopy = append([]byte{}, val...)
				
				return nil
			})
			if err != nil {
				respondWithError(w, http.StatusBadRequest, "{result: 'no value'}")
				log.Println("no value")
			}
			} 
			return nil
		})
		
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "{result: 'no value'}")
		log.Println("err 3", err)
	}
	
	// respondWithError(w, http.StatusBadRequest, "{result: 'no value'}")
}

func GetDeviceInfo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("DeviceInfo")
	device_info, err := get_device_info()
	var devices []DeviceInfo
	err = json.Unmarshal(device_info, &devices)
	if err != nil {
		log.Println(err)
		respondWithError(w, 0, "cannot get device info")
	}
	 log.Println(devices)
	 respondWithJson(w, http.StatusCreated, devices)
}

/* type ResponseSensorValue struct {
	device_id string `bson:"device_id,omitempty" json:"device_id,omitempty"`
	sensor_values []interface `bson:"sensor_values,omitempty" json:"sensor_values,omitempty"`
} */

func (auto *Automator) DeviceValues(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("DeviceValues")
	
	// read device value from sensors and actuator
	for _, device := range auto.devices {
		if len(device.Id) > 0 && device.PinMode == 0 {
			device.read_sensor(auto)
		}
	}
	
	// var outSensorValue []ResponseSensorValue = []ResponseSensorValue{}
	//  for key, sensorVals := range auto.sensor_values {
	//  	log.Println("sensorVals", key, sensorVals.value)
	// // 	outSensorValue= append(outSensorValue, ResponseSensorValue { device_id: key, sensor_values: sensorVals})
	//  }
	// sensors := make([]SensorValue, len(auto.sensor_values))
	// copy(sensors, auto.sensor_values)
	//  for key, sensorVals := range sensors {
	//  	log.Println("sensors", key, sensorVals.value)
	// // 	outSensorValue= append(outSensorValue, ResponseSensorValue { device_id: key, sensor_values: sensorVals})
	//  }
	 
	// outputByte, err := json.Marshal(sensors)
	// if err != nil {
	// 	log.Println(err)
	// 	respondWithError(w, 0, "cannot marshal json")
	// }
	// log.Println(outputByte)
	 
	respondWithJson(w, http.StatusCreated, auto.sensor_values)
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
		if err := txn.Set(DEVICE_KEY, body); err != nil {
			return err
		}
		log.Println("pins setup successful")
		
		// item, err := txn.Get(DEVICE_KEY)
		// if err == nil {
		// 	err = item.Value(func(val []byte) error {
		// 	// This func with val would only be called if item.Value encounters no error.
				
		// 	// Accessing val here is valid.
		// 	// log.Println("chaged item: ", string(val))
			
		// 	// Copying or parsing val is valid.
		// 	// valCopy = append([]byte{}, val...)

		// 	// Assigning val slice to another variable is NOT OK.
		// 	// valNot = val // Do not do this.
		// 	return nil
		// 	})
		// 	return err
		// } else {
		// 	return err
		// }
	
		return nil
	})
	
	if err != nil {
		respondWithJson(w, http.StatusCreated, "{status: 'setting fail'}")
	} else {
		GetDevices()
		automator.cron.Stop()
		automator.setup_cron()
		respondWithJson(w, http.StatusCreated, "{success: true}")
	}
}

func GetDevices() error {
	log.Println("GetDevices")
	err := bgdb.View(func(txn *badger.Txn) error {
		var dvs_json []byte
		// Getting devices into auto.devices
		item, err := txn.Get(DEVICE_KEY)
		if err != nil {
			log.Println("error getting DEVICE_KEY")
			return nil
		}
		
		err = item.Value(func(val []byte) error {
			log.Println("successful getting DEVICE_KEY")
			dvs_json = append([]byte{}, val...)
			return nil
		})
		
		if err != nil {
			log.Println("no value on DEVICE_KEY")
			return nil
		}
		
		// log.Println(string(dvs_json))
		var devices []Device
		err = json.Unmarshal(dvs_json, &devices)
		if err != nil {
			log.Println(err)
			return nil
		}
		
		log.Println("devices: ", devices)
		
		if len(devices) > 0 {
			for _, device := range devices {
				if len(device.Id) > 0 {
					automator.devices[device.Id] = device;
					fmt.Printf("device: %s \n", device.Id)
				}
			}
		}
		
		// log.Println("devices: ", automator.devices)
		return nil
	})
	return err
}

// not implemented yet
func CamList(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("CamListt")
	
	respondWithJson(w, http.StatusCreated, "CamListt")
}

// Factor Family

// not implemented yet
func DeviceFactorAirTemp(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("DeviceFactorAirTemp")
	
	respondWithJson(w, http.StatusCreated, "DeviceFactorAirTemp")
}

// Device Family

// not implemented yet
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
	pin := vars["pin"]
	
	pinNum, err := strconv.Atoi(pin)
	if err != nil {
		log.Println(err, pin)
	}
	gpio, err := GetGPIO(uint8(pinNum))
	if err != nil {
		log.Println(err, pin)
		return	
	}
	log.Println("PinOn: ", gpio)
	pinHdl := rpio.Pin(gpio)
	pinHdl.Output()
	pinHdl.High()
	
	respondWithJson(w, http.StatusCreated, "{status: 'ok'}")
}

func GPIO_off(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	
	vars := mux.Vars(r)
	pin := vars["pin"]
	
	pinNum, err := strconv.Atoi(pin)
	if err != nil {
		log.Println(err, pin)
	}
	gpio, err := GetGPIO(uint8(pinNum))
	if err != nil {
		log.Println(err, pin)
		return	
	}
	log.Println("PinOff: ", gpio)
	pinHdl := rpio.Pin(gpio)
	pinHdl.Output()
	pinHdl.Low()
	

	respondWithJson(w, http.StatusCreated, "{status: 'ok'}")
}

func GetGPIO(pin uint8) (uint8, error) {
	for _, device := range automator.devices {
		if device.Pin == pin {
			return device.GPIO, nil
		}
	}
	return 0, errors.New("no gpio")
}

func (device *Device) read_sensor(auto *Automator) error {
	log.Println("Read sensor")
	if len(device.DeviceId) > 0 {
		log.Println("PIN", device.Pin, ", GPIO ", device.GPIO)
		// FT_SOIL_HUMIDITY
		// FT_SOIL_TEMPERATURE
		// FT_AIR_HUMIDITY
		// FT_AIR_TEMPERATURE
		if device.GPIO >= 0 && (device.DeviceId == "dht11" || device.DeviceId == "dht2" || device.DeviceId == "am2302" ){
			log.Println("dht")
			humi, temp := "", ""
			var err error
			if device.DeviceId == "dht11" {
				humi, temp, err = read_from_dht(strconv.Itoa(int(device.GPIO)), "11")
			} else if device.DeviceId == "dht22" {
				humi, temp, err = read_from_dht(strconv.Itoa(int(device.GPIO)),"22")
			} else if device.DeviceId == "am2302" {
				humi, temp, err = read_from_dht(strconv.Itoa(int(device.GPIO)), "2302")
			}
				
			log.Println("reading dht "+ fmt.Sprintf("%s", humi) + fmt.Sprintf("%s", temp))
			if err == nil {
				// push value to sensor_values
				auto.update_sensor_values(device.Id, FT_AIR_HUMIDITY, humi, false)
				auto.update_sensor_values(device.Id, FT_AIR_TEMPERATURE, temp, false)
			}
			
		} else if device.GPIO >= 0 && device.DeviceId == "sms-lm393" {
			log.Println("Reading mosture sensor")
			moisture_high, err := read_sms_lm393(device.GPIO)
			if err != nil {
				log.Println("reading mosture sensor fail ", err)
				return nil
			}
			auto.update_sensor_values(device.Id, FT_SOIL_HUMIDITY, moisture_high, true)	
		}
		return nil
	}
	return errors.New("fail reading device")
}

func (auto *Automator) update_sensor_values(device_id string, ifactor Factor, ivalue string, is_bool bool) {
	// var update bool = true
	// var f32v float32
	// var bv bool
	log.Println("update_sensor_values")
	updated := false
	is_bool_text := "false"
	if is_bool {
		is_bool_text = "true"
	}
	for key, sv := range auto.sensor_values {
		if sv.Device_id == device_id && Factor(sv.Factor) == ifactor {
			auto.sensor_values[key].Value = ivalue
			auto.sensor_values[key].Is_boolean = is_bool_text
			// sv.Is_boolean = "true"
			log.Println("update exist value: ", ivalue, " (Boolean="+is_bool_text+")")
			updated = true
		} 
	}
	if !updated {
		log.Println("update insert")
		
		auto.sensor_values = append(auto.sensor_values, SensorValue{ Device_id: device_id, Factor: uint8(ifactor), Value: ivalue, Is_boolean: is_bool_text})
	}
	/* if sensorVal.factor == factor && reflect.TypeOf(sensorVal.value) != nil {
		f32v, err := strconv.ParseFloat(ivalue, 32)
		if err != nil {
			log.Println("error parsing float at update_sensor_values")
		}
		
	} else if sensorVal.factor == factor && reflect.TypeOf(sensorVal.boolean) != nil {
		bv, err := strconv.ParseBool(ivalue)
		if err != nil {
			log.Println("error parsing float at update_sensor_values")
		}
		
	} else {
		update = false
	} */
	
	// if update {
		
		
	// }
}

func read_from_dht(gpio string, dht string) (string, string, error){
	
	// expect true,temp,humi -> true,35.3,49.5
	cmd, err := exec.Command("python", "dht.py", dht, gpio).Output()
	fmt.Println("Running dht.py ", dht, gpio)
	// Wait for the Python program to exit.
	// err = cmd.Run()
	if err != nil {
		// cmd = string(cmd)
		fmt.Println("err:", err)
		return "","",errors.New("fail reading DHT11")
	}
	fmt.Println("DHT11: ", string(cmd))
	
	th := strings.Split(strings.TrimSuffix(string(cmd), "\n"), ",")
	// in c
	temp := th[1]
	// in %
	humi := th[2]
	
	return humi, temp, nil
}

func read_sms_lm393(gpio uint8) (string, error) {
	log.Println("read_sms_lm393")
	// pinNum, err := strconv.Atoi(pin)
	// if err == nil {
		// gpio, err := GetGPIO(uint8(pinNum))
		// if err == nil {
	log.Println("read_sms_lm393: ", gpio)
	pinHdl := rpio.Pin(gpio)
	pinHdl.Input()
	res := pinHdl.Read() 
	log.Println("res: ", res)
	if uint8(res) == 0 {
		return "false", nil
	} else if uint8(res) == 1 {
		return "true", nil
	}
		// }
	// }
	// log.Println(err, pin)
	return "", errors.New("cannot read from sms_lm393")
}

// utility func
func clear_ctls() {
	
	bgdb.Update(func(txn *badger.Txn) error {
		
		if err := txn.Set(CTLS_KEY, []byte("")); err != nil {
			log.Println("clear_ctls fail")
			log.Println(err)
			// respondWithError(w, http.StatusBadRequest, "clear_ctls fail")
			return err
		} else {
			log.Println("clear_ctls successful")
			// respondWithJson(w, http.StatusCreated, "clear_ctls")
			return nil
		}	
			return nil
	})
}

// utility func
func clear_devices() {
	
	bgdb.Update(func(txn *badger.Txn) error {
		
		if err := txn.Set(DEVICE_KEY, []byte("")); err != nil {
			log.Println("clear_devices fail")
			log.Println(err)
			// respondWithError(w, http.StatusBadRequest, "clear_ctls fail")
			return err
		} else {
			log.Println("clear_devices successful")
			// respondWithJson(w, http.StatusCreated, "clear_ctls")
			return nil
		}	
			return nil
	})
}
