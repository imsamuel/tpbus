package models

// Represents the details of an incoming bus.
type IncomingBus struct {
	EstimatedArrival string `json:"estimatedArrival"`
	Type string `json:"type"`
	IsWheelChairAccessible bool `json:"isWheelChairAccessible"`
}

// Represents all details of a single bus service.
type Service struct {
	ServiceNumber string `json:"serviceNumber"`
	DestinationCode string `json:"destinationCode"`
	NextBus IncomingBus `json:"nextBus"`
	NextBus2 IncomingBus `json:"nextBus2"`
	NextBus3 IncomingBus `json:"nextBus3"`
}

// Represents all bus services at a bus stop.
type Services []Service

// Defines the structure of how all TP's bus services are stored in this app.
type BusStops struct {
	WestGate Services `json:"westGate"`
	OppWestGate Services `json:"oppWestGate"`
	MainGate Services `json:"mainGate"`
	OppMainGate Services `json:"oppMainGate"`
	EastGate Services `json:"eastGate"`
	OppEastGate Services `json:"oppEastGate"`
}