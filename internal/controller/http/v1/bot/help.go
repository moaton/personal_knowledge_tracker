package bot_v1

import tele "gopkg.in/telebot.v4"

func (h *Handler) Help() func(c tele.Context) error {
	return func(c tele.Context) error {
		return c.Send("Доступные команды:\n" +
			"/start - начать работу\n" +
			"/help - помощь\n" +
			"/resources - список ресурсов\n" +
			"/reviews - мои отзывы")
	}
}
