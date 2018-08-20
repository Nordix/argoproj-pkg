package json

import (
	"bytes"
	"encoding/json"
)

// DisallowUnknownFields configures the JSON decoder to error out if unknown
// fields come along, instead of dropping them by default.
func DisallowUnknownFields(d *json.Decoder) *json.Decoder {
	d.DisallowUnknownFields()
	return d
}

// JSONOpt is a decoding option for decoding from JSON format.
type JSONOpt func(*json.Decoder) *json.Decoder

// Unmarshal is a convenience wrapper around json.Unmarshal to support json decode options
func Unmarshal(j []byte, o interface{}, opts ...JSONOpt) error {
	d := json.NewDecoder(bytes.NewReader(j))
	for _, opt := range opts {
		d = opt(d)
	}
	return d.Decode(&o)
}
