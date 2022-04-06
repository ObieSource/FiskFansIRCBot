package main

import (
	"strconv"
	"strings"
)

const (
	FandomNoArgs      = "Syntax: .fandom <list/id#>"
	FandomTooManyArgs = "Too many args specified. Syntax .fandom <list/id#>"
	FandomIntError    = "The id entered must be an integer between 0 and the number of articles on Fandom."
)

func FandomHandler(argv []string) string {
	if len(argv) == 0 {
		return FandomNoArgs
	} else if len(argv) > 1 {
		return FandomTooManyArgs
	}

	if strings.ToLower(argv[0]) == "list" {
		return GetFandomArticlePrint()
	}

	idstr := argv[0]
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return FandomIntError
	}

	return FandomPage(id)

}
