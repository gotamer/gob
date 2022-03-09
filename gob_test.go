package gob

import (
	"testing"
)

type T struct {
	A int64
	B string
	C float64
}

func TestEncDec(t *testing.T) {
	o := new(T)
	o.A = 987654321
	o.B = "sbs"
	o.C = 987.789

	var err error
	var bs []byte

	if bs, err = Marshal(o); err != nil {
		t.Errorf("Enc Error: %v", err.Error())
	}

	oo := new(T)
	if err = Unmarshal(bs, oo); err != nil {
		t.Errorf("Dec Error: %v", err.Error())
	}

	if oo.A != o.A {
		t.Errorf("Dec %v should be %v", oo.A, o.A)
	}

	if oo.B != o.B {
		t.Errorf("Dec %v should be %v", oo.B, o.B)
	}

	if oo.C != o.C {
		t.Errorf("Dec %v should be %v", oo.C, o.C)
	}
}
