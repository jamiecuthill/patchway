package patchway

import "encoding/json"

type Operation interface {
	String() string
	Bytes() []byte
}

// Operations are a collection of Operation
// that can be applied to a resource.
// It is used to represent a patch document.
type Operations []Operation

func (o Operations) Bytes() []byte {
	b, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	return b
}
