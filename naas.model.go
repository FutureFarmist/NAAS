package main

// type (
// 	Field struct {
// 		FieldId 			string		`bson:"Match,omitempty" json:"Match,omitempty"`
// 		Name					string		`bson:"Match,omitempty" json:"Match,omitempty"`
		
// 		GroundWidth 	uint8			`bson:"Match,omitempty" json:"Match,omitempty"`
//     GroundDepth 	uint8			`bson:"Match,omitempty" json:"Match,omitempty"`
//     GroundHeight 	uint8			`bson:"Match,omitempty" json:"Match,omitempty"`
//     AirHeight			uint8			`bson:"Match,omitempty" json:"Match,omitempty"`

// 		/* Plant ids */
// 		Plants 				[]string	`bson:"Match,omitempty" json:"Match,omitempty"`
		
// 		/* Device id */
// 		Devices 			[]string	`bson:"Match,omitempty" json:"Match,omitempty"`
		
// 	}

// 	Device struct {
// 		DeviceId 		string 				`bson:"Match,omitempty" json:"Match,omitempty"`
// 		Name 				string 				`bson:"Match,omitempty" json:"Match,omitempty"`
		
// 		Kind			// ENUM SENSOR, ACTUATOR

// 		/* These are for sensor */
// 		Factor			// ENUM				`bson:"Match,omitempty" json:"Match,omitempty"`
// 		Unit				// ENUM				`bson:"Match,omitempty" json:"Match,omitempty"`

// 		/* This is for actuator */
// 		ControlScheme // ENUM			`bson:"Match,omitempty" json:"Match,omitempty"`
		
// 		/* These are for sensor */
// 		ValueState
// 		BoolState
		
// 	}

// 	Plant struct {
// 		PlantId			string				`bson:"Match,omitempty" json:"Match,omitempty"`
//     Name				string				`bson:"Match,omitempty" json:"Match,omitempty"`

// 		/* string point to device id */
//     Devices			[]string			`bson:"Match,omitempty" json:"Match,omitempty"`
	
// 		Genus				string				`bson:"Match,omitempty" json:"Match,omitempty"`
//     Species			string				`bson:"Match,omitempty" json:"Match,omitempty"`
//     LifeCycle		uint8					`bson:"Match,omitempty" json:"Match,omitempty"`
		
// 		// Mineral										`bson:"Match,omitempty" json:"Match,omitempty"`
		
//     // PreferredSoilHumidity			`bson:"Match,omitempty" json:"Match,omitempty"`
	
// 		// PreferredAirHumidity 			`bson:"Match,omitempty" json:"Match,omitempty"`
 
// 		// PreferredAirTemperature 	`bson:"Match,omitempty" json:"Match,omitempty"`

// 	}
	
// 	ControlCondition struct {
		
// 		SensorDevices	[]string
		
// 		ActuatorDevices	[]string
		
// 		ControlScheme	// ENUM BOOL, 
	
// 		/* TimeType
// 				Period
// 					Daily
// 					Weekly
// 					Monthly
// 				Ondemand 
// 		*/
// 		TimeType		// ENUM
		
// 		// Daytime HH:MM:SS
// 		DaytimeStart	//
// 		DaytimeEnd		//
		
// 		/* optional Season Start Date - End Date 
// 			Date DD:MM:YYYY in YYYY in CE (common era)
// 		*/
		
// 		SessionStartDate	
// 		SeasonEndDate
		
			
// 		/* value manipulation capability of device */
// 		ValueManipulation	// ENUM INCREASE, DECREASE, BOTH
		
// 		/* Range value optimization 
// 		NAAS allows the value to fructuate between PreferredMin and PreferredMax
// 		When it could only increase value, PreferredMin is evaluated if the value is below.
// 		When it could only decrease value, PreferredMax is evaluated if the value is above.
// 		If device can manipulate in both directions, both PreferredMin and PreferredMax are evaluated for control. 
// 		When, OptimalVal is set, controlling device would leave manipulating the value */
		
// 		// optimization is stop at this value
// 		OptimalVal
		
// 		PreferredMin
		
// 		PreferredMax
		
// 		// true/false
// 		// This setting constrains by time period, not by value
// 		PreferState
		
// 	}
	
// }