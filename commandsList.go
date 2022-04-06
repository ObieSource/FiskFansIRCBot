package main

type Command struct {
	Command     string
	Args        string
	Description string
	Handler     func(argv []string) string
}

var Commands []Command = []Command{
	{"ping", "", "Send a ping to the bot", func(argv []string) string {
		return "Pong!"
	}},
}