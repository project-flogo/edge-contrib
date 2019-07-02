<!--
title: MQTT
weight: 4705
-->
# MQTT
This trigger allows you to listen to messages on MQTT.

## Installation

### Flogo CLI
```bash
flogo install github.com/project-flogo/edge-contrib/trigger/mqtt
```

## Configuration

### Settings:
| Name          | Type   | Description
| :---          | :---   | :---
| broker        | string | The broker URL - ***REQUIRED***
| id            | string | The id of client - ***REQUIRED*** 
| username      | string | The user's name
| password      | string | The user's password
| store         | string | The store for message persistence
| cleanSession  | bool   | Clean session flag
| keepAlive     | int    | Keep Alive time in seconds
| autoReconnect | bool   | Enable Auto-Reconnect
| sslConfig     | object | SSL configuration
 
 #### *sslConfig* Object: 
 | Property      | Type   | Description
 |:---           | :---   | :---     
 | skipVerify    | bool   | Skip SSL validation, defaults to true
 | useSystemCert | bool   | Use the systems root certificate file, defaults to true
 | caFile        | string | The path to PEM encoded root certificates file
 | certFile      | string | The path to PEM encoded client certificate
 | keyFile       | string | The path to PEM encoded client key
 
 *Note: used if broker URI is ssl*
 
### Handler Settings
| Name       | Type   | Description
| :---       | :---   | :---
| topic      | string | The topic to listen on - ***REQUIRED***
| replyTopic | string | The topic to reply on   
| qos        | int    | The Quality of Service

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