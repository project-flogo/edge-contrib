package mqtt

import (
	"encoding/json"
	"testing"

	"github.com/project-flogo/core/action"
	"github.com/project-flogo/core/support"
	"github.com/project-flogo/core/support/test"
	"github.com/project-flogo/core/trigger"
	"github.com/stretchr/testify/assert"
)

const testConfig string = `{
	"id": "trigger-mqtt",
	"ref": "github.com/project-flogo/device-contrib/trigger/mqtt",
	"settings": {
		"broker": "tcp://localhost:1883",
		"qos": "0",
		"cleansess": "false",
		"id":"client-1"
    },
	"handlers": [
	  {
		"settings": {
		  "topic": "flogo"
		},
		"action" : {
		  "id": "dummy"
		}
	  }
	]
  }`

func TestTrigger_Register(t *testing.T) {

	ref := support.GetRef(&MqttTrigger{})
	f := trigger.GetFactory(ref)
	assert.NotNil(t, f)
}

func TestRestTrigger_Initialize(t *testing.T) {

	f := &Factory{}

	config := &trigger.Config{}
	err := json.Unmarshal([]byte(testConfig), config)
	assert.Nil(t, err)

	actions := map[string]action.Action{"dummy": test.NewDummyAction(func() {
	})}

	trg, err := test.InitTrigger(f, config, actions)

	assert.Nil(t, err)
	assert.NotNil(t, trg)
	err = trg.Start()
	for {

	}

}
