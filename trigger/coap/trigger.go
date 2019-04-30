package coap

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"net/url"
	"strings"

	"github.com/dustin/go-coap"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/trigger"
)

var triggerMd = trigger.NewMetadata(&Settings{}, &HandlerSettings{}, &Output{})

func init() {
	_ = trigger.Register(&CoApTrigger{}, &Factory{})
}

type Factory struct {
}

// Metadata implements trigger.Factory.Metadata
func (*Factory) Metadata() *trigger.Metadata {
	return triggerMd
}

// New implements trigger.Factory.New
func (*Factory) New(config *trigger.Config) (trigger.Trigger, error) {
	s := &Settings{}
	err := metadata.MapToStruct(config.Settings, s, true)
	if err != nil {
		return nil, err
	}

	return &CoApTrigger{id: config.Id, settings: s}, nil
}

// Trigger CoAp trigger struct
type CoApTrigger struct {
	server    *Server
	settings  *Settings
	resources map[string]*CoapResource
	id        string
	logger    log.Logger
}
type CoapResource struct {
	path     string
	attrs    map[string]string
	handlers map[coap.COAPCode]trigger.Handler
}

func (t *CoApTrigger) Initialize(ctx trigger.InitContext) error {

	t.logger = ctx.Logger()

	mux := coap.NewServeMux()

	mux.Handle("/.well-known/core", coap.FuncHandler(t.handleDiscovery))

	// Init handlers
	for _, handler := range ctx.GetHandlers() {

		s := &HandlerSettings{}
		err := metadata.MapToStruct(handler.Settings(), s, true)
		if err != nil {
			return err
		}

		path := s.Path
		method := s.Method

		t.logger.Debugf("Registering handler [%s]", path)
		resource, exists := t.resources[path]

		if !exists {
			resource = &CoapResource{path: path, attrs: make(map[string]string), handlers: make(map[coap.COAPCode]trigger.Handler)}
			t.resources[path] = resource
		}

		resource.handlers[toCoapCode(method)] = handler

		mux.Handle(path, newActionHandler(t, resource))
	}
	t.server = NewServer("udp", t.settings.Port, mux)

	return nil
}

func (t *CoApTrigger) Start() error {
	return t.server.Start()
}

// Stop implements util.Managed.Stop
func (t *CoApTrigger) Stop() error {
	return t.server.Stop()
}

func newActionHandler(t *CoApTrigger, resource *CoapResource) coap.Handler {

	return coap.FuncHandler(func(conn *net.UDPConn, addr *net.UDPAddr, msg *coap.Message) *coap.Message {
		t.logger.Debugf("CoAP Trigger: Recieved request")

		uriQuery := msg.Option(coap.URIQuery)
		var data map[string]interface{}

		if uriQuery != nil {
			//todo handle error
			queryValues, _ := url.ParseQuery(uriQuery.(string))

			queryParams := make(map[string]string, len(queryValues))

			for key, value := range queryValues {
				queryParams[key] = strings.Join(value, ",")
			}

			data = map[string]interface{}{
				"queryParams": queryParams,
				"payload":     string(msg.Payload),
			}
		} else {
			data = map[string]interface{}{
				"payload": string(msg.Payload),
			}
		}
		handler, exists := resource.handlers[msg.Code]

		if !exists {
			res := &coap.Message{
				Type:      coap.Reset,
				Code:      coap.MethodNotAllowed,
				MessageID: msg.MessageID,
				Token:     msg.Token,
			}

			return res
		}

		_, err := handler.Handle(context.Background(), data)

		if err != nil {
			//todo determining if 404 or 500
			res := &coap.Message{
				Type:      coap.Reset,
				Code:      coap.NotFound,
				MessageID: msg.MessageID,
				Token:     msg.Token,
				Payload:   []byte(fmt.Sprintf("Unable to execute handler '%s'", handler)),
			}

			return res
		}

		t.logger.Debugf("Ran: %s", handler)

		if msg.IsConfirmable() {
			res := &coap.Message{
				Type:      coap.Acknowledgement,
				Code:      0,
				MessageID: msg.MessageID,
				Token:     msg.Token,
			}
			//res.SetOption(coap.ContentFormat, coap.TextPlain)

			t.logger.Debugf("Transmitting %#v", res)
			return res
		}

		return nil

	})
}

func (t *CoApTrigger) handleDiscovery(conn *net.UDPConn, addr *net.UDPAddr, msg *coap.Message) *coap.Message {

	//todo add filter support
	if msg.Code == coap.GET {
		var buffer bytes.Buffer

		numResources := len(t.resources)

		i := 0
		for _, resource := range t.resources {

			i++

			buffer.WriteString("<")
			buffer.WriteString(resource.path)
			buffer.WriteString(">")

			if len(resource.attrs) > 0 {
				for k, v := range resource.attrs {
					buffer.WriteString(";")
					buffer.WriteString(k)
					buffer.WriteString("=")
					buffer.WriteString(v)
				}
			}

			if i < numResources {
				buffer.WriteString(",\n")
			} else {
				buffer.WriteString("\n")
			}
		}

		payloadStr := buffer.String()

		res := &coap.Message{
			Type:      msg.Type,
			Code:      coap.Content,
			MessageID: msg.MessageID,
			Token:     msg.Token,
			Payload:   []byte(payloadStr),
		}
		res.SetOption(coap.ContentFormat, coap.AppLinkFormat)

		t.logger.Debugf("Transmitting %#v", res)

		return res
	}
	return nil

}
