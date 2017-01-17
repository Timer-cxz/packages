package errors

import (
	"testing"
)

func TestTrace(t *testing.T) {
	err := New("try this errors package")
	err = Trace(err)
	t.Logf(Stack(err).String())
}

func TestEqual(t *testing.T) {
	err1 := Errorf("test equal")
	err2 := Trace(New("test equal"))

	if NotEqual(err1, err2) {
		t.Errorf("error:\"%s\" should equal error:\"%s\"", err1.Error(), err2.Error())
	}
}
