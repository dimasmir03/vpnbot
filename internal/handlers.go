package internal

import (
	"log"
	"strconv"

	"github.com/dimasmir03/vpnbot/config"
	"gopkg.in/telebot.v3"
)

func SetupHandlers(bot *telebot.Bot, cfg *config.Config) {
	apiClient := NewAPIClient(cfg.APIBaseURL)

	bot.Handle("/start", func(c telebot.Context) error {
		return c.Send("Добро пожаловать! Используйте кнопки для управления VPN.")
	})

	bot.Handle("/create_user", func(c telebot.Context) error {
		args := c.Args()
		if len(args) < 2 {
			return c.Send("Использование: /create_user <имя> <лимит>")
		}

		name := args[0]
		limit, err := strconv.Atoi(args[1])
		if err != nil {
			return c.Send("Лимит должен быть числом.")
		}

		user := map[string]interface{}{
			"name":  name,
			"limit": limit,
		}

		res, err := apiClient.CreateUSer(user)
		if err != nil {
			log.Printf("Ошибка создания пользователя: %v", err)
			return c.Send("Не удалось создать пользователя.")
		}

		return c.Send("Пользователем успешно создан: " + res)
	})

	bot.Handle("/menu", func(c telebot.Context) error {
		menu := &telebot.ReplyMarkup{}
		btnUsers := menu.Text("Управление пользователями")
		btnServer := menu.Text("Управление сервером")

		menu.Reply(menu.Row(btnUsers, btnServer))

		return c.Send("Меню управления:", menu)
	})
}
