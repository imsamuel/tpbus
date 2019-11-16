package handlers

import "tpbus/constants"

// Path to retrieve the bus service timings from all bus stops at Temasek Poly.
var SERVICES_PATH = "/services"

// Path to retrieve the bus service timings from a specified bus stop.
var BUS_STOP_LOCATION_PATH = "/services/:busStopLocation"

// "west-gate"
var WEST_GATE = constants.BusStopsAtGates.West.Name
