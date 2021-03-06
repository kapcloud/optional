// tool that generates 'optional' type wrappers base type
package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"reflect"
	"strings"
)

// TplParams the template params
type TplParams struct {
	TypeName string
	WarpName string
}

func genKind(kind reflect.Kind) error {
	log.Println("generate", kind)

	typeName := kind.String()
	warpName := strings.Title(typeName)
	fileName := "../" + typeName + ".go"

	params := &TplParams{
		TypeName: typeName,
		WarpName: warpName,
	}

	tpl, err := template.New("").Parse(tmpl)
	if err != nil {
		return err
	}

	buf := bytes.Buffer{}
	if err := tpl.Execute(&buf, params); err != nil {
		return err
	}

	err = ioutil.WriteFile(fileName, buf.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	kinds := []reflect.Kind{
		reflect.Bool,
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Float32,
		reflect.Float64,
		reflect.String,
	}

	for _, kind := range kinds {
		if err := genKind(kind); err != nil {
			panic(err)
		}
	}
}

var tmpl = `// This file was generated by cmd
package optional

import "encoding/json"

// {{ .WarpName }} represents an {{ .TypeName }} that may be optional
type {{ .WarpName }} struct {
	value *{{ .TypeName }}
}

// New{{ .WarpName }} creates an optional.{{ .WarpName }}
func New{{ .WarpName }}(val {{ .TypeName }}) {{ .WarpName }} {
	return {{ .WarpName }}{
		value: &val,
	}
}

// SetValue sets the {{ .TypeName }} value
func (opt *{{ .WarpName }}) SetValue(val {{ .TypeName }}) {
	opt.value = &val
}

func (opt {{ .WarpName }}) initValue() {{ .TypeName }} {
	var val {{ .TypeName }}
	return val
}

// Value returns the {{ .TypeName }} value or init value if not present
func (opt {{ .WarpName }}) Value() {{ .TypeName }} {
	if opt.value != nil {
		return *opt.value
	} else {
		return opt.initValue()
	}
}

// IsPresent returns whether or not the value is present
func (opt {{ .WarpName }}) IsPresent() bool {
	return opt.value != nil
}

// MarshalJSON implements the json.MarshalJSON interface
func (opt {{ .WarpName }}) MarshalJSON() ([]byte, error) {
	if opt.value != nil {
		return json.Marshal(opt.value)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON implements the json.UnmarshalJSON interface
func (opt *{{ .WarpName }}) UnmarshalJSON(data []byte) error {
	opt.value = nil

	if data == nil {
		return nil
	}

	if string(data) == "null" {
		return nil
	}

	var val {{ .TypeName }}
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}

	opt.value = &val
	return nil
}
`
