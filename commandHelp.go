package main

import (
	"bytes"
	"fmt"
)

func HelpHandler(argv []string) string {
	buf := new(bytes.Buffer)

	for _, c := range Commands {
		fmt.Fprintf(buf, "%-8s%-15s%s\n", c.Command, c.Args, c.Description)
	}
	return buf.String()
}
