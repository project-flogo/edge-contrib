package gpio

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	Method    string `md:"method,required,allowed(Direction,Set State,Read State, Pull)"`
	PinNumber int    `md:"pinNumber,required"`
}

type Input struct {
	Direction string `md:"direction,allowed(Input,Output)"`
	State     string `md:"state,allowed(High, Low)"`
	Pull      string `md:"pull,allowed(Up,Down,Off)"`
}

type Output struct {
	Result int `md:"result"`
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"direction": i.Direction,
		"state":     i.State,
		"pull":      i.Pull,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {
	var err error

	i.Direction, err = coerce.ToString(values["direction"])
	if err != nil {
		return err
	}
	i.State, err = coerce.ToString(values["state"])
	if err != nil {
		return err
	}
	i.Pull, err = coerce.ToString(values["pull"])
	if err != nil {
		return err
	}
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"result": o.Result,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {
	var err error

	o.Result, err = coerce.ToInt(values["direction"])
	if err != nil {
		return err
	}
	return nil
}
