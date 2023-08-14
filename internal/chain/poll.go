package chain

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"telegram-bot/internal/models"
)

type Poll struct {
	BaseAdviser
}

func (a *Poll) Execute(chatID int64, filter *Filter, update tgbotapi.Update) {
	if a.isMoodSelected {
		a.Next.Execute(chatID, filter, update)
		return
	}
	filter = &Filter{
		Attr: make([]string, 3),
	}
	if err := a.sendPoll(chatID, "Mood", models.PollQuestionsMood); err != nil {
		a.SomethingWentWrong(chatID)
	}
	a.isMoodSelected = true
}
