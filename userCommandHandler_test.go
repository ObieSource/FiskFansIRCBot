package main

import (
	"strings"
	"testing"
)

var BasicTestCases map[string]string = map[string]string{
	".ping":       "Pong!",
	".ping hello": "Pong!",
	".yo":         UnknownCommand,
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
