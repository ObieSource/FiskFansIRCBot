package main

import (
	"github.com/ergochat/irc-go/ircmsg"
	"github.com/ergochat/irc-go/ircutils"
	"strings"
)

func OnPrivMsg(e ircmsg.Message) {
	if len(e.Params) < 2 {
		return
	}
	text := e.Params[1]
	output := UserCommandHandler(text)
	for _, line := range strings.Split(output, "\n") {
		irc.Privmsg(e.Params[0], ircutils.SanitizeText(line, 384))
	}
}
