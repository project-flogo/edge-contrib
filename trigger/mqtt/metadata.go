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

type HandlerSettings struct {
	Topic string `md:"topic"`
	Qos   int    `md:"qos"`
}

type Output struct {
	Message string `md:"message"`
}

type Reply struct {
	Data interface{} `md:"data"`
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"message": o.Message,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {

	var err error
	o.Message, err = coerce.ToString(values["message"])
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
