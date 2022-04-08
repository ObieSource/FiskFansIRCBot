package main

import (
	"errors"
	"strings"
)

var UseTlsError = errors.New("Unknown value for environment variable IRCTLS. Acceptable values \"true\", \"false\"")

var UseTlsAllowableResults map[string]bool = map[string]bool{
	"":      true, // default value if not set
	"1":     true,
	"0":     false,
	"true":  true,
	"false": false,
	"t":     true,
	"f":     false,
	"yes":   true,
	"no":    false,
	"y":     true,
	"n":     false,
}

func UseTls(in string) (bool, error) {
	input := strings.ToLower(strings.TrimSpace(in))
	result, ok := UseTlsAllowableResults[input]
	if !ok {
		return false, UseTlsError
	}
	return result, nil
}
