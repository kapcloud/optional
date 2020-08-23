package optional

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString_Value(t *testing.T) {
	val1 := String{}
	assert.False(t, val1.IsPresent())
	assert.Equal(t, val1.Value(), "")
	assert.Nil(t, val1.value)

	val2 := NewString("")
	assert.True(t, val2.IsPresent())
	assert.Equal(t, val2.Value(), "")
	assert.NotNil(t, val2.value)

	val3 := NewString("foo")
	assert.True(t, val3.IsPresent())
	assert.Equal(t, val3.Value(), "foo")
	assert.NotNil(t, val3.value)
}

func TestString_MarshalJSON(t *testing.T) {
	type fields struct {
		WithValue     String
		WithZeroValue String
		WithNoValue   String
		Unused        String
	}

	val := fields{
		WithValue:     NewString("foo"),
		WithZeroValue: NewString(""),
		WithNoValue:   String{},
	}

	data, err := json.Marshal(val)
	assert.NoError(t, err)

	want := `{"WithValue":"foo","WithZeroValue":"","WithNoValue":null,"Unused":null}`
	assert.Equal(t, string(data), want)
}

func TestString_UnmarshalJSON(t *testing.T) {
	type fields struct {
		WithValue     String
		WithZeroValue String
		WithNoValue   String
		Unused        String
	}

	var jsonString = `{"WithValue":"foo","WithZeroValue":"","WithNoValue":null}`

	val := fields{}
	err := json.Unmarshal([]byte(jsonString), &val)
	assert.NoError(t, err)

	assert.True(t, val.WithValue.IsPresent())
	assert.Equal(t, val.WithValue.Value(), "foo")
	assert.NotNil(t, val.WithValue.value)

	assert.True(t, val.WithZeroValue.IsPresent())
	assert.Equal(t, val.WithZeroValue.Value(), "")
	assert.NotNil(t, val.WithZeroValue.value)

	assert.False(t, val.WithNoValue.IsPresent())
	assert.Equal(t, val.WithNoValue.Value(), "")
	assert.Nil(t, val.WithNoValue.value)

	assert.False(t, val.Unused.IsPresent())
	assert.Equal(t, val.Unused.Value(), "")
	assert.Nil(t, val.Unused.value)
}

func TestString_UnmarshalJSON_Overwritten(t *testing.T) {
	type fields struct {
		WithValue     String
		WithZeroValue String
		WithNoValue   String
		Unused        String
	}

	val := fields{
		WithValue:     NewString("seed_a"),
		WithZeroValue: NewString("seed_b"),
		WithNoValue:   NewString("seed_c"),
		Unused:        NewString("prev_value"),
	}

	var jsonString = `{"WithValue":"foo","WithZeroValue":"","WithNoValue":null}`
	err := json.Unmarshal([]byte(jsonString), &val)
	assert.NoError(t, err)

	assert.True(t, val.WithValue.IsPresent())
	assert.Equal(t, val.WithValue.Value(), "foo")
	assert.NotNil(t, val.WithValue.value)

	assert.True(t, val.WithZeroValue.IsPresent())
	assert.Equal(t, val.WithZeroValue.Value(), "")
	assert.NotNil(t, val.WithZeroValue.value)

	assert.False(t, val.WithNoValue.IsPresent())
	assert.Equal(t, val.WithNoValue.Value(), "")
	assert.Nil(t, val.WithNoValue.value)

	assert.True(t, val.Unused.IsPresent())
	assert.Equal(t, val.Unused.Value(), "prev_value")
	assert.NotNil(t, val.Unused.value)
}
