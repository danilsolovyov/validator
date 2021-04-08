package validator

import (
	"encoding/json"
)

type Decoder struct {
	*json.Decoder
}

func (dec *Decoder) DecodeAndValidate(s interface{}, v Validator) error {
	err := dec.Decoder.Decode(s)
	if err != nil {
		return err
	}
	v.AddValues(s)
	err = v.Validate()
	return err
}
