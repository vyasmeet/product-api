package data

import (
	"encoding/json"
	"io"
)

//	ToJSON serialise the given interface into a string JSON format
func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}

//	FromJSON deserialise the object from JSON string into an io.Reader to given interface
func FromJSON(i interface{}, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(i)
}
