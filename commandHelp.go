package main

import (
	"bytes"
	"fmt"
	"strings"
)

func HelpHandler(argv []string) string {
	buf := new(bytes.Buffer)

	blankLine := "+" + strings.Repeat("-", 10) + "+" + strings.Repeat("-", 17) + "+" + strings.Repeat("-", 48) + "+"

	for _, c := range Commands {
		fmt.Fprintln(buf, blankLine)
		fmt.Fprintf(buf, "| %-8s | %-15s | %-46s |\n", "."+c.Command, c.Args, c.Description)
	}
	fmt.Fprintln(buf, blankLine)
	return buf.String()
}
