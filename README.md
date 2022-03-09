GoTamer gob v2
==============

#### Example
```go
type Foo struct {
	A int
	B string
}

p := &Foo{111,"A string"}

byteslice, err := gob.Marshal(p)
...

foo := new(Foo)
err := gob.Unmarshal(byteslice, foo)
...
```

