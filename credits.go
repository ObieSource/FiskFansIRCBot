package main

import (
	_ "embed"
)

//go:embed credits.txt
var Credits string

func CreditsHandler(argv []string) string {
	return Credits
}
