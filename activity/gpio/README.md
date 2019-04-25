# GPIO
This activity allows you to control GPIO pins on a Raspberry Pi.

## Installation

### Flogo CLI
```bash
flogo install github.com/project-flogo/edge-contrib/activity/gpio
```

## Configuration

### Settings:
| Name      | Type   | Description
| :---      | :---   | :---
| method    | string | 	The method to take action for specified pin (Allowed values are Direction, Set State, Read State, and Pull)   
| pinNumber | string | The pin number of the GPIO 
 
### Input: 

| Name       | Type   | Description
| :---       | :---   | :---
| direction| string | Set the direction of the pin (Allowed values are Input and Output)   
| state  | string | Set the state of the pin (Allowed values are High and Low)
| pull    | string | Pull the pin to the specified value (Allowed values are Up, Down, and Off)   
 

### Output:

| Name       | Type   | Description
| :---       | :---   | :---
| result   | string | The result of the operation

## Example

```json
{
  "id": "gpio-activity",
  "name": "GPIO Activity",
  "description": "GPIO Example",
  "activity": {
    "ref": "github.com/project-flogo/edge-contrib/activity/gpio",
    "settings": {
      "method" : "Pull",
      "uri": "coap://localhost:5683/flogo"
    },
    "input" : {
        "direction" : "Input",
        "State": "High"
    }
  }
}
```