package handlers

import "tpbus/constants"

// Path to retrieve the bus service timings from all bus stops at Temasek Poly.
var SERVICES_PATH = "/services"

// Path to retrieve the timings of all bus services at specified bus stop.
var BUS_STOP_LOCATION_PATH = "/services/:busStopLocation"

// Path to retrieve the timings of a single bus service at specified bus stop.
var SERVICE_PATH = "/services/:busStopLocation/:serviceNumber"

// "west-gate"
var WEST_GATE = constants.BusStopsAtGates.West.Name
