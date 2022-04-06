package main

import (
	"fmt"
	"strings"
)

var PodIncorrectArgs string = "Incorrect number of parameters.\nSyntax: .pod <o/s> <id> / .pod q <key1>..."

func PodNotOSQ(spec string) string {
	return fmt.Sprintf("Unknown specifier \"%s\". Syntax .pod <o/s/q> ...", spec)
}

func PodHandler(argv []string) string {
	if len(argv) < 2 {
		return PodIncorrectArgs
	}

	podArgs := argv[1:]
	switch strings.ToLower(argv[0]) {
	case "q":
		return strings.Join(PodKeywordSearch(podArgs), "\n")
	default:
		return PodNotOSQ(argv[0])
	}

}
