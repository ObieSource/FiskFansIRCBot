package main

import (
	"testing"
)

var OrganStopTestCases map[string]string = map[string]string{
	"baxoncillo": "http://www.organstops.org/f/Fagotto.html",
	"acuta":      "http://www.organstops.org/s/Scharf.html",
	"decembass":  "http://www.organstops.org/d/Decembass.html",
}

func TestGetOrganStops(t *testing.T) {
	var stops map[string]string = map[string]string{}
	GetOrganStops(stops)

	for key, exp := range OrganStopTestCases {
		got, ok := stops[key]
		if !ok {
			t.Errorf("Stop with name %s not present in output of GetOrganStops\n", key)
		} else if got != exp {
			t.Errorf("For stop name %s, expected to get link %s, got %s.\n", key, exp, got)
		}
	}
}
