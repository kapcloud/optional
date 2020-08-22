package optional

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt_Value(t *testing.T) {
	v1 := Int{}
	assert.False(t, v1.IsPresent())
	assert.Equal(t, v1.Value(), 0)
	assert.Nil(t, v1.value)

	v2 := NewInt(0)
	assert.True(t, v2.IsPresent())
	assert.Equal(t, v2.Value(), 0)
	assert.NotNil(t, v2.value)

	v3 := NewInt(1)
	assert.True(t, v3.IsPresent())
	assert.Equal(t, v3.Value(), 1)
	assert.NotNil(t, v3.value)
}

func TestInt_MarshalJSON(t *testing.T) {
	type fields struct {
		WithValue     Int
		WithZeroValue Int
		WithNoValue   Int
		Unused        Int
	}

	val := fields{
		WithValue:     NewInt(37),
		WithZeroValue: NewInt(0),
		WithNoValue:   Int{},
	}

	data, err := json.Marshal(val)
	assert.NoError(t, err)

	want := `{"WithValue":37,"WithZeroValue":0,"WithNoValue":null,"Unused":null}`
	assert.Equal(t, string(data), want)
}

func TestInt_UnmarshalJSON(t *testing.T) {
	type fields struct {
		WithValue     Int
		WithZeroValue Int
		WithNoValue   Int
		Unused        Int
	}

	var jsonString = `{"WithValue":37,"WithZeroValue":0,"WithNoValue":null}`

	val := fields{}
	err := json.Unmarshal([]byte(jsonString), &val)
	assert.NoError(t, err)

	assert.True(t, val.WithValue.IsPresent())
	assert.Equal(t, val.WithValue.Value(), 37)
	assert.NotNil(t, val.WithValue.value)

	assert.True(t, val.WithZeroValue.IsPresent())
	assert.Equal(t, val.WithZeroValue.Value(), 0)
	assert.NotNil(t, val.WithZeroValue.value)

	assert.False(t, val.WithNoValue.IsPresent())
	assert.Equal(t, val.WithNoValue.Value(), 0)
	assert.Nil(t, val.WithNoValue.value)

	assert.False(t, val.Unused.IsPresent())
	assert.Equal(t, val.Unused.Value(), 0)
	assert.Nil(t, val.Unused.value)
}

func TestInt_UnmarshalJSON_Overwritten(t *testing.T) {
	type fields struct {
		WithValue     Int
		WithZeroValue Int
		WithNoValue   Int
		Unused        Int
	}

	val := fields{
		WithValue:     NewInt(1),
		WithZeroValue: NewInt(2),
		WithNoValue:   NewInt(3),
		Unused:        NewInt(4),
	}

	var jsonString = `{"WithValue":37,"WithZeroValue":0,"WithNoValue":null}`

	err := json.Unmarshal([]byte(jsonString), &val)
	assert.NoError(t, err)

	assert.True(t, val.WithValue.IsPresent())
	assert.Equal(t, val.WithValue.Value(), 37)
	assert.NotNil(t, val.WithValue.value)

	assert.True(t, val.WithZeroValue.IsPresent())
	assert.Equal(t, val.WithZeroValue.Value(), 0)
	assert.NotNil(t, val.WithZeroValue.value)

	assert.False(t, val.WithNoValue.IsPresent())
	assert.Equal(t, val.WithNoValue.Value(), 0)
	assert.Nil(t, val.WithNoValue.value)

	assert.True(t, val.Unused.IsPresent())
	assert.Equal(t, val.Unused.Value(), 4)
	assert.NotNil(t, val.Unused.value)
}
