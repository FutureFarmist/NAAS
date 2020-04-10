package main

import (
	_ "bytes"
	_ "encoding/binary"
	"encoding/json"
	"fmt"
	_ "github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/stianeikeland/go-rpio"
	"strconv"
	"io/ioutil"
	badger "github.com/dgraph-io/badger"
	// _ IteratorOptions "github.com/dgraph-io/badger/v2/options"
	// cron "github.com/robfig/cron/v3"	
	"time"
	"reflect"
)

type (
	
	// The algorithm is 
	// 1 check for active controller that would work for that day 
	// 2 run checking every 10 second for every active controller according to time/sensor 
	Controller struct {
		Id uint16 `bson:"Id,omitempty" json:"Id,omitempty"`	
		Name string `bson:"INamed,omitempty" json:"Name,omitempty"`	
		Desc string `Desc:"Desc,omitempty" json:"Desc,omitempty"`	

		// Active controller or not. For manual control
		Active bool `bson:"Active,omitempty" json:"Active,omitempty"`	
		
		// sensor pin id
		Sensor string `bson:"Sensor,omitempty" json:"Sensor,omitempty"`	
		
		Factor Factor  `bson:"Factor,omitempty" json:"Factor,omitempty"`	
		
		// DeviceLinks []DeviceLink `bson:"DeviceLinks,omitempty" json:"DeviceLinks,omitempty"`
		
		/* ENUM
		TIME_POLICY - control based on time
		MEASUREMENT_POLICY - control based on value, this includes VALUE_CONTROL and BOOLEAN_CONTROL scheme
		TIME_MEASUREMENT_POLICY - control value for period(s) of time
		*/
		Policy Policy `bson:"Policy,omitempty" json:"Policy,omitempty"`	
		
		// ENUM: BOOLEAN_CONTROL, VALUE_CONTROL
		ControlScheme ControlScheme `bson:"ControlScheme,omitempty" json:"ControlScheme,omitempty"`
		
		// cron code values
		Cron Cron `bson:"Cron,omitempty" json:"Cron,omitempty"`	
		
		// optional Season Start Date - End Date
		// Date DD:MM:YYYY in YYYY in CE (common era)
		// Control for these period of time
		// If not specified either of them, this controller always run following ranges 
		// TODO add real Date
		// SessionStartDate string	`bson:"SessionStartDate,omitempty" json:"SessionStartDate,omitempty"`
		// SeasonEndDate		string	`bson:"SeasonEndDate,omitempty" json:"SeasonEndDate,omitempty"`
		
		/* set to active everyday */
		// ActiveDaily bool `bson:"ActiveDaily,omitempty" json:"ActiveDaily,omitempty"`	
		
		// // set week days to active in 7 days format start with Sunday
		// // 1 for Sunday, 2 for Monday and so on
		// ActiveWeekDayRanges []uint8 `bson:"ActiveWeekDayRanges,omitempty" json:"ActiveWeekDayRanges,omitempty"`	
		
		// // set month days to active in 31 
		// // 1 for the first day of the month and so on
		// ActiveMonthDayRanges []uint8 `bson:"ActiveMonthDayRanges,omitempty" json:"ActiveMonthDayRanges,omitempty"`	
		
		// // set months to active
		// // 1 for January and so on
		// ActiveMonthRanges []uint8  `bson:"ActiveMonthRanges,omitempty" json:"ActiveMonthRanges,omitempty"`	
		
		// TimePeriods []TimePeriod `bson:"TimePeriods,omitempty" json:"TimePeriods,omitempty"`
		
		// Daytime HH:MM:SS
		// DaytimeStart Datetime `bson:"DaytimeStart,omitempty" json:"DaytimeStart,omitempty"`
		// DaytimeEnd Datetime `bson:"DaytimeEnd,omitempty" json:"DaytimeEnd,omitempty"`
		
		
		// value manipulation capability of device
		// interface for a device are increase(), decrease()
		// ENUM INCREASE, DECREASE, BOTH
		// ValueManipulation uint8 `bson:"ValueManipulation,omitempty" json:"ValueManipulation,omitempty"` 
		
		// VALUE_CONTROL policy
		
		// Range value optimization
		// NAAS allows the value to fructuate between PreferredMin and PreferredMax
		// When it could only increase value, PreferredMin is evaluated if the value is below.
		// When it could only decrease value, PreferredMax is evaluated if the value is above.
		// If device can manipulate in both directions, both PreferredMin and PreferredMax are evaluated for control.
		// When, OptimalVal is set, controlling device would leave manipulating the value
		// optimization is stop at this value
		OptimalVal	float32	`bson:"OptimalVal,omitempty" json:"OptimalVal,omitempty"`	
		PreferredMin	float32 `bson:"PreferredMin,omitempty" json:"PreferredMin,omitempty"`	
		PreferredMax	float32 `bson:"PreferredMax,omitempty" json:"PreferredMax,omitempty"`	
		
		// for value control 
		IncreasingDevices []string `bson:"IncreasingDevices,omitempty" json:"IncreasingDevices,omitempty"`	
		DecreasingDevices []string `bson:"DecreasingDevices,omitempty" json:"DecreasingDevices,omitempty"`	
		
		// BOOLEAN_CONTROL policy
		
		// true/false
		// This setting constrained by time period, not by value
		// PreferredState	bool
		
		// acting on device when value is true
		BoolTrueDevices []string `bson:"BoolTrueDevices,omitempty" json:"BoolTrueDevices,omitempty"`	
		
		// acting on device when value is false
		BoolFalseDevices []string `bson:"BoolFalseDevices,omitempty" json:"BoolFalseDevices,omitempty"`	
		
		
		/* ENUM
		TimeType
		Period
		Daily
		Weekly
		Monthly
		Ondemand
		*/
		// TimeType Date `bson:"TimeType,omitempty" json:"TimeType,omitempty"`
		
		ControlStatus ControlStatus `bson:"ControlStatus,omitempty" json:"ControlStatus,omitempty"`	
		
	}
	
	DeviceLink struct {
		// id of a sensor device/pin
		SensorDevices string `bson:"SensorDevices,omitempty" json:"SensorDevices,omitempty"`
		
		// ids of actuator devices/pins
		ActuatorDevices []string `bson:"ActuatorDevices,omitempty" json:"ActuatorDevices,omitempty"`

	}
	
	TimePeriod struct {
		// starting time
		// TODO 
		Start string `bson:"Start,omitempty" json:"Start,omitempty"` 
		
		// end of time range. this doesn't need to be specified
		// TODO  
		End string  `bson:"End,omitempty" json:"End,omitempty"`
		
		// in case of End doesn't specified, this must be specified to identify ending of period from starting period in Hour range
		HourFromStart uint8 `bson:"HourFromStart,omitempty" json:"HourFromStart,omitempty"`
		
	}
	
	Cron struct {
		Second string `bson:"Second,omitempty" json:"Second,omitempty"`	
		Minute string `bson:"Minute,omitempty" json:"Minute,omitempty"`	
		Hour string `bson:"Hour,omitempty" json:"Hour,omitempty"`	
		Dom string `bson:"Dom,omitempty" json:"Dom,omitempty"`	
		Month string `bson:"Month,omitempty" json:"Month,omitempty"`	
		Dow string `bson:"Dow,omitempty" json:"Dow,omitempty"`	
		Year string `bson:"Year,omitempty" json:"Year,omitempty"`	
	}

	
	DeviceStatus uint16
	DeviceKind uint16
	ControlScheme uint16
	Policy uint8
	ControlStatus uint8
	Factor uint8
	
)

