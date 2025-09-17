package bot_v1

import (
	"context"
	"fmt"
	botTypes "personal_knowledge_tracker/internal/controller/http/v1/bot/types"
	"personal_knowledge_tracker/internal/controller/http/v1/bot/ui"
	"personal_knowledge_tracker/internal/dto"
	"personal_knowledge_tracker/internal/types"
	"strconv"
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
			ui.Menu.Row(ui.BtnMain, ui.BtnResourcesDelete),
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
				ui.Menu.Row(ui.BtnResourcesAdd, ui.BtnResourcesDelete),
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
			ui.Menu.Row(ui.BtnResourcesAdd, ui.BtnResourcesDelete),
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
		ui.Menu.Row(ui.BtnResourcesAdd, ui.BtnResourcesDelete),
		ui.Menu.Row(ui.BtnMain),
	)

	return c.Send(
		fmt.Sprintf("🗂️ Пожалуйста, вот твои ресурсы\nСтраница %d из %d",
			state.Step, total%defaultLimit,
		),
		tele.ModeMarkdownV2, ui.Menu,
	)
}

func (h *Handler) ResourcesDelete() func(c tele.Context) error {
	return func(c tele.Context) error {
		menu, total, err := h.deleteResourceRenderPage(c.Sender().ID, 1)
		if err != nil {
			return c.Send("🛑 Произошла ошибка при получении списка ресурсов")
		}

		h.userStates[c.Sender().ID] = &dto.State{
			State: botTypes.StateResourceDelete.String(),
			Step:  1,
		}

		return c.Send(
			fmt.Sprintf("📂 Удаление ресурса\nСтраница %d из %d", 1, total),
			menu,
		)
	}
}

func (h *Handler) deleteResourceRenderPage(userID, page int64) (*tele.ReplyMarkup, int64, error) {
	menu := &tele.ReplyMarkup{}

	resources, total, err := h.usecases.Resource().List(context.Background(), userID, page, defaultLimit)
	if err != nil {
		return nil, 0, fmt.Errorf("🛑 Произошла ошибка при получении списка ресурсов")
	}

	btns := []tele.Btn{}
	for _, r := range resources {
		btn := menu.Data(r.Title, "delete_resource", r.ID)
		btns = append(btns, btn)
	}

	pageCount := (total + defaultLimit - 1) / defaultLimit

	paginationRow := []tele.Btn{}
	if page != 1 {
		paginationRow = append(paginationRow, menu.Data("⬅️ Назад", "delete_resource_page", strconv.Itoa(int(page)-1)))
	}
	if pageCount != page {
		paginationRow = append(paginationRow, menu.Data("Вперёд ➡️", "delete_resource_page", strconv.Itoa(int(page)+1)))
	}

	menu.Inline(
		menu.Row(btns...),
		menu.Row(paginationRow...),
	)

	return menu, pageCount, nil
}

func (h *Handler) deleteResourcePagination() func(c tele.Context) error {
	return func(c tele.Context) error {
		page, _ := strconv.ParseInt(c.Data(), 16, 64)
		menu, total, err := h.deleteResourceRenderPage(c.Sender().ID, page)
		if err != nil {
			return fmt.Errorf("🛑 Произошла ошибка при получении списка ресурсов")
		}

		return c.Edit(
			fmt.Sprintf("📂 Удаление ресурса\nСтраница %d из %d", page, total),
			menu,
		)
	}
}

func (h *Handler) deleteResourceByID() func(c tele.Context) error {
	return func(c tele.Context) error {
		err := h.usecases.Resource().DeleteByID(context.Background(), c.Data())
		if err != nil {
			switch err.(type) {
			case *types.NotFound:
				return c.Edit("🟠 Ресурс не найден")
			}
			return c.Edit("🛑 Произошла ошибка при удалении ресурса")
		}

		return c.Edit(fmt.Sprintf("Ресурс %s удалён ✅", c.Data()))
	}
}
