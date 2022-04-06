package main

import (
	"testing"
)

var PGSIFOIDcases map[int]string = map[int]string{
	728:   "https://pipeorgandatabase.org/stoplist/15828",
	12489: "https://pipeorgandatabase.org/stoplist/12457",
}

func TestPodGetStoplistIndexFromOrganId(t *testing.T) {
	for key, expect := range PGSIFOIDcases {
		result, err := PodGetStoplistIndexFromOrganId(key)
		if err != nil {
			t.Fatalf("For key %d, shouldn't have recieved error but recieved %+v\n", key, err)
		}
		t.Logf("Expected %s, recieved %s.\n", expect, result)
		if result != expect {
			t.Errorf("For id %d, expected %s but recieved %s.\n", key, expect, result)
		}
	}
}
