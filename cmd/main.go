package main

import (
	"log"

	"github.com/dimasmir03/vpnbot/config"
	"github.com/dimasmir03/vpnbot/internal"
	"gopkg.in/telebot.v3"
)

func main() {
	configPath := "./config/config.yaml"
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("failed load config: %v", err)
	}

	pref := telebot.Settings{
		Token: cfg.TelegramBotToken,
		Poller: &telebot.LongPoller{
			Timeout: 10,
		},
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatalf("Failed create bot: %v", err)
	}

	internal.SetupHandlers(bot, cfg)

	log.Println("Bot is running...")
	bot.Start()
}