const (
	DEVICE_INACTIVE DeviceStatus = iota // 0
	DEVICE_ACTIVE 
	DEVICE_ERROR	
)

const (
	TIME_POLICY Policy = iota  // 0
	MEASUREMENT_POLICY
	TIME_MEASUREMENT_POLICY
)

const (
	VALUE_CONTROL ControlScheme = iota  // 0
	BOOLEAN_CONTROL
)

// Control Status
const (
	CS_ONHOLD ControlStatus = iota  // 0
	CS_BOOL_TRUE_DEVICES_ACTIVE
	CS_BOOL_FALSE_DEVICES_ACTIVE
	CS_INCREASING_DEVICES_ACTIVE
	CS_DECREASING_DEVICES_ACTIVE
)

// Controlling Factor FT 
const (
	FT_SOIL_HUMIDITY Factor = iota + 1 // 1, 2, 3, ...
	FT_SOIL_TEMPERATURE
	FT_AIR_HUMIDITY
	FT_AIR_TEMPERATURE
)


// not implemented yet, replaced by run_controllers
// run this every day at 00:00:00
// this function help run cron of everyday
// first it clear active controller then it go through all controller
// to find which on is to run today
func (c *Controller) startScheduler() bool {
	
	return true
}

// not implemented yet 
func (c *Controller) getControllers() []Controller {
	
	return []Controller{}
}

