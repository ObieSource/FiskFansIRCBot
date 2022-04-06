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
	{"fandom", "<list/id#>", "Read articles from the Pipe organ Wiki", FandomHandler},
	{"pod", "<o/s/q> <id/keyword>...", "Perform a pipe organ database keyword search", PodHandler},
}
