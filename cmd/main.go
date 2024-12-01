package main

import (
	"log"

	"github.com/dimasmir03/vpnbot/config"
	"github.com/dimasmir03/vpnbot/internal"
	"gopkg.in/telebot.v3"
)

func main() {
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	apiClient := internal.NewAPIClient(cfg.APIBaseURL, cfg.APILogin, cfg.APIPassword)
	// if err != nil {
	// 	log.Fatalf("Failed to initialize API client: %v", err)
	// }

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

	internal.SetupHandlers(bot, apiClient)

	log.Println("Bot is running...")
	bot.Start()

}
