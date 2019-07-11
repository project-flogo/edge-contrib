package mqtt

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	Broker        string                 `md:"broker,required"` // The broker URL
	Id            string                 `md:"id,required"`     // The id of client
	Username      string                 `md:"username"`        // The user's name
	Password      string                 `md:"password"`        // The user's password
	Store         string                 `md:"store"`           // The store for message persistence
	CleanSession  bool                   `md:"cleanSession"`    // Clean session flag
	KeepAlive     int                    `md:"keepAlive"`       // Keep Alive time in seconds
	AutoReconnect bool                   `md:"autoReconnect"`   // Enable Auto-Reconnect
	SSLConfig     map[string]interface{} `md:"sslConfig"`       // SSL Configuration
}

type HandlerSettings struct {
	Topic      string `md:"topic,required"` // The topic to listen on
	ReplyTopic string `md:"replyTopic"`     // The topic to reply on
	Qos        int    `md:"qos"`            // The Quality of Service
	Retain     bool   `md:"retain"`         // Retain messages
}

type Output struct {
	Message     string            `md:"message"`     // The message recieved
	Topic       string            `md:"topic"`       // The MQTT topic
	TopicParams map[string]string `md:"topicParams"` // The topic parameters
}

type Reply struct {
	Data interface{} `md:"data"` // The data to reply with
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"message":     o.Message,
		"topic":       o.Topic,
		"topicParams": o.TopicParams,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {

	var err error
	o.Message, err = coerce.ToString(values["message"])
	if err != nil {
		return err
	}
	o.Topic, err = coerce.ToString(values["topic"])
	if err != nil {
		return err
	}
	o.TopicParams, err = coerce.ToParams(values["topicParams"])
	if err != nil {
		return err
	}

	return nil
}

func (r *Reply) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"data": r.Data,
	}
}

func (r *Reply) FromMap(values map[string]interface{}) error {

	r.Data = values["data"]
	return nil
}
