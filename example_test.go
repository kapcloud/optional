package optional_test

import (
	"encoding/json"
	"fmt"

	"github.com/v8lab/optional"
)

func Example_value() {
	values := []optional.String{
		{},
		optional.NewString(""),
		optional.NewString("foo"),
		optional.NewString("bar"),
	}

	for _, v := range values {
		fmt.Println(v.Value())
	}

	// Output:
	//
	//
	// foo
	// bar
}

func Example_valueCheck() {
	values := []optional.String{
		{},
		optional.NewString(""),
		optional.NewString("foo"),
		optional.NewString("bar"),
	}

	for _, v := range values {
		if v.IsPresent() {
			fmt.Println(v.Value())
		} else {
			fmt.Println("not present")
		}
	}

	// Output:
	// not present
	//
	// foo
	// bar
}

func Example_set() {
	val := optional.NewString("baz")

	strs := []string{
		"",
		"foo",
		"bar",
	}

	for _, s := range strs {
		val.SetValue(s)
		if val.Value() == s {
			fmt.Println("ok")
		}
	}

	// Output:
	// ok
	// ok
	// ok
}

func Example_marshalJSON() {
	type example struct {
		FieldOne optional.String `json:"field_one,omitempty"`
		FieldTwo optional.String `json:"field_two"`
	}

	list := []example{
		example{},
		example{
			FieldOne: optional.NewString(""),
			FieldTwo: optional.NewString("foo"),
		},
		example{
			FieldOne: optional.NewString("foo"),
			FieldTwo: optional.NewString(""),
		},
		example{
			FieldOne: optional.NewString(""),
			FieldTwo: optional.NewString("bar"),
		},
		example{
			FieldOne: optional.NewString("bar"),
			FieldTwo: optional.NewString(""),
		},
	}

	for _, v := range list {
		out, _ := json.Marshal(&v)
		fmt.Println(string(out))
	}

	// Output:
	// {"field_one":null,"field_two":null}
	// {"field_one":"","field_two":"foo"}
	// {"field_one":"foo","field_two":""}
	// {"field_one":"","field_two":"bar"}
	// {"field_one":"bar","field_two":""}
}

func Example_unmarshalJSON() {
	var values = []string{
		"{}",
		`{"field_one":"","field_two":"foo"}`,
		`{"field_one":"foo","field_two":""}`,
		`{"field_one":null,"field_two":"foo"}`,
		`{"field_one":"foo","field_two":null}`,
		`{"field_one":null,"field_two":null}`,
	}

	type example struct {
		FieldOne optional.String `json:"field_one"`
		FieldTwo optional.String `json:"field_two"`
	}

	for _, v := range values {
		var obj = &example{}
		if err := json.Unmarshal([]byte(v), obj); err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s%s\n", obj.FieldOne.Value(), obj.FieldTwo.Value())
	}

	// Output:
	//
	// foo
	// foo
	// foo
	// foo
	//
}
