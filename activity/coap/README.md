# CoAp
This activity allows you to send a CoAP message.

## Installation

### Flogo CLI
```bash
flogo install github.com/project-flogo/edge-contrib/activity/coap
```

## Configuration

### Settings:
| Name    | Type   | Description
| :---    | :---   | :---
| method  | string | The CoAP method to use (allowed values are GET, POST, PUT, DELETE)  - ***REQUIRED***   
| uri     | string | The CoAP resource URI - ***REQUIRED***
| type    | string | The message type (allowed values are Confirmable, NonConfirmable, Acknowledgement, Reset),  *Confirmable* is the default 
| options | string | The CoAP options to set     

### Input: 

| Name       | Type   | Description
| :---       | :---   | :---
| queryParams| string | The query params of the CoAP message    
| payload    | string | The payload of the CoAP message   
| messageId  | int    | ID used to detect duplicates and for optional reliability
 

### Output:

| Name       | Type   | Description
| :---       | :---   | :---
| response   | string | The response

## Example

```json
{
  "id": "coap-activity",
  "name": "Coap Activity",
  "description": "CoAP Get Example",
  "activity": {
    "ref": "github.com/project-flogo/edge-contrib/activity/coap",
    "settings": {
      "method" : "GET",
      "uri": "coap://localhost:5683/flogo"
    },
    "input" : {
        "payload" : "Hello from Flogo"
    }
  }
}
```