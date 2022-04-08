package main

import (
	"errors"
	"testing"
)

var UseTlsExtraCases map[string]struct {
	Result bool
	Err    error
} = map[string]struct {
	Result bool
	Err    error
}{
	"True":    {true, nil},
	"TRUE":    {true, nil},
	"tRUe":    {true, nil}, // all ok
	"nope":    {false, UseTlsError},
	" ":       {true, nil},
	" true  ": {true, nil},
	" 0":      {false, nil},
}

func TestUseTls(t *testing.T) {
	for input, expect := range UseTlsAllowableResults {
		result, err := UseTls(input)
		if err != nil {
			t.Errorf("For allowable input \"%s\", should not have returned error but returned %+v\n", input, err)
		}
		if result != expect {
			t.Errorf("For allowable input \"%s\", should have recieved %v but recieved %v.\n", input, expect, result)
		}
	}

	for input, ExpErr := range UseTlsExtraCases {
		var exp2 bool = ExpErr.Result
		var err2 error = ExpErr.Err
		result, err := UseTls(input)
		if !errors.Is(err, err2) {
			t.Errorf("For input \"%s\", should have recieved error %+v but recieved %+v\n", input, err2, err)
		}
		if result != exp2 {
			t.Errorf("For input \"%s\", should have recieved bool %v but recieved %v.\n", input, exp2, result)
		}
	}
}
