package coap

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"

	"github.com/dustin/go-coap"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
)

func init() {
	_ = activity.Register(&CoAPActivity{}, New)
}

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

// New creates a new CoAP activity
func New(ctx activity.InitContext) (activity.Activity, error) {

	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	coapURI, err := url.Parse(s.URI)
	if err != nil {
		return nil, err
	}

	scheme := coapURI.Scheme
	if scheme != "coap" {
		return nil, errors.New("URI scheme must be 'coap'")
	}

	msgType := coap.Confirmable
	if s.MessageType != "" {
		msgType = toCoapMsgType(s.MessageType)
	}

	return &CoAPActivity{settings: s, msgType: msgType, methodCode: toCoapMethodCode(s.Method), coapURI: coapURI}, nil
}

// CoAPActivity is an Activity that is used to send a CoAP message
// inputs : {method,type,payload,messageId}
// outputs: {result}
type CoAPActivity struct {
	settings *Settings

	coapURI    *url.URL
	msgType    coap.COAPType
	methodCode coap.COAPCode
}

func (act *CoAPActivity) Metadata() *activity.Metadata {
	return activityMd
}

//todo enhance CoAP client code

// Eval implements api.Activity.Eval - Invokes a REST Operation
func (act *CoAPActivity) Eval(ctx activity.Context) (done bool, err error) {

	log := ctx.Logger()
	input := &Input{}

	req := coap.Message{
		Type: act.msgType,
		Code: act.methodCode,
	}

	req.SetPathString(act.coapURI.Path)

	for k, v := range act.settings.Options {
		op, val := toOption(k, v)
		req.SetOption(op, val)
	}

	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}

	if input.MessageId == 0 {
		req.MessageID = uint16(rand.Intn(10))
	} else {
		req.MessageID = uint16(input.MessageId)
	}

	if input.Payload != "" {
		req.Payload = []byte(input.Payload)
	}

	if input.QueryParams != nil {

		qp := url.Values{}

		for key, value := range input.QueryParams {
			qp.Set(key, value)
		}

		queryStr := qp.Encode()
		req.SetOption(coap.URIQuery, queryStr)

		log.Debugf("CoAP Message: [%s] %s?%s\n", act.settings.Method, act.coapURI.Path, queryStr)

	} else {
		log.Debugf("CoAP Message: [%s] %s\n", act.settings.Method, act.coapURI.Path)
	}

	c, err := coap.Dial("udp", act.coapURI.Host)
	if err != nil {
		return false, activity.NewError(err.Error(), "", nil)
	}

	log.Debugf("conn: %v\n", c)

	rv, err := c.Send(req)
	if err != nil {
		return false, err
	}

	if rv != nil {

		if rv.Code > 100 {
			return false, fmt.Errorf("CoAP Error: %s ", rv.Code.String())
		}

		if rv.Payload != nil {
			log.Tracef("Response payload: %s", rv.Payload)

			out := &Output{}
			out.Response = string(rv.Payload)
			err := ctx.SetOutputObject(out)
			if err != nil {
				return false, err
			}
		}
	}

	return true, nil
}
