package sbs

import (
	"testing"
)

type Foo struct {
	A int
	B string
}

func TestEnc(t *testing.T) {
	p := &Foo{111, "A string"}

	bytesdata, err := Enc(p)
	if err != nil {
		t.Errorf("Enc Error: %s", err)
	}

	foo := new(Foo)
	_, err = Dec(foo, bytesdata)
	if err != nil {
		t.Errorf("Dec Error: %s", err)
	}
}