func (auto *Automator) UpdateControllers(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	
	log.Println(r.Body)
	ctls_blob, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("converting error")
		respondWithError(w, http.StatusBadRequest, "converting error")
	}

	log.Println("controller setting")
	
	/* var ctls []string
	err = json.Unmarshal(ctls_blob, &ctls)
	if err != nil {
		log.Println("unmarshalling blob fail")
		log.Println(ctls_blob)
		respondWithError(w, http.StatusBadRequest, "unmarshalling blob fail")
	} */
	
	err = bgdb.Update(func(txn *badger.Txn) error {
		
		if err := txn.Set(CTLS_KEY, ctls_blob); err != nil {
			log.Println("set CTLS_COUNT_KEY fail")
			log.Println(err)
			respondWithError(w, http.StatusBadRequest, "setting fail")
			return err
		} else {
			log.Println("set CTLS_COUNT_KEY successful")
			
			// resetting devices
			GetDevices()
			
			automator.cron.Stop()
			automator.setup_cron()
			
			respondWithJson(w, http.StatusCreated, "{ success: true }")
			return nil
		}	
		respondWithJson(w, http.StatusCreated, "{ success: false }")
		return nil
		// // set_res := []byte("");
		// set_res := bytes.Buffer{}
  	// set_res.WriteString("[")
  
		// // result := buf.String()
	
		// var successful_set int
		// // iterate with ctls 	
		// for i, ctl := range ctls {
		// 	fmt.Println("inserting ctl: ", i, ctl)
		// 	// id := binary.BigEndian.Int()	
		// 	id := []byte(strconv.Itoa(i+1))
		// 	if err != nil {
		// 		fmt.Println(err, id)
		// 	}
		// 	if err := txn.Set(append(CTL_KEY, id...), []byte(ctl)); err != nil {
		// 		// // set_res = append(set_res, []byte({"{id: ", i+1,", success: true},"}))
		// 		set_res.WriteString("{id: ")
		// 		set_res.WriteString(string(id))
		// 		set_res.WriteString(", success: true},")
		// 		successful_set++
		// 		log.Println("set ctl", i+1, " successful")
		// 	} else {
		// 		// set_res = append(set_res,[]byte("{id: ", i+1,", success: false},"))
		// 	}

		// }	
		
		// set_res.WriteString("]")
		// respondWithJson(w, http.StatusCreated, set_res.String())
		
		// if err := txn.Set(CTLS_COUNT_KEY, []byte(strconv.Itoa(successful_set))); err != nil {
		// 	log.Println("set CTLS_COUNT_KEY successful")
		// } else {
		// 	log.Println("set CTLS_COUNT_KEY fail")
		// }	
			
		// item, err := txn.Get(CTLS_COUNT_KEY)
		// if err == nil {
		// 	_ = item.Value(func(val []byte) error {

		// 		// TODO: controllers backup with key controllers_setId
		// 		// with active_controllers_id_key

		// 		// Accessing val here is valid.
		// 		log.Println("chaged item: ", string(val))

		// 		var new_count = 0
	
		// 		return nil
		// 	})
		// } 
		
	})
		
	if err != nil {
		log.Println("update error: ", err)
		respondWithJson(w, http.StatusCreated, "{status: 'fail'}")
	}
	
}

