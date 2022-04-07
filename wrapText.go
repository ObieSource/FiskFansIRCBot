package main

import (
	"fmt"
)

const TextMaxWidth = 192

func GetWrappedText(text string) (lines []string) {

	for i, c := range text {
		if len(lines) == 0 || i%TextMaxWidth == 0 {
			lines = append(lines, "")
		}
		x := len(lines) - 1
		old := lines[x]
		lines[x] = fmt.Sprintf("%s%c", old, c)
	}

	return
}
