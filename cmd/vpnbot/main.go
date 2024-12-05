package main

import (
	"log"

	"github.com/dimasmir03/vpnbot/config"
	"github.com/dimasmir03/vpnbot/internal/api"
	"github.com/dimasmir03/vpnbot/internal/database"
	"github.com/dimasmir03/vpnbot/internal/telegramHandlers"
	"gopkg.in/telebot.v3"
)

func main() {
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	apiClient := api.NewAPIClient(cfg.APIBaseURL, cfg.APILogin, cfg.APIPassword)

	db, err := database.NewStore(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Error when try load db: %v", err)
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

	telegramHandlers.SetupHandlers(bot, apiClient, db)

	log.Println("Bot is running...")
	bot.Start()

}
