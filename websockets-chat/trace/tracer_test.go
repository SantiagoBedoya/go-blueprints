package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buff bytes.Buffer
	tracer := New(&buff)
	if tracer == nil {
		t.Error("Return form New should not be nil")
	} else {
		tracer.Trace("Hello trace package.")
		if buff.String() != "Hello trace package.\n" {
			t.Errorf("Trace should not write '%s'.", buff.String())
		}
	}
}

func TestOff(t *testing.T) {
	var silentTracer Tracer = Off()
	silentTracer.Trace("something")
}