func ControllerList(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	
	log.Println("ControllerList")
	
	var valCopy []byte
	err := bgdb.View(func(txn *badger.Txn) error {
		
		item, err := txn.Get(CTLS_KEY) 
		if err != nil {
			log.Println("error getting ctls")
			respondWithError(w, http.StatusBadRequest, "cannot get ctls")
			return nil
		}
		err = item.Value(func(val []byte) error {
			log.Println("successful getting ctls")
			valCopy = append([]byte{}, val...)
			respondWithJson(w, http.StatusCreated, string(valCopy))
			
	// 			// This func with val would only be called if item.Value encounters no error.
	// 			log.Println("controllers_count: ", val)
	// 			ctl_n := binary.BigEndian.Uint16(val)	
				
	// 			var ctls []Controller = []Controller{}

	// 			opt := badger.IteratorOptions
	// 			opt.PrefetchSize = 10
	// 			opt.PrefetchValues = true
	// 			opt.Prefix = []byte("ctl")
				
	// 			var output_ctls = "";
				
	// 			var count int
	// 			// err = db.View(func(txn *Txn) error {
	// 			it := txn.NewIterator(opt)
	// 			defer it.Close()
	// 			for it.Rewind(); it.Valid(); it.Next() {
	// 				it.Item().Value(func(val) {
	// 					output_ctls = append(output_ctls, val)
	// 					count++
	// 				})
	// 			}
	// 			// return nil
	// 			// })

	// 			// TODO loop through ctl1, ctl2, ctl3
	// 			// for i := 1; i < ctl_n; i++ {
						
	// 			// 	ctls.append()
	// 			// }
	// 			if count > 0 {
	// 				respondWithJson(w, http.StatusCreated, output_ctls )	
	// 			} else {
	// 				respondWithJson(w, http.StatusCreated, "[]" )	
	// 			}
				
	// 			// Copying or parsing val is valid.
	// 			// valCopy = append([]byte{}, val...)
				
			return nil
		})
		if err != nil {
			// respondWithError(w, http.StatusBadRequest, "getting ctls fail")
			log.Println("no value")
			return nil
		}
		// } 
		return nil
	})
	
	if err != nil {
		log.Println("err 3", err)
		respondWithError(w, http.StatusCreated, "initializing bgdb view error")
	}

}
		
