package bot_v1

import (
	tele "gopkg.in/telebot.v4"
)

func (h *Handler) Review() func(c tele.Context) error {
	return func(c tele.Context) error {
		return c.Send("У вас пока нет отзывов")
	}
}
