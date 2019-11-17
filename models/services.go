package models

// Represents the details of an incoming bus.
type IncomingBus struct {
	EstimatedArrival string `json:"estimatedArrival,omitempty"`
	Type string `json:"type,omitempty"`
	IsWheelChairAccessible bool `json:"isWheelChairAccessible,omitempty"`
}

// Represents all details of a single bus service.
type Service struct {
	ServiceNumber string `json:"serviceNumber,omitempty"`
	DestinationCode string `json:"destinationCode,omitempty"`
	NextBus IncomingBus `json:"nextBus,omitempty"`
	NextBus2 IncomingBus `json:"nextBus2,omitempty"`
	NextBus3 IncomingBus `json:"nextBus3,omitempty"`
}

// Represents all bus services at a bus stop.
type Services []Service

// Defines the structure of how all TP's bus services are stored in this app.
type BusStops struct {
	WestGate Services `json:"westGate,omitempty"`
	OppWestGate Services `json:"oppWestGate,omitempty"`
	MainGate Services `json:"mainGate,omitempty"`
	OppMainGate Services `json:"oppMainGate,omitempty"`
	EastGate Services `json:"eastGate,omitempty"`
	OppEastGate Services `json:"oppEastGate,omitempty"`
}