func (auto *Automator) run_controllers() error {
	log.Println("Running Controllers")
	var ctls_json []byte
	// var dvs_json []byte
	err := bgdb.View(func(txn *badger.Txn) error {
		
		// Getting controllers into auto.ctls
		item, err := txn.Get(CTLS_KEY) 
		if err != nil {
			log.Println("Error getting controllers")
			return nil
		}
		err = item.Value(func(val []byte) error {
			log.Println("Successful getting controllers")
			ctls_json = append([]byte{}, val...)
			return nil
		})
		if err != nil {
			log.Println("No value on controllers")
			return nil
		}
		log.Println(string(ctls_json))
		
		var controllers []Controller
		err = json.Unmarshal(ctls_json, &controllers)
		if err != nil {
			log.Println(err)
			return nil
		}
		
		log.Println("Controlers: ", controllers)
		
		if len(controllers) > 0 {
			for _, ctl := range controllers {
				
				auto.ctls[ctl.Id] = ctl;
				 
				log.Printf("ctl: %d \n", ctl.Id)
				
			}
		}
		
		if len(auto.ctls) > 0 {
			for _, ctl := range auto.ctls {
				if auto.check_controller(ctl.Id) {
					
					log.Println("Add function for controller ", ctl.Id)
					
					if ctl.Policy == TIME_POLICY {
						// activate only BoolTrueDevices on time
						// if ctl.Cron != Cron{} {
							if len(ctl.BoolTrueDevices) > 0 {
								f := activate_by_time(ctl.Id, auto);
								cron := concatCronForTimePolicy(&ctl.Cron)
								entryId, err := auto.cron.AddFunc(cron, f)
								if err != nil {
									log.Println("\nError adding function f: ", err)
								}
								log.Println("EntryId ", entryId)
								auto.ctl_entry_map[ctl.Id] = entryId
								// log.Println(auto.ctl_entry_map)
								log.Println("TIME POLICY ", ctl.Id)
								
								// d := deactivate_by_time(ctl.Id, auto);
								// auto.cron.AddFunc(cron, d)
							}
						// }
							
					} else if ctl.Policy == MEASUREMENT_POLICY {
						// activate based on measurement always
						// activate_by_time(ctl.Id, auto);
						f := resolve_control_by_measurement(&ctl, auto)
								
						// activate every 10 seconds
						// auto.cron.AddFunc("*/10 * * * * *", f)
						entryId, err := auto.cron.AddFunc("*/10 * * * * *", f)
						if err != nil {
							log.Println("\nError adding function f: ", err)
						}
						log.Println("EntryId ", entryId)
						auto.ctl_entry_map[ctl.Id] = entryId
						// log.Println(auto.ctl_entry_map)
						log.Println("MEASUREMENT POLICY ", ctl.Id)
						
					} else if ctl.Policy == TIME_MEASUREMENT_POLICY {
						// sensor activate based on time and measurement
						// activate_by_time(ctl.Id, auto);
						f := resolve_control_by_measurement(&ctl, auto)
						cron := concatCron(&ctl.Cron)
						// automator.cron.AddFunc(cron, f)
						entryId, err := auto.cron.AddFunc(cron, f)
						if err != nil {
							log.Println("\nError adding function f: ", err)
						}
						log.Println("EntryId ", entryId)
						auto.ctl_entry_map[ctl.Id] = entryId
						// log.Println(auto.ctl_entry_map)
						log.Println("TIME MEASUREMENT POLICY ", ctl.Id)
						// if ctl.Cron != Cron{} {
							
						// }
						
					}
				
						
				}
			}
		}
			
		return nil
	})
	return err
}

func concatCron(cron *Cron) string {
	return "*/" + cron.Second + " " + cron.Minute + " " + cron.Hour + " " + cron.Dom + " " + cron.Month + " " + cron.Dow
}

// for time policy cron second is 60
func concatCronForTimePolicy(cron *Cron) string {
	return "*/60" + " " + cron.Minute + " " + cron.Hour + " " + cron.Dom + " " + cron.Month + " " + cron.Dow
}

// not implemented
func (auto *Automator) check_controller(ctl_id uint16) bool {
	// if ctl, exist := auto.ctls[ctl_id]; exist {
	// 	if auto.ctls[ctl_id] != nil {
			// fmt.Println(ctl)
	// 	}
	// 	return true
	// }
	return true
}
// ctl_id uint16
type Control func() 

// func resolve_control_by_measurement(ctl *Controller) {
	
// }

