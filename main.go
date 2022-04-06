package main

import (
	"github.com/ergochat/irc-go/ircmsg"
	"log"
	"strings"
)

func main() {
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
	irc.Loop()

}
