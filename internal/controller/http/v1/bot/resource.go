package bot_v1

import (
	"context"
	"fmt"
	botTypes "personal_knowledge_tracker/internal/controller/http/v1/bot/types"
	"personal_knowledge_tracker/internal/controller/http/v1/bot/ui"
	"personal_knowledge_tracker/internal/dto"
	"strings"
	"time"

	tele "gopkg.in/telebot.v4"
)

const (
	defaultLimit = 5
)

func (h *Handler) Resources() func(c tele.Context) error {
	return func(c tele.Context) error {
		ui.Menu.Reply(
			ui.Menu.Row(ui.BtnResourcesAdd, ui.BtnResourcesList),
			ui.Menu.Row(ui.BtnMain),
		)

		return c.Send("📚 Хотите посмотреть свои ресурсы или добавить новый?", ui.Menu)
	}
}

func (h *Handler) ResourcesAdd() func(c tele.Context) error {
	return func(c tele.Context) error {
		h.userStates[c.Sender().ID] = &dto.State{
			State:  botTypes.StateResourceCreate.String(),
			Step:   1,
			Buffer: make(map[string]string, 4),
		}
		return c.Send("💡 Давай создадим ресурс. Напиши название ресурса:")
	}
}

func (h *Handler) resourcesAdd(c tele.Context) error {
	state := h.userStates[c.Sender().ID]

	switch state.Step {
	case 1:
		state.Buffer["title"] = c.Text()
		state.Step = 2

		return c.Send("📝 Отлично, какой тип у ресурса будет?")
	case 2:
		state.Buffer["type"] = c.Text()
		state.Step = 3

		return c.Send("📝 Мхм, отправь мне контент ресурса")
	case 3:
		state.Buffer["content"] = c.Text()
		state.Step = 4

		return c.Send("📝 Еще чуть чуть, будем ли накидывать определнные теги?")
	case 4:
		if strings.ToLower(c.Text()) != "нет" {
			state.Buffer["tags"] = strings.ToLower(c.Text())
		}
	default:
		return c.Send("😬 Упс, что то не так")
	}
	var tags []string
	if state.Buffer["tags"] != "" {
		tags = strings.Split(state.Buffer["tags"], ",")
	}

	delete(h.userStates, c.Sender().ID)
	now := time.Now().UTC()
	err := h.usecases.Resource().Create(context.Background(), &dto.Resource{
		Title:   state.Buffer["title"],
		Type:    state.Buffer["type"],
		Content: state.Buffer["content"],
		Tags:    tags,
		Metadata: map[string]interface{}{
			"userID": c.Sender().ID,
		},
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		h.logger.Error(err, "failed to create resource")
		return c.Send("🛑 Произошла ошибка при создании ресурса")
	}

	ui.Menu.Reply(
		ui.Menu.Row(ui.BtnResourcesAdd, ui.BtnResourcesList),
	)

	return c.Send("🥳 Ресурс создан", ui.Menu)
}

func (h *Handler) ResourcesList() func(c tele.Context) error {
	return func(c tele.Context) error {
		h.userStates[c.Sender().ID] = &dto.State{
			State: botTypes.StateResourceList.String(),
			Step:  1,
		}

		resources, total, err := h.usecases.Resource().List(context.Background(), c.Sender().ID, 1, defaultLimit)
		if err != nil {
			return c.Send("🛑 Произошла ошибка при получении списка ресурсов")
		}
		if len(resources) == 0 {
			ui.Menu.Reply(
				ui.Menu.Row(ui.BtnResourcesAdd),
				ui.Menu.Row(ui.BtnMain),
			)
			return c.Send("🗂️ У вас пока нет ресурсов", ui.Menu)
		}

		for _, r := range resources {
			c.Send(fmt.Sprintf("ℹ️ ID: `%s`\n📚 Название: %s\n✨ Тип: %s\n🏷️ Теги: %s\n📝 Контент:\n%s\n",
				r.ID, r.Title, r.Type, strings.Join(r.Tags, ", "), r.Content,
			), tele.ModeMarkdownV2)
		}

		var btns []tele.Btn = []tele.Btn{}
		if total > defaultLimit {
			btns = append(btns, ui.BtnPaginationNext)
		}

		ui.Menu.Reply(
			ui.Menu.Row(btns...),
			ui.Menu.Row(ui.BtnResourcesAdd),
			ui.Menu.Row(ui.BtnMain),
		)

		return c.Send(
			fmt.Sprintf("🗂️ Пожалуйста, вот твои ресурсы\nСтраница %d из %d",
				1, total%defaultLimit,
			),
			tele.ModeMarkdownV2, ui.Menu,
		)
	}
}
func (h *Handler) resourcesList(c tele.Context) error {
	state := h.userStates[c.Sender().ID]

	switch c.Text() {
	case "👉 Следующая страница":
		state.Step++
	case "👈 Предыдущая страница":
		state.Step--
		if state.Step <= 0 {
			state.Step = 1
		}
	default:
		ui.Menu.Reply(
			ui.Menu.Row(ui.BtnResources),
			ui.Menu.Row(ui.BtnReviews, ui.BtnHelp),
		)
		return c.Send("😬 Упс, вот вам кнопки", ui.Menu)
	}
	resources, total, err := h.usecases.Resource().List(context.Background(), c.Sender().ID, int64(state.Step), defaultLimit)
	if err != nil {
		return c.Send("🛑 Произошла ошибка при получении списка ресурсов")
	}
	if len(resources) == 0 {
		ui.Menu.Reply(
			ui.Menu.Row(ui.BtnPaginationPrev),
			ui.Menu.Row(ui.BtnMain),
		)
		return c.Send("🗂️ На этом ресурсы закончились", ui.Menu)
	}

	for _, r := range resources {
		c.Send(fmt.Sprintf("ℹ️ ID: `%s`\n📚 Название: %s\n✨ Тип: %s\n🏷️ Теги: %s\n📝 Контент:\n%s\n",
			r.ID, r.Title, r.Type, strings.Join(r.Tags, ", "), r.Content,
		), tele.ModeMarkdownV2)
	}

	var paginationBtn []tele.Btn = []tele.Btn{ui.BtnPaginationPrev, ui.BtnPaginationNext}
	if state.Step == 1 {
		paginationBtn = []tele.Btn{ui.BtnPaginationNext}
	}
	ui.Menu.Reply(
		ui.Menu.Row(paginationBtn...),
		ui.Menu.Row(ui.BtnResourcesAdd),
		ui.Menu.Row(ui.BtnMain),
	)

	return c.Send(
		fmt.Sprintf("🗂️ Пожалуйста, вот твои ресурсы\nСтраница %d из %d",
			state.Step, total%defaultLimit,
		),
		tele.ModeMarkdownV2, ui.Menu,
	)
}
