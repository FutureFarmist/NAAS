package main

type (
	
	// ValueManipulation uint8
	
	Field struct {
		Id string `bson:"Id,omitempty" json:"Id,omitempty"`
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

	Plant struct {
		
		PlantId string `bson:"PlantId,omitempty" json:"PlantId,omitempty"`
		Name string `bson:"Name,omitempty" json:"Name,omitempty"`

		/* string point to device id */
		Devices []string `bson:"Devices,omitempty" json:"Devices,omitempty"`
		
		// Genus string `bson:"Genus,omitempty" json:"Genus,omitempty"`
		// Species string `bson:"Species,omitempty" json:"Species,omitempty"`
		// LifeCycle uint8 `bson:"LifeCycle,omitempty" json:"LifeCycle,omitempty"`
		// Mineral `bson:"Match,omitempty" json:"Match,omitempty"`
		// PreferredSoilHumidity `bson:"Match,omitempty" json:"Match,omitempty"`
		// PreferredAirHumidity `bson:"Match,omitempty" json:"Match,omitempty"`
		// PreferredAirTemperature `bson:"Match,omitempty" json:"Match,omitempty"`

	}
	
)
