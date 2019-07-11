<!--
title: MQTT
weight: 4705
-->
# MQTT
This activity allows you to send message on MQTT Queue.

## Installation

### Flogo CLI
```bash
flogo install github.com/project-flogo/edge-contrib/activity/mqtt
```

## Configuration

### Settings:
| Name         | Type   | Description
| :---         | :---   | :---
| broker       | string | The broker URL - ***REQUIRED***
| id           | string | The id of client - ***REQUIRED***
| username     | string | The name of the user
| password     | string | The password of the user
| store        | string | The store for message persistence
| cleanSession | bool   | Clean session flag
| topic        | string | The topic to publish to - ***REQUIRED***
| retain       | bool   | Retain Messages       
| qos          | int    | The quality of service
| sslConfig    | object | SSL configuration

 #### *sslConfig* Object:
 | Property      | Type   | Description
 |:---           | :---   | :---     
 | skipVerify    | bool   | Skip SSL validation, defaults to true
 | useSystemCert | bool   | Use the systems root certificate file, defaults to true
 | caFile        | string | The path to PEM encoded root certificates file
 | certFile      | string | The path to PEM encoded client certificate
 | keyFile       | string | The path to PEM encoded client key

 *Note: used if broker URI is ssl*

#### Topics
A substitution syntax is supported. For example if the topic is '/x/:/y/:' then the first ':' will be substituted with the value stored in `topicParams` `input` key '0' and the second ':' with key '1'. This can also be done with names: '/x/:param1/y/:param2'. The keys used for substitution are now 'param1' and 'param2'

### Input:

| Name        | Type   | Description
| :---        | :---   | :---
| message     | string | The message to send  
| topicParams | params | The topic parameters

### Output:

| Name  | Type   | Description
| :---  | :---   | :---
| data  | string | The data recieved

## Example

```json
{
  "id": "mqtt-activity",
  "name": "MQTT Activity",
  "description": "MQTT Example",
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
