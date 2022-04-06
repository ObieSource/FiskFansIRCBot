package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	PodHostname = "https://pipeorgandatabase.org"
)

var PodIncorrectArgs string = "Incorrect number of parameters.\nSyntax: .pod <o/s> <id> / .pod q <key1>..."

func PodNotOSQ(spec string) string {
	return fmt.Sprintf("Unknown specifier \"%s\". Syntax .pod <o/s/q> ...", spec)
}

func PodGetOrganUrl(id int) string {
	return fmt.Sprintf("%s/organ/%d", PodHostname, id)
}

func PodHandler(argv []string) string {
	if len(argv) < 2 {
		return PodIncorrectArgs
	}

	podArgs := argv[1:]
	switch strings.ToLower(argv[0]) {
	case "q":
		return strings.Join(PodKeywordSearch(podArgs), "\n")
	case "s":
		id, err := strconv.Atoi(podArgs[0])
		if err != nil {
			return fmt.Sprintf("Returned error %+v", err)
		}
		href, err := PodGetStoplistIndexFromOrganId(id)
		if err != nil {
			return fmt.Sprintf("Returned error %+v", err)
		}
		stoplist, err := GetStoplist(href)
		if err != nil {
			return fmt.Sprintf("Returned error %+v", err)
		}
		return stoplist
	default:
		return PodNotOSQ(argv[0])
	}

}
