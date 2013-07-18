// Copyright © 2013 Dennis T Kaplan <http://www.robotamer.com>

// sbs can encode a struct to a byte slice and back
// This is useful if you need to save data in a key value database such as
// a leveldb, because a leveldb only takes bytes as values.
package sbs

import (
	"bytes"
	"encoding/gob"
)

// Enc takes a populated struct and returns a []byte
func Enc(o interface{}) (b []byte, err error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err = enc.Encode(o)
	if err != nil {
		return
	}
	b = buf.Bytes()
	return
}

// Dec takes an empty struct as well as the []byte returned by Enc(),
// and returns a populated struct.
func Dec(o interface{}, b []byte) (interface{}, error) {
	var buf bytes.Buffer
	_, err := buf.Write(b)
	if err != nil {
		return o, err
	}
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(o)
	return o, err
}