// resovlign control decision based on measurement from sensor
func resolve_control_by_measurement(controller *Controller, autom *Automator) Control {
	log.Println("resolve_control_by_measurement")
	ctl := controller
	auto := autom
	var sensor_value float32 = ctl.OptimalVal;
	var sensor_bool bool = true;
	return func() {
		// fmt.Println("\nresolve_control_by_measurement: ", ctl.Id)
		// log.Println("\n\nControlling Code ", ctl.ControlStatus)	
		// log.Println(auto.sensor_values)
		// read sensor from device
		// read_sensor
		log.Println(auto.devices)
		if len(ctl.Sensor) > 0 && reflect.TypeOf(auto.devices[ctl.Sensor]) != nil {
			if dv, ok := auto.devices[ctl.Sensor]; ok {
				dv.read_sensor(auto)
			}
		}
		
		// whether update value successful
		var update_fail = false;
		// update sensor_value from auto.sensor_values[device_id]SensorValue
		if len(ctl.Sensor) > 0 && ctl.Factor > 0 {
			log.Println("updating value from sensor")
			log.Println(auto.sensor_values)
			for _, sv := range auto.sensor_values {
				if len(sv.Device_id) > 0 && len(sv.Device_id) > 0 && len(sv.Value) > 0 && sv.Device_id == ctl.Sensor && uint8(ctl.Factor) == sv.Factor  {
					if sv.Is_boolean == "true" {
						bl, err := strconv.ParseBool(sv.Value)
						if err != nil {
							update_fail = true
							log.Println("Parsing bool false ", err)
						}
						log.Println("Updating boolean ", bl)
						sensor_bool = bl
					} else {
						vl, err := strconv.ParseFloat(sv.Value, 32)
						if err != nil {
							update_fail = true
							log.Println("Parsing float false ", err)
						}
						sensor_value = float32(vl)
						log.Println("Updating float ", vl)
					}
				}
			} 
		}

		// decide scheme boolean or value
		var err error = nil
		if ctl.ControlScheme == VALUE_CONTROL && !update_fail {
			log.Println("VALUE CONTROL")
			log.Println("Sensor Value: ", fmt.Sprintf("%f", sensor_value))
			log.Println("Optimal Value: ", ctl.OptimalVal)
			log.Println("Preferred Max: ", ctl.PreferredMax)
			log.Println("Preferred Min: ", ctl.PreferredMin)
			
			if sensor_value > ctl.PreferredMax && 
			ctl.ControlStatus != CS_DECREASING_DEVICES_ACTIVE {
				log.Println("Decreasing Devices Active")

				// activate DecreasingDevices until reaching OptimalVal
				if err = activateDevices(ctl.DecreasingDevices, auto); err == nil {
					ctl.ControlStatus = CS_DECREASING_DEVICES_ACTIVE
				}
				
			} else if sensor_value < ctl.PreferredMin && 
				ctl.ControlStatus != CS_INCREASING_DEVICES_ACTIVE {
				log.Println("Increasing Devices Active")	
				// activate IncreasingDevices until reaching OptimalVal
				if err = activateDevices(ctl.IncreasingDevices, auto); err == nil {
					ctl.ControlStatus = CS_INCREASING_DEVICES_ACTIVE
				}
				
				
			} else if ((sensor_value > ctl.OptimalVal && 
				ctl.ControlStatus == CS_INCREASING_DEVICES_ACTIVE) || (sensor_value < ctl.OptimalVal && 
				ctl.ControlStatus == CS_DECREASING_DEVICES_ACTIVE)) {
					
				log.Println("Devices Onhold")	

				if ctl.ControlStatus == CS_INCREASING_DEVICES_ACTIVE {
					if err = deactivateDevices(ctl.IncreasingDevices, auto); err == nil {
						ctl.ControlStatus = CS_ONHOLD
					}	
				} else if ctl.ControlStatus == CS_DECREASING_DEVICES_ACTIVE {
					if err = deactivateDevices(ctl.DecreasingDevices, auto); err == nil {
						ctl.ControlStatus = CS_ONHOLD
					}	
				} 
				
			}
			
			// simulating increasing value of sensor
			// if ctl.ControlStatus == CS_INCREASING_DEVICES_ACTIVE {
			// 	sensor_value = sensor_value + 1
			// } else if ctl.ControlStatus == CS_DECREASING_DEVICES_ACTIVE {
			// 	sensor_value = sensor_value - 1
			// } else {
			// 	sensor_value = sensor_value - 1
			// }
			
		} else if ctl.ControlScheme == BOOLEAN_CONTROL && !update_fail {
			log.Println("BOOLEAN_CONTROL")
			log.Println("Boolean: ", sensor_bool)
			
			if sensor_bool == true && 
				ctl.ControlStatus != CS_BOOL_TRUE_DEVICES_ACTIVE {

				log.Println("True Devices Active")		
					// activate BoolTrueDevices
					if err = activateDevices(ctl.BoolTrueDevices, auto); err == nil {
						ctl.ControlStatus = CS_BOOL_TRUE_DEVICES_ACTIVE
					}
					
			} else if sensor_bool == false && 
				ctl.ControlStatus != CS_BOOL_FALSE_DEVICES_ACTIVE {
					
				log.Println("False Devices Active")	
				// activate BoolFalseDevices
				if err = activateDevices(ctl.BoolFalseDevices, auto); err == nil {
					ctl.ControlStatus = CS_BOOL_FALSE_DEVICES_ACTIVE
				}
				
			}
			// sensor_bool = !sensor_bool
		}
		// entries := auto.cron.Entries()
		// if len(entries) > 0 {
		// 	fmt.Println(entries)
		// }
	}
}

