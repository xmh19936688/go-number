package number

import (
	"bytes"
	"encoding/json"
	"testing"
)

type T struct {
	N Number `json:"n"`
}

func TestMarshalJSON(t *testing.T) {
	cases := []struct {
		input  T
		expect []byte
	}{
		{input: T{N: Number{Val: 0, Str: "null"}}, expect: []byte(`{"n":0}`)},
		{input: T{N: Number{Val: 0, Str: "abc"}}, expect: []byte(`{"n":0}`)},
		{input: T{N: Number{Val: 0, Str: ""}}, expect: []byte(`{"n":null}`)},
		{input: T{N: Number{Val: 0, Str: "0"}}, expect: []byte(`{"n":0}`)},
		{input: T{N: Number{Val: 123, Str: "123"}}, expect: []byte(`{"n":123}`)},
		{input: T{N: Number{Val: 123, Str: "000"}}, expect: []byte(`{"n":123}`)},
	}

	for _, c := range cases {
		output, _ := json.Marshal(c.input)
		if !bytes.Equal(output, c.expect) {
			t.Errorf("text failed. input [%+v], expect [%v], output [%v]", c.input, string(c.expect), string(output))
		}
	}
}

func TestUnmarshalJSON(t *testing.T) {
	cases := []struct {
		input  []byte
		expect T
	}{
		{input: []byte(`{"n":null}`), expect: T{N: Number{Val: 0, Str: ""}}},
		{input: []byte(`{"n":"null"}`), expect: T{N: Number{Val: 0, Str: "null"}}},
		{input: []byte(`{"n":"abc"}`), expect: T{N: Number{Val: 0, Str: "abc"}}},
		{input: []byte(`{"n":""}`), expect: T{N: Number{Val: 0, Str: ""}}},
		{input: []byte(`{"n":"0"}`), expect: T{N: Number{Val: 0, Str: "0"}}},
		{input: []byte(`{"n":0}`), expect: T{N: Number{Val: 0, Str: "0"}}},
		{input: []byte(`{"n":123}`), expect: T{N: Number{Val: 123, Str: "123"}}},
		{input: []byte(`{"n":"123"}`), expect: T{N: Number{Val: 123, Str: "123"}}},
	}

	for _, c := range cases {
		output := &T{}
		json.Unmarshal(c.input, &output)
		if (*output).N != c.expect.N {
			t.Errorf("text failed. input [%v], expect [%+v], output [%+v]", string(c.input), c.expect, output)
		}
	}
}
