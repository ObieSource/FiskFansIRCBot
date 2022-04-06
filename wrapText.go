package main

import (
	"fmt"
	"strings"
)

const TextMaxWidth = 256

func GetWrappedText(text string) string {
	var lines []string = []string{}

	for i, c := range text {
		if len(lines) == 0 || i%TextMaxWidth == 0 {
			lines = append(lines, "")
		}
		x := len(lines) - 1
		old := lines[x]
		lines[x] = fmt.Sprintf("%s%c", old, c)
	}

	return strings.Join(lines, "\n")
}
