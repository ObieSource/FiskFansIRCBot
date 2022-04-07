# FiskFansIRCBot
Pipe Organ IRC Bot. Inspired by FiskFans[Discord]Bot

# Installation
- Install headless Chrome, this will vary depending on your operating system.
- clone this repository onto your machine and install it into bin.
```bash
git clone https://github.com/ObieSource/FiskFansIRCBot.git
cd FiskFansIRCBot
go install
```

# Run
This bot uses environment variables. Feel free to make use of the following script to easily fire up the bot with one command.
```bash
#!/usr/bin/env bash

IRCSERVER="hostname:port" \
  IRCTLS="1" \
  IRCCHANS="#chan1,#chan2" \
  IRCNICK="botnick" \
  IRCCOMM="command to run on login, such as oper login" \
  IRCSUSER="sasl username" \
  IRCSPASS="sasl password" \
  /path/to/FiskFansIRCBot

```
Note that sasl login is not required. If not set, FiskFansBot will connect without authenticating. The environmental variable IRCTLS is not required either. If not specified, defaults to true.
