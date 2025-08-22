package telegram

import (
	"fmt"
	"time"

	tele "gopkg.in/telebot.v4"
)

type Bot struct {
	*tele.Bot
}

func NewBot(token string) (*Bot, error) {
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		return nil, fmt.Errorf("failed to create bot: %w", err)
	}

	return &Bot{b}, nil
}
