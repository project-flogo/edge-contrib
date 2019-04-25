# MQTT
This activity allows you to send message on Mqtt Queue.

## Installation

### Flogo CLI
```bash
flogo install github.com/project-flogo/edge-contrib/activity/mqtt
```

## Configuration

### Settings:
| Name      | Type   | Description
| :---      | :---   | :---
| broker    | string | 	The broker URL
| id | string | The id of client 
| user | string | The name of the user
| password | string | The password of the user
| store | string | The path containing certificates
| cleansess | string | Clean sesssion
| topic | string | The topic to publish to
| qos | string | The quality of service
 
### Input: 

| Name       | Type   | Description
| :---       | :---   | :---
| message | string | The message to send  
    
### Output:

| Name  | Type   | Description
| :---  | :---   | :---
| data  | string | The data recieved

## Example

```json
{
  "id": "mqtt-activity",
  "name": "Mqtt Activity",
  "description": "Mqtt Example",
  "activity": {
    "ref": "github.com/project-flogo/edge-contrib/activity/mqtt",
    "settings": {
      "broker" : "tcp://localhost:1883",
      "qos": "0",
      "id":"client-1",
      "topic": "flogo"
    },
    "input" : {
        "message" : "Hello From Flogo"
    }
  }
}
```