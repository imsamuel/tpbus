# A REST API for the timings of the bus services at Temasek Polytechnic *EDIT: CORS enabled!

## Preface

### Details
|||
|-------------|----------------------------------------------------------------------------------------------------------------------------------|
| URL         | https://tpbus.herokuapp.com |
| Description | Returns real-time Bus Arrival information of Bus Services at Temasek Polytechnic, including Est. Arrival Time, Destination Code. |
| Update Freq | 1 minute |

### Attributes
| Attributes             | Description                                                                   | Example                   |
|------------------------|-------------------------------------------------------------------------------|---------------------------|
| serviceNumber          | Bus service number                                                            | 15                        |
| destinationCode        | Reference code of the last bus stop where this bus will terminate its service | 77131                     |
| nextBus                | Bus level attributes for next bus                                             |                           |
| nextBus2               | Bus level attributes for subsequent bus                                       |                           |
| nextBus3               | Bus level attributes for bus after subsequent one                             |                           |
| ^ estimatedArrival     | Date-time of this bus’ estimated time of arrival expressed in the UTC standard, GMT+8 for Singapore Standard Time (SST)                                                                                                    | 2019-04-29T07:20:24+08:00 |
| ^ type                   | Vehicle type: single deck OR double deck OR bendy                           | single deck               |
| ^ isWheelChairAccessible | Indicates if bus is wheel-chair accessible: true OR false                   | true                      |

## Usage

All requests to a malformed URI will be returned with:

- `404 Not Found` API could not map to a resource
```json
{
    "id": "Resource could not be found",
    "description": "API could not find a resource mapped to <currentURLPath>"
}
```

### List the timings of all bus services at all bus stops

**Definition**

`GET /services`

**Response**

- `200 OK` on success
```json
{
    "westGate": [
        {
            "serviceNumber":"118",
            "destinationCode": 97009,
            "nextBus": {
                "estimatedArrival": "2019-11-12T17:52:49+08:00",
                "type": "double deck",
                "isWheelChairAccessible": true 
            },
            "nextBus2": {
                "estimatedArrival": "2019-11-12T17:58:24+08:00",
                "type": "bendy",
                "isWheelChairAccessible": false
            },
            "nextBus3": {
                "estimatedArrival": "2019-11-12T18:09:53+08:00",
                "type": "single deck",
                "isWheelChairAccessible": true
            }
        }
    ],
    "oppWestGate": [
        {
            "serviceNumber":"118",
            "destinationCode": "97009",
            "nextBus": {
                "estimatedArrival": "2019-11-12T17:52:49+08:00",
                "type": "single deck",
                "isWheelChairAccessible": true 
            },
            "nextBus2": {
                "estimatedArrival": "2019-11-12T17:58:24+08:00",
                "type": "double deck",
                "isWheelChairAccessible": true 
            },
            "nextBus3": {
                "estimatedArrival": "2019-11-12T18:09:53+08:00",
                "type": "bendy",
                "isWheelChairAccessible": false
            }
        }
    ],
    "mainGate": [
        {
            "serviceNumber":"118",
            "destinationCode": "97009",
            "nextBus": {
                "estimatedArrival": "2019-11-12T17:52:49+08:00",
                "type": "double deck",
                "isWheelChairAccessible": false
            },
            "nextBus2": {
                "estimatedArrival": "2019-11-12T17:58:24+08:00",
                "type": "single deck",
                "isWheelChairAccessible": true
            },
            "nextBus3": {
                "estimatedArrival": "2019-11-12T18:09:53+08:00",
                "type": "double deck",
                "isWheelChairAccessible": true
            }
        }
    ],
    "oppMainGate": [
        {
            "serviceNumber":"118",
            "destinationCode": "97009",
            "nextBus": {
                "estimatedArrival": "2019-11-12T17:52:49+08:00",
                "type": "bendy",
                "isWheelChairAccessible": true
            },
            "nextBus2": {
                "estimatedArrival": "2019-11-12T17:58:24+08:00",
                "type": "single deck",
                "isWheelChairAccessible": false
            },
            "nextBus3": {
                "estimatedArrival": "2019-11-12T18:09:53+08:00",
                "type": "single deck",
                "isWheelChairAccessible": true
            }
        }
    ],
    "eastGate": [
        {
            "serviceNumber":"118",
            "destinationCode": "97009",
            "nextBus": {
                "estimatedArrival": "2019-11-12T17:52:49+08:00",
                "type": "double deck",
                "isWheelChairAccessible": true
            },
            "nextBus2": {
                "estimatedArrival": "2019-11-12T17:58:24+08:00",
                "type": "single deck",
                "isWheelChairAccessible": true
            },
            "nextBus3": {
                "estimatedArrival": "2019-11-12T18:09:53+08:00",
                "type": "double deck",
                "isWheelChairAccessible": true
            }
        }
    ],
    "oppEastGate": [
        {
            "serviceNumber":"118",
            "destinationCode": "97009",
            "nextBus": {
                "estimatedArrival": "2019-11-12T17:52:49+08:00",
                "type": "double deck",
                "isWheelChairAccessible": true
            },
            "nextBus2": {
                "estimatedArrival": "2019-11-12T17:58:24+08:00",
                "type": "single deck",
                "isWheelChairAccessible": false
            },
            "nextBus3": {
                "estimatedArrival": "2019-11-12T18:09:53+08:00",
                "type": "single deck",
                "isWheelChairAccessible": true
            }
        }
    ]
}
```

