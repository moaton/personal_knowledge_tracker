package bot_v1

import (
	"personal_knowledge_tracker/config"
	botTypes "personal_knowledge_tracker/internal/controller/http/v1/bot/types"
	"personal_knowledge_tracker/internal/controller/http/v1/bot/ui"
	"personal_knowledge_tracker/internal/dto"
	"personal_knowledge_tracker/internal/interfaces"
	"personal_knowledge_tracker/pkg/bot/telegram"

	tele "gopkg.in/telebot.v4"

	"github.com/go-logr/logr"
)

type Handler struct {
	bot        *telegram.Bot
	usecases   interfaces.Usecases
	logger     logr.Logger
	userStates map[int64]*dto.State
}

type Dependency struct {
	Config   *config.Config
	Usecases interfaces.Usecases
	Logger   logr.Logger
}

func NewHandler(deps *Dependency) (*Handler, error) {
	bot, err := telegram.NewBot(deps.Config.Token)
	if err != nil {
		return nil, err
	}

	return &Handler{
		bot:        bot,
		usecases:   deps.Usecases,
		logger:     deps.Logger.WithName("[bot]"),
		userStates: make(map[int64]*dto.State, 10),
	}, nil
}

func (h *Handler) StartBot() {
	go h.bot.Start()
}

func (h *Handler) StopBot() {
	h.bot.Stop()
}

func (h *Handler) Register() {
	h.bot.Handle("/start", h.Start())

	h.bot.Handle("/help", h.Help())
	h.bot.Handle(&ui.BtnHelp, h.Help())
	h.bot.Handle(&ui.BtnMain, h.Main())

	h.bot.Handle("/resources", h.Resources())
	h.bot.Handle(&ui.BtnResources, h.Resources())
	h.bot.Handle(&ui.BtnResourcesAdd, h.ResourcesAdd())
	h.bot.Handle(&ui.BtnResourcesList, h.ResourcesList())

	h.bot.Handle("/reviews", h.Review())
	h.bot.Handle(&ui.BtnReviews, h.Review())

	h.bot.Handle(tele.OnText, h.messageHandler())
}

func (h *Handler) messageHandler() func(c tele.Context) error {
	return func(c tele.Context) error {
		if _, ok := h.userStates[c.Sender().ID]; !ok {
			return c.Send("😕 Упс, я вас не понял")
		}
		state := h.userStates[c.Sender().ID]
		switch state.State {
		case botTypes.StateResourceCreate.String():
			return h.resourcesAdd(c)
		case botTypes.StateResourceList.String():
			return h.resourcesList(c)
		}

		return c.Send("😕 Упс, я вас не понял")
	}
}
