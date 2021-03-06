package main

import (
	"crypto/tls"
	"flag"
	"log"
	"os"
	"strings"

	"github.com/ergochat/irc-go/ircevent"
	"github.com/ergochat/irc-go/ircmsg"
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
	debug := new(bool)
	flag.BoolVar(debug, "d", false, "Print logs of server connections to stdout")
	flag.Parse()

	BotUseTls, err := UseTls(os.Getenv("IRCTLS"))
	if err != nil {
		log.Fatalln(err)
	}

	/*
		Define the bot
	*/
	irc = ircevent.Connection{
		Server:      env("IRCSERVER"),
		Nick:        env("IRCNICK"),
		Debug:       *debug,
		UseTLS:      BotUseTls,
		TLSConfig:   &tls.Config{},
		RequestCaps: []string{},
		RealName:    "https://github.com/ObieSource/FiskFansIRCBot",
	}
	saslUser := os.Getenv("IRCSUSER")
	saslPasswd := os.Getenv("IRCSPASS")
	if saslUser != "" && saslPasswd != "" {
		irc.SASLLogin = saslUser
		irc.SASLPassword = saslPasswd
		irc.UseSASL = true
	} else {
		irc.UseSASL = false
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
	err = irc.Connect()
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
