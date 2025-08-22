package bot_v1

import (
	"personal_knowledge_tracker/internal/controller/http/v1/bot/ui"

	tele "gopkg.in/telebot.v4"
)

func (h *Handler) Main() func(c tele.Context) error {
	return func(c tele.Context) error {
		ui.Menu.Reply(
			ui.Menu.Row(ui.BtnResources),
			ui.Menu.Row(ui.BtnReviews, ui.BtnHelp),
		)

		return c.Send("🏠 Вы вернулись в меню", ui.Menu)
	}
}
