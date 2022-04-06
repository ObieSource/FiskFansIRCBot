package main

import (
	"reflect"
	"strings"
	"testing"
)

var PodKeywordSearchResults map[string][]string = map[string][]string{
	"this one should not return any results": []string{},
	"brombaugh oberlin": []string{
		"98, John Brombaugh & Associates (Opus 11a 1973) Oberlin Conservatory of Music Oberlin Conservatory of Music Oberlin OH 44074 US",
		"1776, John Brombaugh & Associates (1976) Residence: L. Dean Nuernberger Residence: L. Dean Nuernberger Oberlin OH 44074 US",
		"1896, John Brombaugh & Associates (Opus 23b 1976) Residence: David Boe Residence: David Boe Oberlin OH 44074 US",
		"1986, John Brombaugh & Associates (Opus 5A 1969) Residence: L. Dean Nuernberger Residence: L. Dean Nuernberger Oberlin OH 44074 US",
		"2575, John Brombaugh & Associates (Opus 15 1974) First United Methodist Church First United Methodist Church Oberlin OH 44074 US",
		"8003, John Brombaugh & Associates (Opus 25 1981) Oberlin Conservatory of Music Oberlin Conservatory of Music Oberlin OH 44074 US Fairchild Chapel"},
	"flentrop oberlin warner":              []string{"728, Flentrop Orgelbouw (1974) Oberlin Conservatory of Music Oberlin Conservatory of Music Oberlin OH 44074 US Warner Hall"},
	"flentrop fl fl o ober oberlin warner": []string{"728, Flentrop Orgelbouw (1974) Oberlin Conservatory of Music Oberlin Conservatory of Music Oberlin OH 44074 US Warner Hall"}, // this is ok.
}

func TestPodKeywordSearch(t *testing.T) {
	for keyr, expect := range PodKeywordSearchResults {
		key := strings.Split(keyr, " ")
		result := PodKeywordSearch(key)
		if !reflect.DeepEqual(result, expect) && (len(result) != 0 || len(expect) != 0) {
			t.Errorf("For key %+v, did not recieve correct result.\n", key)
			t.Errorf("Expected %+v, recieved %+v.\n", expect, result)
		}
	}
}
