package coap

import (
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
	server   *Server
	settings *Settings
	id       string
	logger   log.Logger
}

func (t *CoApTrigger) Initialize(ctx trigger.InitContext) error {

	t.logger = ctx.Logger()

	mux := coap.NewServeMux()
	//TODO .. is it neccessary?
	//mux.Handle("/.well-known/core", coap.FuncHandler(t.handleDiscovery))

	// Init handlers
	for _, handler := range ctx.GetHandlers() {

		s := &HandlerSettings{}
		err := metadata.MapToStruct(handler.Settings(), s, true)
		if err != nil {
			return err
		}

		path := s.Path

		t.logger.Debugf("Registering handler [%s]", path)

		mux.Handle(path, newActionHandler(t, handler))
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

func newActionHandler(t *CoApTrigger, handler trigger.Handler) coap.Handler {

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
