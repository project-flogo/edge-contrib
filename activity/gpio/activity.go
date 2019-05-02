package gpio

import (
	"errors"
	"fmt"
	"strings"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/metadata"
	"github.com/stianeikeland/go-rpio/v4"
)

// log is the default package logger

const (
	setState  = "Set State"
	readState = "Read State"
	pull      = "Pull"

	in   = "Input"
	high = "High"
	//low = "Low"
	direction = "Direction"

	up   = "Up"
	down = "Down"
	//off = "off"

	//ouput

	result = "result"
)

func init() {
	_ = activity.Register(&GPIOActivity{}, New)
}

type GPIOActivity struct {
	settings *Settings
}

var GPIOActivitymd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

func New(ctx activity.InitContext) (activity.Activity, error) {
	s := &Settings{}

	err := metadata.MapToStruct(ctx.Settings(), s, true)

	if err != nil {
		return nil, err
	}

	return &GPIOActivity{settings: s}, nil
}

func (a *GPIOActivity) Metadata() *activity.Metadata {
	return GPIOActivitymd
}

// Eval implements api.Activity.Eval - Invokes a REST Operation
func (a *GPIOActivity) Eval(ctx activity.Context) (done bool, err error) {

	log := ctx.Logger()
	//getmethod
	log.Debug("Running gpio activity.")

	input := &Input{}

	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}
	if input == nil {
		return true, errors.New("Input Not Set")
	}

	out := &Output{}
	//Open pin
	openErr := rpio.Open()

	defer rpio.Close()

	if openErr != nil {
		log.Errorf("Open RPIO error: %+v", openErr.Error())
		return true, errors.New("Open RPIO error: " + openErr.Error())
	}

	pin := rpio.Pin(a.settings.PinNumber)

	switch a.settings.Method {
	case direction:

		if strings.EqualFold(in, input.Direction) {
			log.Debugf("Set pin %d direction to input", pin)
			pin.Input()
		} else {
			log.Debugf("Set pin %d direction to output", pin)
			pin.Output()
		}
	case setState:

		if strings.EqualFold(high, input.State) {
			log.Debugf("Set pin %d state to High", pin)
			pin.High()
		} else {
			log.Debugf("Set pin %d state to low", pin)
			pin.Low()
		}

	case readState:

		log.Debugf("Read pin %d state..", pin)
		state := pin.Read()
		log.Debugf("Read state and state: %s", state)

		out.Result, err = coerce.ToInt(state)
		if err != nil {
			return false, err
		}
		err = ctx.SetOutputObject(out)

		if err != nil {
			return false, err
		}
		return true, nil
	case pull:

		if strings.EqualFold(up, input.Pull) {
			log.Debugf("Pull pin %d  to Up", pin)
			pin.PullUp()
		} else if strings.EqualFold(down, input.Pull) {
			log.Debugf("Pull pin %d to Down", pin)
			pin.PullDown()
		} else {
			log.Debugf("Pull pin %d to Up", pin)
			pin.PullOff()
		}
	default:
		log.Errorf("Cannot find method %s ", a.settings.Method)
		return true, fmt.Errorf("Cannot find method %s", a.settings.Method)
	}

	out.Result = 0

	err = ctx.SetOutputObject(out)
	return true, nil
}
