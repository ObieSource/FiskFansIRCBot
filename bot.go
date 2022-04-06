package main

import (
	"crypto/tls"
	"github.com/ergochat/irc-go/ircevent"
	"log"
	"os"
)

var irc ircevent.Connection

func env(e string) string {
	out, ok := os.LookupEnv(e)
	if !ok {
		log.Fatalf("Environment variable %s not found, exiting.\n", e)
	}
	return out
}

func init() {
	irc = ircevent.Connection{
		Server:      env("IRCSERVER"),
		Nick:        env("IRCNICK"),
		Debug:       true,
		UseTLS:      true,
		TLSConfig:   &tls.Config{},
		RequestCaps: []string{},
	}
}
