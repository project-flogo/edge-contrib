package coap

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	Code      string            `md:"code,required"`
	URI       string            `md:"uri,required"`
	Type      string            `md:"type"`
	MessageId int               `md:"messageId"`
	Options   map[string]string `md:"options"`
}

type Input struct {
	QueryParams map[string]string `md:"queryParams"`
	Payload     string            `md:"payload"`
}

type Output struct {
	Response string `md:"response"`
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"queryParams": i.QueryParams,
		"payload":     i.Payload,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {
	var err error

	i.QueryParams, err = coerce.ToParams(values["queryParams"])
	if err != nil {
		return err
	}

	i.Payload, err = coerce.ToString(values["payload"])
	if err != nil {
		return err
	}
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"response": o.Response,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {
	var err error

	o.Response, err = coerce.ToString(values["response"])
	if err != nil {
		return err
	}
	return nil
}
