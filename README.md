# A REST API for the timings of the bus services at Temasek Polytechnic

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
            "nextBus": "2019-11-12T17:52:49+08:00",
            "nextBus2": "2019-11-12T17:58:24+08:00",
            "nextBus3": "2019-11-12T18:09:53+08:00",
            "type": "double deck",
            "isWheelChairAccessible": true
         }
    ],
    "oppWestGate": [
        {
            "serviceNumber":"118",
            "destinationCode": "97009",
            "nextBus": "2019-11-12T17:52:49+08:00",
            "nextBus2": "2019-11-12T17:58:24+08:00",
            "nextBus3": "2019-11-12T18:09:53+08:00",
            "type": "single deck",
            "isWheelChairAccessible": false
         }
    ],
    "mainGate": [
        {
            "serviceNumber":"118",
            "destinationCode": "97009",
            "nextBus": "2019-11-12T17:52:49+08:00",
            "nextBus2": "2019-11-12T17:58:24+08:00",
            "nextBus3": "2019-11-12T18:09:53+08:00",
            "type": "double deck",
            "isWheelChairAccessible": false
         }
    ],
    "oppMainGate": [
        {
            "serviceNumber":"118",
            "destinationCode": "97009",
            "nextBus": "2019-11-12T17:52:49+08:00",
            "nextBus2": "2019-11-12T17:58:24+08:00",
            "nextBus3": "2019-11-12T18:09:53+08:00",
            "type": "bendy",
            "isWheelChairAccessible": true
         }
    ],
    "eastGate": [
        {
            "serviceNumber":"118",
            "destinationCode": "97009",
            "nextBus": "2019-11-12T17:52:49+08:00",
            "nextBus2": "2019-11-12T17:58:24+08:00",
            "nextBus3": "2019-11-12T18:09:53+08:00",
            "type": "double deck",
            "isWheelChairAccessible": true
         }
    ],
    "oppEastGate": [
        {
            "serviceNumber":"118",
            "destinationCode": "97009",
            "nextBus": "2019-11-12T17:52:49+08:00",
            "nextBus2": "2019-11-12T17:58:24+08:00",
            "nextBus3": "2019-11-12T18:09:53+08:00",
            "type": "single deck",
            "isWheelChairAccessible": true
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
        "nextBus": "2019-11-12T17:52:49+08:00",
        "nextBus2": "2019-11-12T17:58:24+08:00",
        "nextBus3": "2019-11-12T18:09:53+08:00",
        "type": "single deck",
        "isWheelChairAccessible": false
    },
    {
        "serviceNumber":"129",
        "destinationCode": ,
        "nextBus": "2019-11-12T17:56:21+08:00",
        "nextBus2": "2019-11-12T18:08:15+08:00",
        "nextBus3": "2019-11-12T18:29:48+08:00",
        "type": "bendy",
        "isWheelChairAccessible": true
    }
]
```

### List the timings of a single bus service at a specified bus stop

**Definition**

`GET /services/<busStopCode>?ServiceNumber=<serviceNumber>`

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

**Response** (Example request to `/services/east-gate?ServiceNumber=118`)

- `200 OK` on success
```json
{
    "serviceNumber":"118",
    "destinationCode": "97009",
    "nextBus": "2019-11-12T17:52:49+08:00",
    "nextBus2": "2019-11-12T17:58:24+08:00",
    "nextBus3": "2019-11-12T18:09:53+08:00",
    "type": "double deck",
    "isWheelChairAccessible": false
}
```


