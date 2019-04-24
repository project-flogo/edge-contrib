package coap

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	Port string `md:"port"`
}
type HandlerSettings struct {
	Path string `md:"path,required"`
}

type Output struct {
	QueryParams map[string]string `md:"queryparams"`
	Payload     string            `md:"payload"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	var err error

	o.QueryParams, err = coerce.ToParams(values["queryparams"])
	if err != nil {
		return err
	}

	o.Payload, err = coerce.ToString(values["payload"])
	if err != nil {
		return err
	}

	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"queryparams": o.QueryParams,
		"payload":     o.Payload,
	}
}
