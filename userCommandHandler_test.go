package main

import (
	"os"
	"strings"
	"testing"
)

var BasicTestCases map[string]string = map[string]string{
	".ping":         "Pong!",
	".ping hello":   "Pong!",
	".help":         HelpHandler([]string{}),
	".help hello":   HelpHandler([]string{}),
	".yo":           UnknownCommand,
	".pod":          PodIncorrectArgs,
	".pod s":        PodIncorrectArgs,
	".pod o":        PodIncorrectArgs,
	".pod q":        PodIncorrectArgs,
	".pod g 69":     PodNotOSQ("g"),
	".pod hello 23": PodNotOSQ("hello"),
	".fandom":       FandomNoArgs,
	".fandom hello": FandomIntError,
	".fandom 69420": FandomIntError,
	".fandom 23 45": FandomTooManyArgs,
	".stop":         OrganStopNoArgs,
	".stop jfkadsf": OrganStopNotFound,
	".credits": func() string {
		a, b := os.ReadFile("credits.txt")
		if b != nil {
			panic(b)
		}
		return string(a)
	}(),
}

func TestUserCommandHandler(t *testing.T) {
	for comm, expect := range BasicTestCases {
		t.Logf("Command: %s", comm)
		t.Logf("Expects: %s", expect)
		result := UserCommandHandler(comm)
		t.Logf("    Got: %s", result)

		if strings.TrimSpace(result) != strings.TrimSpace(expect) {
			t.Errorf("For command %s, did not get correct answer", comm)
		}
	}
}
