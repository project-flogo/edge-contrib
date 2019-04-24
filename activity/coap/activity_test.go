package coap

import (
	"testing"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&CoAPActivity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestSimple(t *testing.T) {
	settings := &Settings{Code: "GET", URI: "coap://localhost:5683/flogo"}

	iCtx := test.NewActivityInitContext(settings, nil)
	act, err := New(iCtx)
	assert.Nil(t, err)

	tc := test.NewActivityContext(act.Metadata())

	act.Eval(tc)
}
