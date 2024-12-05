package telegramHandlers

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/dimasmir03/vpnbot/internal/api"
	"github.com/dimasmir03/vpnbot/internal/database"
	"github.com/dimasmir03/vpnbot/internal/keyboards"
	"github.com/dimasmir03/vpnbot/internal/responses"
	"gopkg.in/telebot.v3"
)

func SetupHandlers(bot *telebot.Bot, apiClient *api.APIClient, db *database.Store) {
	// Меню кнопок
	bot.Handle("/menu", func(c telebot.Context) error {
		return c.Send("Меню управления:")
	})

	bot.Handle(&keyboards.Btncntr, func(c telebot.Context) error {

		// TODO Добавляем id <c.Sender().ID> пользователя в бд
		log.Print("Sender: " + strconv.Itoa(int(c.Sender().ID)))
		err := db.Create(int(c.Sender().ID))
		if err != nil {
			log.Printf("Error add user id: %v", err)
			return c.Send("Что-то пошло не так, попробуйте еще раз")
		}
		// TODO Выводим кнопки /menu
		menu := keyboards.MainKeyboard()
		return c.Send("Здарово лох российский", menu)

	})
	// Главная команда /start
	bot.Handle("/start", func(c telebot.Context) error {
		// TODO вывод приведственного сообщения и кнопки "создать впн"
		return c.Send(responses.WelcomeMessage, keyboards.StartKeyboard())
	})

	bot.Handle("/devices", func(c telebot.Context) error {

		return c.Send("1 Дилдо и всё")
	})
	bot.Handle("/invite", func(c telebot.Context) error {
		// TODO вывод ссылки для приглашения друга <ссылка на бота + payload(id_client)>
		return c.Send("У тебя нету друзей. НЕЕТУУУ!!!")
	})
	bot.Handle("/help", func(c telebot.Context) error {
		// TODO вывод сообщения о боте, кнопок /menu /invite /device, кнопок поддержки </report>
		return c.Send("Дурка тебе токо поможет!!!!!")
	})
	bot.Handle("/privacy", func(c telebot.Context) error {
		// TODO вывод та хуй пойми чего
		return c.Send("Та ииииибал я эту политику \"Конфиденциальности\"")
	})

	// TODO Сделать проверку на админа и вывод команд для админов
	// Управление сервером
	bot.Handle("/server_status", func(c telebot.Context) error {
		resp, err := apiClient.GetServerStatus()
		if err != nil {
			log.Printf("Ошибка поулченися статуса сервера: %v", err)
			return c.Send("не удалось получить статус сервера.")
		}
		d, _ := json.Marshal(resp)
		return c.Send("Статус сервера: " + string(d))
	})

	bot.Handle("/panel_restart", func(c telebot.Context) error {
		_, err := apiClient.RestartPanel()
		if err != nil {
			log.Printf("Ошибка перезапуска сервера: %v", err)
			return c.Send("не удалось перезапустить сервер.")
		}
		return c.Send("Сервер успешно перезапущен.")
	})

	// Управление пользователями
	bot.Handle("/list_users", func(c telebot.Context) error {
		users, err := apiClient.GetOnlineClients()
		if err != nil {
			log.Printf("Ошибка полученися списка пользователей: %v", err)
			return c.Send("Не удалось получить список пользователей.")
		}
		return c.Send(users)
	})

	bot.Handle("/delete_user", func(c telebot.Context) error {
		args := c.Args()
		if len(args) < 1 {
			return c.Send("Использование: /delete_user <имя>")
		}

		err := apiClient.DeleteUser(args[0])
		if err != nil {
			log.Printf("Ошибка удаления пользователя: %v", err)
			return c.Send("Не удалось удалить пользователя.")
		}
		return c.Send("Пользователь успешно удален.")
	})

	bot.Handle("/generate_link", func(c telebot.Context) error {
		args := c.Args()
		if len(args) < 1 {
			return c.Send("Использование: /generate_link <имя>")
		}

		link, err := apiClient.GenerateUserLink(args[0])
		if err != nil {
			log.Printf("Ошибка генерации ссылки: %v", err)
			return c.Send("Не удалось сгенерировать ссылку.")
		}
		return c.Send("Ссылка для подключения: " + link)
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

	// bot.Handle("/id", func(c telebot.Context) error {

	// }
}
