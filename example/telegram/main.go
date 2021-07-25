package main

import (
	"log"

	"go-notif/channel/telegram"
)

func main() {
	tg := telegram.NewTelegramChannel("{YOUR_TELEGRAM_TOKEN}")

	err := tg.SendTextMessage("Hello World", []string{"{RECIPIENTS_CHAT_ID}"})
	if err != nil {
		log.Print(err)
	}
}
