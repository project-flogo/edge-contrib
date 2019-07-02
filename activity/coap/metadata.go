package coap

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	Method      string            `md:"method,required"` // The CoAP method to use
	URI         string            `md:"uri,required"`    // The CoAP resource URI
	MessageType string            `md:"messageType"`     // The message type
	Options     map[string]string `md:"options"`         // The CoAP options to set
}

type Input struct {
	QueryParams map[string]string `md:"queryParams"` // The query params of the CoAP message
	MessageId   int               `md:"messageId"`   // ID used to detect duplicates and for optional reliability
	Payload     string            `md:"payload"`     // The payload of the CoAP message
}

type Output struct {
	Response string `md:"response"` // The response
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"queryParams": i.QueryParams,
		"payload":     i.Payload,
		"messageId":   i.MessageId,
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
	i.MessageId, err = coerce.ToInt(values["messageId"])

	return err
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
