package mqtt

import (
	"encoding/json"
	"net"
	"os"
	"os/exec"
	"testing"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"github.com/project-flogo/core/action"
	"github.com/project-flogo/core/support"
	"github.com/project-flogo/core/support/test"
	"github.com/project-flogo/core/trigger"
	"github.com/stretchr/testify/assert"
)

func Pour(port string) {
	for {
		conn, _ := net.Dial("tcp", net.JoinHostPort("", port))
		if conn != nil {
			conn.Close()
			break
		}
	}
}

func TestMain(m *testing.M) {
	command := exec.Command("docker", "start", "mqtt")
	err := command.Run()
	if err != nil {
		command := exec.Command("docker", "run", "-p", "1883:1883", "-p", "9001:9001", "--name", "mqtt", "-d", "eclipse-mosquitto")
		err := command.Run()
		if err != nil {
			panic(err)
		}
	}
	Pour("1883")
	os.Exit(m.Run())
}

const testConfig string = `{
	"id": "trigger-mqtt",
	"ref": "github.com/project-flogo/device-contrib/trigger/mqtt",
	"settings": {
		"autoreconnect": true,
        "broker": "ssl://mqtt.bosch-iot-hub.com:8883",
        "certstore": "*****8",
        "cleansess": false,
        "enableTLS": true,
		"id": "FlogoTest237123",
        "keepalive": 30,
        "password": "****",
        "store": ":memory:",
        "user": "*****"
    },
	"handlers": [
	  {
		"settings": {
			"topic": "control/+/+/req/#"
		},
		"action" : {
		  "id": "dummy"
		}
	  }
	]
  }`

const testConfigLocal string = `{
		"id": "trigger-mqtt",
		"ref": "github.com/project-flogo/device-contrib/trigger/mqtt",
		"settings": {
			"autoreconnect": true,
	        "broker": "tcp://localhost:1883",
	        "cleansess": false,
	        "enableTLS": false,
					"id": "FlogoTest237123",
	        "keepalive": 30,
	        "store": ":memory:"
	    },
		"handlers": [
		  {
			"settings": {
				"topic": "control/+/+/req/#"
			},
			"action" : {
			  "id": "dummy"
			}
		  }
		]
	  }`

const testConfigLocalgetHandler string = `{
				"id": "trigger1-mqtt",
				"ref": "github.com/project-flogo/device-contrib/trigger/mqtt",
				"settings": {
					"autoreconnect": true,
			        "broker": "tcp://localhost:1883",
			        "cleansess": false,
			        "enableTLS": false,
							"id": "FlogoTest2371231",
			        "keepalive": 30,
			        "store": ":memory:"
			    },
				"handlers": [
				  {
					"settings": {
						"topic": "test/+/+/req/#"
					},
					"action" : {
					  "id": "dummyTest"
					}
				  }
				]
			  }`

func TestParseTopic(t *testing.T) {
	test := func(input, output string) {
		parsed := ParseTopic(input)
		assert.Equal(t, output, parsed.String())
	}
	test("/a/+x/b/#y", "/a/+/b/#")
	test("/a/+/b/#", "/a/+/b/#")
	test("a/+/b/#", "a/+/b/#")
	test("a/+/b", "a/+/b")
	test("a/+/b/", "a/+/b/")
	test("", "")
	test("+", "+")
	test("#", "#")
	test("/", "/")
	test("/+", "/+")
	test("/#", "/#")
}

func TestTopic_Match(t *testing.T) {
	test := func(match, in string, params map[string]string) {
		parsed, input := ParseTopic(match), ParseTopic(in)
		found := parsed.Match(input)
		for key, value := range params {
			assert.Equal(t, found[key], value)
		}
	}
	test("/a/+x/b/#y", "/a/x/b/y/z/", map[string]string{"x": "x", "y": "y/z/"})
	test("a/+x/b/#y", "a/x/b/y/z/", map[string]string{"x": "x", "y": "y/z/"})
	test("/a/+x/b/#y", "/a/x/b/y/z", map[string]string{"x": "x", "y": "y/z"})
	test("/a/+x/b/#y", "/a/x/b/y", map[string]string{"x": "x", "y": "y"})
	test("/a/+x/b/#y", "/a/x/b", map[string]string{"x": "x", "y": ""})
	test("/a/+x/b/", "/a/x/b/", map[string]string{"x": "x"})
	test("/a/+x/b", "/a/x/b", map[string]string{"x": "x"})
	test("/a/+x", "/a/x", map[string]string{"x": "x"})
	test("/+x", "/x/gah", map[string]string{"x": "x"})
	test("/#x", "/x/gah", map[string]string{"x": "x/gah"})
	test("+x", "x/gah", map[string]string{"x": "x"})
	test("#x", "x/gah", map[string]string{"x": "x/gah"})
}

func TestTrigger_Register(t *testing.T) {

	ref := support.GetRef(&Trigger{})
	f := trigger.GetFactory(ref)
	assert.NotNil(t, f)
}

func TestRestTrigger_Initialize(t *testing.T) {

	f := &Factory{}

	config := &trigger.Config{}
	err := json.Unmarshal([]byte(testConfigLocal), config)
	assert.Nil(t, err)

	actions := map[string]action.Action{"dummy": test.NewDummyAction(func() {
	})}

	trg, err := test.InitTrigger(f, config, actions)
	assert.Nil(t, err)
	assert.NotNil(t, trg)

	err = trg.Start()
	assert.Nil(t, err)

	err = trg.Stop()
	assert.Nil(t, err)
}

func TestRestTrigger_getHanlder(t *testing.T) {

	f := &Factory{}

	config := &trigger.Config{}
	err := json.Unmarshal([]byte(testConfigLocalgetHandler), config)
	assert.Nil(t, err)

	done := make(chan bool, 1)
	actions := map[string]action.Action{"dummyTest": test.NewDummyAction(func() {
		done <- true
	})}

	trg, err := test.InitTrigger(f, config, actions)
	assert.Nil(t, err)
	assert.NotNil(t, trg)

	err = trg.Start()
	assert.Nil(t, err)

	options := mqtt.NewClientOptions()
	options.AddBroker("tcp://localhost:1883")
	options.SetClientID("TestAbc123")
	client := mqtt.NewClient(options)
	token := client.Connect()
	token.Wait()
	assert.Nil(t, token.Error())

	token = client.Publish("test/a/b/req/c/d", 0, true, []byte(`{"message": "hello world"}`))
	token.Wait()
	assert.Nil(t, token.Error())
	select {
	case <-done:
	case <-time.Tick(time.Second):
		t.Fatal("didn't get message in time")
	}
	client.Disconnect(50)

	err = trg.Stop()
	assert.Nil(t, err)
}
