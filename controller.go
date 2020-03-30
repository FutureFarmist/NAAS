package main

import (
	_ "bytes"
	_ "encoding/binary"
	_ "encoding/json"
	_ "fmt"
	_ "github.com/gorilla/mux"
	"log"
	"net/http"
	_ "github.com/stianeikeland/go-rpio"
	_ "strconv"
	"io/ioutil"
	
	badger "github.com/dgraph-io/badger"
	// _ IteratorOptions "github.com/dgraph-io/badger/v2/options"
)

var (
	// controller (ctl)

	CTL_KEY = []byte("ctl")
	
	CTLS_KEY = []byte("ctls")
	
	// ctl1, ctl2, ctl3, ...
	CTLS_COUNT_KEY = []byte("ctls_count")
	
	// ctls_update_DATE
	CTLS_UPDATE_KEY = []byte("ctls_update")

	// every update controllers would give it controller 
	// ctls_ID
	// active_controllers_id_key = []byte("active_controllers_id")
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
		Sensors []string `bson:"Sensors,omitempty" json:"Sensors,omitempty"`	
		
		Factor string  `bson:"Factor,omitempty" json:"Factor,omitempty"`	
		
		// DeviceLinks []DeviceLink `bson:"DeviceLinks,omitempty" json:"DeviceLinks,omitempty"`
		
		/* ENUM
		TIME_POLICY - control based on time
		MEASUREMENT_POLICY - control based on value, this includes VALUE_CONTROL and BOOLEAN_CONTROL scheme
		TIME_MEASUREMENT_POLICY - control value for period(s) of time
		*/
		Policy uint8 `bson:"Policy,omitempty" json:"Policy,omitempty"`	
		
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
		OptimalVal	uint
		PreferredMin	uint
		PreferredMax	uint
		
		// for value control 
		IncreasingDevices []string
		DecreasingDevices []string
		
		// BOOLEAN_CONTROL policy
		
		// true/false
		// This setting constrained by time period, not by value
		// PreferredState	bool
		
		// acting on device when value is true
		BoolTrueDevices []string
		
		// acting on device when value is false
		BoolFalseDevices []string
		
		
		/* ENUM
		TimeType
		Period
		Daily
		Weekly
		Monthly
		Ondemand
		*/
		// TimeType Date `bson:"TimeType,omitempty" json:"TimeType,omitempty"`
		
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


// run this every day at 00:00:00
// this function help run cron of everyday
// first it clear active controller then it go through all controller
// to find which on is to run today
func (c *Controller) startScheduler() bool {
	
	return true
}

func (c *Controller) getControllers() []Controller {
	
	return []Controller{}
}

func UpdateControllers(w http.ResponseWriter, r *http.Request) {
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
			respondWithJson(w, http.StatusCreated, "{success: true}")
			return nil
		}	
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

func clear_ctls_key() {
	
	bgdb.Update(func(txn *badger.Txn) error {
		
		if err := txn.Set(CTLS_KEY, []byte("")); err != nil {
			log.Println("clear_ctls_key fail")
			log.Println(err)
			// respondWithError(w, http.StatusBadRequest, "clear_ctls_key fail")
			return err
		} else {
			log.Println("clear_ctls_key successful")
			// respondWithJson(w, http.StatusCreated, "clear_ctls_key")
			return nil
		}	
			return nil
	})
}
		