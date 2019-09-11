package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/connection"
)

func init() {

	connection.RegisterManagerFactory(&Factory{})
}

type Settings struct {
	Broker    string `md:"broker"`    //The Broker to connect to
	Id        string `md:"id"`        // Id of the client
	User      string `md:"user"`      // User name of the client
	Password  string `md:"password"`  //Password of the client
	Store     string `md:"store"`     //Cert Store
	Cleansess bool   `md:"cleansess"` //Cleansess flag
	Close     uint   `md:"close"`     //Time in millisecond to disconnect
}

type MqttSharedConn struct {
	settings *Settings
	conn     mqtt.Client
}

type Factory struct {
}

func (f *Factory) Type() string {

	return "mqtt:paho.mqtt.golang"
}

func (*Factory) NewManager(settings map[string]interface{}) (connection.Manager, error) {
	settingStruct := &Settings{}
	err := metadata.MapToStruct(settings, settingStruct, true)
	if err != nil {
		return nil, err
	}
	conn, err := getMqttConnection(settingStruct)
	if err != nil {
		//fmt.Printf("Mqtt Client initialization got error: [%s]", err.Error())
		return nil, err
	}
	if token := conn.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	sharedConn := &MqttSharedConn{settings: settingStruct, conn: conn}

	return sharedConn, nil
}

func (h *MqttSharedConn) Type() string {

	return "mqtt:paho.mqtt.golang"
}

func (h *MqttSharedConn) GetConnection() interface{} {

	return h.conn
}

func (h *MqttSharedConn) ReleaseConnection(connection interface{}) {

	h.conn.Disconnect(h.settings.Close)

}

func (h *MqttSharedConn) Start() error {

	return nil
}

func getMqttConnection(settings *Settings) (mqtt.Client, error) {
	options := initClientOption(settings)

	mqttClient := mqtt.NewClient(options)

	return mqttClient, nil
}

func initClientOption(settings *Settings) *mqtt.ClientOptions {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(settings.Broker)
	opts.SetClientID(settings.Id)
	opts.SetUsername(settings.User)
	opts.SetPassword(settings.Password)
	b, err := coerce.ToBool(settings.Cleansess)
	if err != nil {
		//log.Error("Error converting \"cleansess\" to a boolean ", err.Error())
		return nil
	}
	opts.SetCleanSession(b)
	if storeType := settings.Store; storeType != ":memory:" {
		if settings.Store != "" {
			opts.SetStore(mqtt.NewFileStore(settings.Store))
		}

	}
	return opts
}
