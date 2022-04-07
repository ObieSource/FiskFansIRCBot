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
	go func() {
		output := UserCommandHandler(text)
		// eachLine := strings.Split(output, "\n")
		var eachLine []string
		for _, ogLine := range strings.Split(output, "\n") {
			for _, wrappedLine := range GetWrappedText(ogLine) {
				eachLine = append(eachLine, wrappedLine)
			}
		}
		if len(eachLine) > PasteBinCutoff {
			eachLine = []string{UploadToPastebin(strings.Join(eachLine, "\n"))}
		}
		for _, line := range eachLine {
			if strings.TrimSpace(line) != "" {
				irc.Privmsg(e.Params[0], ircutils.SanitizeText(line, 384))
			}
		}
	}()
}
