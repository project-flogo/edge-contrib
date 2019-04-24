package coap

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/dustin/go-coap"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
)

func init() {
	_ = activity.Register(&CoAPActivity{}, New)
}

// CoAPActivity is an Activity that is used to send a CoAP message
// inputs : {method,type,payload,messageId}
// outputs: {result}
type CoAPActivity struct {
	settings *Settings
	req      coap.Message
	coapURL  *url.URL
}

var CoAPActivitymd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

// New creates a new CoAP activity
func New(ctx activity.InitContext) (activity.Activity, error) {
	s := &Settings{}

	err := metadata.MapToStruct(ctx.Settings(), s, true)

	if err != nil {
		return nil, err
	}

	req := coap.Message{
		Type:      toCoapType(s.Type),
		Code:      toCoapCode(s.Code),
		MessageID: uint16(s.MessageId),
	}

	coapURI, err := url.Parse(s.URI)
	if err != nil {
		return nil, err
	}

	scheme := coapURI.Scheme
	if scheme != "coap" {
		return nil, errors.New("URI scheme must be 'coap'")
	}
	for k, v := range s.Options {
		op, val := toOption(k, v)
		req.SetOption(op, val)
	}
	req.SetPathString(coapURI.Path)

	return &CoAPActivity{settings: s, req: req, coapURL: coapURI}, nil
}

func (act *CoAPActivity) Metadata() *activity.Metadata {
	return CoAPActivitymd
}

//todo enhance CoAP client code

// Eval implements api.Activity.Eval - Invokes a REST Operation
func (act *CoAPActivity) Eval(ctx activity.Context) (done bool, err error) {

	log := ctx.Logger()
	input := &Input{}
	out := &Output{}

	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}

	if input.Payload != "" {
		act.req.Payload = []byte(input.Payload)
	}

	if input.QueryParams != nil {

		qp := url.Values{}

		for key, value := range input.QueryParams {
			qp.Set(key, value)
		}

		queryStr := qp.Encode()
		act.req.SetOption(coap.URIQuery, queryStr)
		log.Debugf("CoAP Message: [%s] %s?%s\n", act.settings.Code, act.coapURL.Path, queryStr)

	} else {
		log.Debugf("CoAP Message: [%s] %s\n", act.settings.Code, act.coapURL.Path)
	}

	c, err := coap.Dial("udp", act.coapURL.Host)
	if err != nil {
		return false, activity.NewError(err.Error(), "", nil)
	}

	log.Debugf("conn: %v\n", c)

	rv, err := c.Send(act.req)
	if err != nil {
		return false, err
	}

	if rv != nil {

		if rv.Code > 100 {
			return false, fmt.Errorf("CoAP Error: %s ", rv.Code.String())
		}

		if rv.Payload != nil {
			log.Debugf("Response payload: %s", rv.Payload)

			out.Response = string(rv.Payload)
			ctx.SetOutputObject(out)
		}
	}

	return true, nil
}
