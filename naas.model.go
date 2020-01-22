package main

type (
	
	DeviceStatus uint16
	DeviceKind uint16
	ControlScheme uint16
	
	ValueManipulation uint8
	
	Field struct {
		FieldId string `bson:"Match,omitempty" json:"Match,omitempty"`
		Name string `bson:"Match,omitempty" json:"Match,omitempty"`
		
		GroundWidth int `bson:"Match,omitempty" json:"Match,omitempty"`
		GroundDepth int `bson:"Match,omitempty" json:"Match,omitempty"`
		GroundHeight int `bson:"Match,omitempty" json:"Match,omitempty"`
		AirHeight int `bson:"Match,omitempty" json:"Match,omitempty"`

		/* Plant ids */
		Plants []string `bson:"Match,omitempty" json:"Match,omitempty"`
		
		/* Device id */
		Devices []string `bson:"Match,omitempty" json:"Match,omitempty"`
		
	}

	Device struct {
		DeviceId string `bson:"DeviceId" json:"DeviceId"`
		Name string `bson:"Name,omitempty" json:"Name,omitempty"`

		// ENUM ACTIVE, INACTIVE, ERROR
		Status uint16 `bson:"Status,omitempty" json:"Status,omitempty"`
		
		// ENUM SENSOR, ACTUATOR
		Kind uint16 `bson:"Kind,omitempty" json:"Kind,omitempty"`

		/* These are for sensor */
		Factor uint16 `bson:"Factor,omitempty" json:"Factor,omitempty"`
		Unit uint16 `bson:"Unit,omitempty" json:"Unit,omitempty"`

		/* This is for actuator */
		ControlScheme uint16 `bson:"ControlScheme,omitempty" json:"ControlScheme,omitempty"`
		
		/* These are for sensor */
		ValueState uint16 `bson:"ValueState,omitempty" json:"ValueState,omitempty"`
		BoolState uint16 `bson:"BoolState,omitempty" json:"BoolState,omitempty"`
	
		/* Associated pin */
		Pin	uint8 `bson:"Pin,omitempty" json:"Pin,omitempty"`
	}

	// Plant struct {
		
	// 	PlantId string `bson:"PlantId,omitempty" json:"PlantId,omitempty"`
	// 	Name string `bson:"Name,omitempty" json:"Name,omitempty"`

	// 	/* string point to device id */
	// 	Devices []string `bson:"Devices,omitempty" json:"Devices,omitempty"`
		
	// 	// Genus string `bson:"Genus,omitempty" json:"Genus,omitempty"`
	// 	// Species string `bson:"Species,omitempty" json:"Species,omitempty"`
	// 	// LifeCycle uint8 `bson:"LifeCycle,omitempty" json:"LifeCycle,omitempty"`
	// 	// Mineral `bson:"Match,omitempty" json:"Match,omitempty"`
	// 	// PreferredSoilHumidity `bson:"Match,omitempty" json:"Match,omitempty"`
	// 	// PreferredAirHumidity `bson:"Match,omitempty" json:"Match,omitempty"`
	// 	// PreferredAirTemperature `bson:"Match,omitempty" json:"Match,omitempty"`

	// }
	
	// Controller struct {
	// 	SensorDevices []string `bson:"SensorDevices,omitempty" json:"SensorDevices,omitempty"`
	// 	ActuatorDevices []string `bson:"ActuatorDevices,omitempty" json:"ActuatorDevices,omitempty"`
		
	// 	// ENUM BOOL, VALUE 
	// 	ControlScheme `bson:"ControlScheme,omitempty" json:"ControlScheme,omitempty"`
		
	// 	/* ENUM
	// 	TIME_BASED
	// 	VALUE_BASED
	// 	*/
	// 	Policy uint16 `bson:"Policy,omitempty" json:"Policy,omitempty"`	
		
	// 	/* ENUM
	// 	TimeType
	// 	Period
	// 	Daily
	// 	Weekly
	// 	Monthly
	// 	Ondemand
	// 	*/
	// 	TimeType Date `bson:"TimeType,omitempty" json:"TimeType,omitempty"`
		
	// 	// Daytime HH:MM:SS
	// 	DaytimeStart `bson:"DaytimeStart,omitempty" json:"DaytimeStart,omitempty"`
	// 	DaytimeEnd `bson:"DaytimeEnd,omitempty" json:"DaytimeEnd,omitempty"`
		
	// 	/* optional Season Start Date - End Date
	// 	Date DD:MM:YYYY in YYYY in CE (common era)
	// 	*/
	// 	SessionStartDate	`bson:"SessionStartDate,omitempty" json:"SessionStartDate,omitempty"`
	// 	SeasonEndDate			`bson:"SeasonEndDate,omitempty" json:"SeasonEndDate,omitempty"`
		
	// 	/*
	// 	value manipulation capability of device
	// 	interface for a device are increase(), decrease()
	// 	ENUM INCREASE, DECREASE, BOTH
	// 	*/
	// 	ValueManipulation uint8 `bson:"ValueManipulation,omitempty" json:"ValueManipulation,omitempty"` 
		
	// 	/* Range value optimization
	// 	NAAS allows the value to fructuate between PreferredMin and PreferredMax
	// 	When it could only increase value, PreferredMin is evaluated if the value is below.
	// 	When it could only decrease value, PreferredMax is evaluated if the value is above.
	// 	If device can manipulate in both directions, both PreferredMin and PreferredMax are evaluated for control.
	// 	When, OptimalVal is set, controlling device would leave manipulating the value */
	// 	// optimization is stop at this value
	// 	OptimalVal	uint
	// 	PreferredMin	uint
	// 	PreferredMax	uint
		
	// 	// true/false
	// 	// This setting constrained by time period, not by value
	// 	PreferState	
		
	// }
	
)

// const (
// 	INACTIVE DeviceStatus = iota
// 	ACTIVE
// 	ERROR	
// )

// const (
// 	INACTIVE DeviceStatus = iota
// 	ACTIVE
// 	ERROR	
// )
