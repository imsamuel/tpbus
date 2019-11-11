/*
Package constants contains values which are commonly used throughout the
application.
*/
package constants

// Container for a bus stop's details of a bus service.
type busStopOfService struct {
	Name string // name of bus stop
	Code string  // bus stop code of bus stop
}

// Container for the details of a bus stop at Temasek Polytechnic.
type busStopAtGate struct {
	Name string // the Temasek Poly gate in which the bus stop is located at
	Code string // bus stop code of bus stop
}

/*
Holds the data about the bus stops at the beginning/end
destinations of the bus bus_services at Temasek Polytechnic.
*/
var BusStopsOfServices = struct {
	BefPunggolRd busStopOfService
	ChangiBusinessPkTer busStopOfService
	OppITECol busStopOfService
	PasirRisInt busStopOfService
	PunggolTempInt busStopOfService
	StMichaelTer busStopOfService
	TampinesInt busStopOfService
	TampinesConcourseInt busStopOfService
	ToaPayohInt busStopOfService
}{
	busStopOfService{"Bef Punggol Rd", "65191"},
	busStopOfService{"Changi Biz Pk Ter", "97009"},
	busStopOfService{"Opp ITE Col", "96111"},
	busStopOfService{"Pasir Ris Int", "77009"},
	busStopOfService{"Punggol Temp Int", "65009"},
	busStopOfService{"St. Michael's Ter", "52499"},
	busStopOfService{"Tamp Int", "75009"},
	busStopOfService{"Tamp Conc Int", "75019"},
	busStopOfService{"Toa Payoh Int", "52009"},
}

// Holds the data of the bus stops at Temasek Polytechnic.
var BusStopsAtGates = struct {
	West busStopAtGate
	OppWest busStopAtGate
	Main busStopAtGate
	OppMain busStopAtGate
	East busStopAtGate
	OppEast busStopAtGate
}{
	busStopAtGate{"west", "75249"},
	busStopAtGate{"opp-west", "75241"},
	busStopAtGate{"main", "75239"},
	busStopAtGate{"opp-main", "75231"},
	busStopAtGate{"east", "75229"},
	busStopAtGate{"opp-east", "75221"},
}
