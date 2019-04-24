package mqtt

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	Broker    string `md:"broker"`
	Id        string `md:"id"`
	User      string `md:"user"`
	Password  string `md:"password"`
	Store     string `md:"store"`
	Cleansess bool   `md:"cleansess"`
}

type Input struct {
	Topic   string      `md:"topic"`
	Qos     int         `md:"qos"`
	Message interface{} `md:"message"`
}

type Output struct {
	Data interface{} `md:"data"`
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"message": i.Message,
		"topic":   i.Topic,
		"qos":     i.Qos,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Message, _ = values["message"]

	i.Topic, err = coerce.ToString(values["topic"])
	if err != nil {
		return err
	}

	i.Qos, err = coerce.ToInt(values["qos"])
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
