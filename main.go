package main

import (
	"fmt"
	"time"

	button "github.com/crescent-network/cre_tele_bot/button"
	command "github.com/crescent-network/cre_tele_bot/command"
	msg "github.com/crescent-network/cre_tele_bot/msg"

	tele "gopkg.in/telebot.v3"
)

type BotHere struct {
	tele.Bot
}

func main() {

	// Bot define
	b, err := tele.NewBot(tele.Settings{
		Token:  msg.TelegramToken_Address,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		return
	}

	command.Define_command(b)

	button.HandleButtons(b)

	fmt.Println("Starting Crescent Telegram Bot...... (To stop, press ^C.)")
	b.Start()
	fmt.Println("Finished Crescent Telegram Bot!")
}
