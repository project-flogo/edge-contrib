package mqtt

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	Broker       string                 `md:"broker,required"` // The broker URL
	Id           string                 `md:"id,required"`     // The id of client
	Username     string                 `md:"username"`        // The user's name
	Password     string                 `md:"password"`        // The user's password
	Store        string                 `md:"store"`           // The store for message persistence
	CleanSession bool                   `md:"cleanSession"`    // Clean session flag

	Retain       bool                   `md:"retain"`          // Retain Messages
  Topic        string                 `md:"topic,required"`  // The topic to publish to
	Qos          int                    `md:"qos"`             // The Quality of Service
	SSLConfig    map[string]interface{} `md:"sslConfig"`       // SSL Configuration
}

type Input struct {
	Message     interface{}       `md:"message"`     // The message to send
	TopicParams map[string]string `md:"topicParams"` // The topic parameters
}

type Output struct {
	Data interface{} `md:"data"` // The data recieved
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"message":     i.Message,
		"topicParams": i.TopicParams,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {
	var err error
	i.Message, _ = values["message"]
	i.TopicParams, err = coerce.ToParams(values["topicParams"])
	if err != nil {
		return err
	}
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"data": o.Data,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {

	o.Data = values["data"]
	return nil
}
