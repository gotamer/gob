// structbytes can encode a struct to a byte slice and back
// This is usful if you need to save data in a key value database such as
// a leveldb, because a leveldb only takes bytes as values.
package sbs

import (
	"bytes"
	"encoding/gob"
)

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
