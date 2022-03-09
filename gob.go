//Package gob implements Marshal and Unmarshal using the standard GOB package
package gob

import (
	"bytes"
	"encoding/gob"
)

// Marshal takes a populated struct and returns a []byte in GOB format
func Marshal(o interface{}) (b []byte, err error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err = enc.Encode(o)
	if err != nil {
		return
	}
	b = buf.Bytes()
	return
}

// Unmarshal takes an empty struct as well as the []byte returned by Marshal(),
// and populates the given empty struct from the []byte.
func Unmarshal(b []byte, o interface{}) (err error) {
	var buf bytes.Buffer
	_, err = buf.Write(b)
	if err != nil {
		return
	}
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(o)
	return
}
