# MQTT
This activity allows you to listen to message on Mqtt Queue.

## Installation

### Flogo CLI
```bash
flogo install github.com/project-flogo/edge-contrib/trigger/mqtt
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

### Handler Settings
| Name      | Type   | Description
| :---      | :---   | :---
| cleansess | string | Clean sesssion
| topic | string | The topic to publish to
| qos | string | The quality of service

 
### Output: 

| Name    | Type   | Description
| :---    | :---   | :---
| message | string | The message recieved
    
### Reply:

| Name  | Type   | Description
| :---  | :---   | :---
| data  | object | The data recieved

## Example

```json
{
  "id": "mqtt-trigger",
  "name": "Mqtt Trigger",
  "ref": "github.com/project-flogo/edge-contrib/trigger/mqtt",
  "settings": {
      "broker" : "tcp://localhost:1883",
     	"id":"client-1"
  },
  "handlers": {
    "settings": {
    	"topic": "flogo",
    	"qos": "0"
    
    }
  }
}
```