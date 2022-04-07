package main

import (
	"crypto/tls"
	"github.com/ergochat/irc-go/ircevent"
	"github.com/ergochat/irc-go/ircmsg"
	"log"
	"os"
	"strings"
)

func env(e string) string {
	out, ok := os.LookupEnv(e)
	if !ok {
		log.Fatalf("Environment variable %s not found, exiting.\n", e)
	}
	return out
}

var irc ircevent.Connection

func main() {
	irc = ircevent.Connection{
		Server:      env("IRCSERVER"),
		Nick:        env("IRCNICK"),
		Debug:       false,
		UseTLS:      true,
		TLSConfig:   &tls.Config{},
		RequestCaps: []string{},
		RealName:    "FiskFansIRCBot",
	}
	irc.AddConnectCallback(func(e ircmsg.Message) {
		// attempt to set the BOT mode on ourself:
		if botMode := irc.ISupport()["BOT"]; botMode != "" {
			irc.Send("MODE", irc.CurrentNick(), "+"+botMode)
		}
		channels := env("IRCCHANS")
		for _, c := range strings.Split(channels, ",") {
			irc.Join(c)
		}
	})

	irc.AddCallback("PRIVMSG", OnPrivMsg)
	err := irc.Connect()
	if err != nil {
		log.Fatal(err)
	}
	if val, ok := os.LookupEnv("IRCCOMM"); ok {
		if err := irc.SendRaw(val); err != nil {
			log.Fatal(err)
		}
	}
	irc.Loop()

}