func activateDevices(device_ids []string, auto *Automator) error {
	log.Println("Activate Devices ", device_ids)
	for _, did := range device_ids { 
		device, ok := auto.devices[did] 
		if ok {
			pin := rpio.Pin(device.GPIO)
			pin.Output()
			pin.High()
			log.Println("Activating Pin ", device.Pin)
		} else if device.GPIO > 0 {
			pin := rpio.Pin(device.GPIO)
			pin.Output()
			pin.Low()
			log.Println("Deactivate Pin ", device.Pin)
		}
	}
	return nil
}

func deactivateDevices(device_ids []string, auto *Automator) error {
	log.Println("Deactivate Devices ", device_ids)
	for _, did := range device_ids { 
		device, ok := auto.devices[did]
		if ok {
			pin := rpio.Pin(device.GPIO)
			pin.Output()
			pin.Low()
			log.Println("Deactivate Pin ", device.Pin)
		} else if device.GPIO > 0 {
			pin := rpio.Pin(device.GPIO)
			pin.Output()
			pin.High()
			log.Println("Activate Pin ", device.Pin)
		}
	}
	return nil
}

// 
func activate_by_time(controller_id uint16, auto *Automator) Control {
	ctl_id := controller_id
	autoIn := auto
	to_deactivate_next := 0
	return func() {
		log.Println("activate by time: ", ctl_id)
		if ctl, ok := autoIn.ctls[ctl_id]; ok {
				
			if len(ctl.BoolTrueDevices) > 0 { // activate && 
				activateDevices(ctl.BoolTrueDevices, autoIn)
				to_deactivate_next = to_deactivate_next + 1
				time.Sleep(71 * time.Second)
				if to_deactivate_next <= 1 {
					deactivateDevices(ctl.BoolTrueDevices, autoIn)
				}
				to_deactivate_next = to_deactivate_next - 1
				log.Println("to deactivate next")
				log.Println(to_deactivate_next)
			}
			
			// check if this controller's cron entry has Next time or not. If it has, do nothing. if it hasn't, deactivate devices immediately
			// if entryId, ok := auto.ctl_entry_map[ctl_id]; ok {
			// 	log.Println("entryId exists")
			// 	entry := auto.cron.Entry(entryId) 
			// 	if entry.Valid() {
			// 		log.Println("valid")
			// 		log.Println(entry.Next)
			// 		if entry.Next.IsZero() {
			// 			log.Println("deactivateDevices because Next is Zero")
			// 			deactivateDevices(autoIn.ctls[ctl_id].BoolTrueDevices, autoIn)
			// 			activate = false
			// 		}	
			// 	}
			// }
			// 
			// NOT WORK since NEXT is forever
			
		}
			
			
	}
}

/* func deactivate_by_time(controller_id uint16, auto *Automator) Control {
	ctt_id := controller_id
	autoIn := auto
	return func() {
		log.Println("deactivate_by_time: ", ctt_id)
		if autoIn.ctls != nil && autoIn.ctls.get(ctl) && autoIn.ctls[ctl].BoolTrueDevices {
			activateDevices(autoIn.ctls[ctl].BoolTrueDevices, autoIn)
		}
	}
} */





