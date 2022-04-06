package main

import (
	"strings"
	"testing"
)

func TestWrapText(t *testing.T) {
	input := strings.Repeat("abcde", TextMaxWidth)

	output := GetWrappedText(input)
	outputspl := strings.Split(output, "\n")

	if len(outputspl[0]) != TextMaxWidth {
		t.Fatalf("GetWrappedText returned a line of incorrect length %d\n", len(outputspl[0]))
	}
	firstchar := input[TextMaxWidth]
	if firstchar != outputspl[1][0] {
		t.Fatalf("GetWrappedText second line first character was not correct, wanted %c got %c\n", firstchar, outputspl[1][0])
	}

}