### List the timings of all bus services at a specified bus stop

**Definition**

`GET /services/<busStopLocation>`

Where `busStopLocation` can be one of these 6 values:
- west-gate
- opp-west-gate
- main-gate
- opp-main-gate
- east-gate
- opp-east-gate

**Response**

- `200 OK` on success
```json
[
    {
        "serviceNumber":"118",
        "destinationCode": "97009",
        "nextBus": {
            "estimatedArrival": "2019-11-12T17:52:49+08:00",
            "type": "double deck",
            "isWheelChairAccessible": true
        },
        "nextBus2": {
            "estimatedArrival": "2019-11-12T17:58:24+08:00",
            "type": "single deck",
            "isWheelChairAccessible": false
        },
        "nextBus3": {
            "estimatedArrival": "2019-11-12T18:09:53+08:00",
            "type": "single deck",
            "isWheelChairAccessible": true
        }
    },
    {
        "serviceNumber":"129",
        "destinationCode": "97009",
        "nextBus": {
            "estimatedArrival": "2019-11-12T17:52:49+08:00",
            "type": "single deck",
            "isWheelChairAccessible": false
        },
        "nextBus2": {
            "estimatedArrival": "2019-11-12T17:58:24+08:00",
            "type": "bendy",
            "isWheelChairAccessible": true
        },
        "nextBus3": {
            "estimatedArrival": "2019-11-12T18:09:53+08:00",
            "type": "double deck",
            "isWheelChairAccessible": true
        }
    }
]
```

### List the timings of a single bus service at a specified bus stop

**Definition**

`GET /services/<busStopCode>/<serviceNumber>`

Where `busStopLocation` can be one of these six values:
- west-gate
- opp-west-gate
- main-gate
- opp-main-gate
- east-gate
- opp-east-gate

and `serviceNumber` can be one of these eight values:
- 8
- 15
- 23
- 69
- 118
- 118B
- 129
- 518

**Response** (Example request to `/services/east-gate/118`)

- `200 OK` on success
```json
{
    "serviceNumber": 129,
    "destinationCode": 97009,
    "nextBus": {
        "estimatedArrival": "2019-11-12T17:52:49+08:00",
        "type": "double deck",
        "isWheelChairAccessible": true
    },
    "nextBus2": {
        "estimatedArrival": "2019-11-12T17:52:49+08:00",
        "type": "single deck",
        "isWheelChairAccessible": false
    },
    "nextBus3": {
        "estimatedArrival": "2019-11-12T17:52:49+08:00",
        "type": "bendy",
        "isWheelChairAccessible": true
    }
}
```


