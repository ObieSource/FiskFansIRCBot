package main

import (
	"bytes"
	"fmt"
	"strings"
)

func UserCommandHandler(text string) string {
	buf := new(bytes.Buffer)

	fmt.Fprintln(buf, text)
	return strings.TrimSpace(buf.String())
}
