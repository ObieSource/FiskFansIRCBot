package main

import (
	"strings"
)

const (
	PREFIX         = "."
	UnknownCommand = "Unknown command. Try .help"
)

func UserCommandHandler(text string) string {

	/*
		Check for prefix
	*/
	if !strings.HasPrefix(text, PREFIX) {
		return ""
	}
	text = text[len(PREFIX):] // ignore prefix

	allargv := strings.Split(text, " ")

	command := strings.ToLower(allargv[0])

	argv := allargv[1:]

	if strings.ToLower(command) == "help" {
		return HelpHandler(argv)
	} else {
		for _, c := range Commands {
			if strings.ToLower(c.Command) == strings.ToLower(command) {
				return c.Handler(argv)
			}
		}
	}

	return UnknownCommand
}
