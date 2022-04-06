package main

import (
	"github.com/ergochat/irc-go/ircmsg"
)

func OnPrivMsg(e ircmsg.Message) {
	if len(e.Params) < 2 {
		return
	}
	text := e.Params[1]
	irc.Privmsg(e.Params[0], text)
}
