package keyboards

import "gopkg.in/telebot.v3"

var (
	selector = &telebot.ReplyMarkup{}
	Btncntr  = selector.Data(" Сосал? VPN ↓", "center", "asd")
	BtnPrev  = selector.Data("⬅ нахуй туда", "prev", "asd")
	BtnNext  = selector.Data("➡ в  пизду туда", "next", "asd")
)

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

	menu.Inline(
		menu.Row(btnServerStatus, btnServerRestart),
		menu.Row(btnListUsers, btnCreateUser),
		menu.Row(btnDeleteUser, btngenerateLink),
	)

	return menu
}

func StartKeyboard() *telebot.ReplyMarkup {

	selector.Inline(
		selector.Row(BtnPrev, Btncntr, BtnNext),
	)
	return selector
}
