# CoAp
This activity allows you to Get, Insert, Update and Delete a document in couchbase database.

## Installation

### Flogo CLI
```bash
flogo install github.com/project-flogo/edge-contrib/activity/coap
```

## Configuration

### Settings:
| Name    | Type   | Description
| :---    | :---   | :---
| method  | string | The CoAP method to use    
| uri     | string | The URI of the service to invoke  
| type    | string | The type of the service    
| options | string | The options to set     

### Input: 

| Name       | Type   | Description
| :---       | :---   | :---
| queryParams| string | The query params of the CoAP Message    
| messageId  | string | The message Id
| payload    | string | The payload of the CoAP Message   
 

### Output:

| Name       | Type   | Description
| :---       | :---   | :---
| response   | string | The response from the service"

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