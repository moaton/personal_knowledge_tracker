package ui

import (
	tele "gopkg.in/telebot.v4"
)

var (
	Menu = &tele.ReplyMarkup{ResizeKeyboard: true}

	BtnMain = Menu.Text("🏠 Меню")
	BtnHelp = Menu.Text("ℹ️ Помощь")

	BtnResources     = Menu.Text("📚 Ресурсы")
	BtnResourcesAdd  = Menu.Text("💡 Добавить ресурс")
	BtnResourcesList = Menu.Text("🗂️ Мои ресурсы")

	BtnReviews = Menu.Text("⭐ Мои отзывы")

	BtnPaginationNext = Menu.Text("👉 Следующая страница")
	BtnPaginationPrev = Menu.Text("👈 Предыдущая страница")
)
