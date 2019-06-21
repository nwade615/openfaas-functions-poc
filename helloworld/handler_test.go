package function

import (
	"testing"
)

func TestHandle(t *testing.T) {
	b := []byte("pants")
	msg := Handle(b)
	if msg != "Hello, OpenFaaS Cloud. You said: pants" {
		t.Errorf("Msg was incorrect")
	}
}
