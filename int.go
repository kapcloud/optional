// This file was generated by cmd
package optional

import "encoding/json"

// Int represents an int that may be optional
type Int struct {
	value *int
}

// NewInt creates an optional.Int
func NewInt(val int) Int {
	return Int{
		value: &val,
	}
}

// SetValue sets the int value
func (opt *Int) SetValue(val int) {
	opt.value = &val
}

func (opt Int) initValue() int {
	var val int
	return val
}

// Value returns the int value or init value if not present
func (opt Int) Value() int {
	if opt.value != nil {
		return *opt.value
	} else {
		return opt.initValue()
	}
}

// IsPresent returns whether or not the value is present
func (opt Int) IsPresent() bool {
	return opt.value != nil
}

// MarshalJSON implements the json.MarshalJSON interface
func (opt Int) MarshalJSON() ([]byte, error) {
	if opt.value != nil {
		return json.Marshal(opt.value)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON implements the json.UnmarshalJSON interface
func (opt *Int) UnmarshalJSON(data []byte) error {
	opt.value = nil

	if data == nil {
		return nil
	}

	if string(data) == "null" {
		return nil
	}

	var val int
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}

	opt.value = &val
	return nil
}
