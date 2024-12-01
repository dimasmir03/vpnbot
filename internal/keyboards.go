package internal

import "gopkg.in/telebot.v3"

func MainKeyboard() *telebot.ReplyMarkup {
	menu := &telebot.ReplyMarkup{}
	btnServerStatus := menu.Text("Статус сервера")
	btnServerRestart := menu.Text("Перезапустить сервер")
	btnListUsers := menu.Text("Список пользователей")
	btnCreateUser := menu.Text("Создать пользователя")
	btnDeleteUser := menu.Text("Удалить пользователя")
	btngenerateLink := menu.Text("Сгенерировать ссылку")

	menu.Reply(
		menu.Row(btnServerStatus, btnServerRestart),
		menu.Row(btnListUsers, btnCreateUser),
		menu.Row(btnDeleteUser, btngenerateLink),
	)

	return menu
}
