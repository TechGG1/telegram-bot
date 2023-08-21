package chain

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"telegram-bot/internal/models"
)

type Poll struct {
	BaseAdviser
}

func (a *Poll) Execute(chatID int64, filter *models.Filter, update tgbotapi.Update) {
	if a.isSpecSelected {
		a.Next.Execute(chatID, filter, update)
		return
	}
	if err := a.sendPoll(chatID, "Mood", models.PollQuestionsMood); err != nil {
		a.SomethingWentWrong(chatID)
	}
	a.isSpecSelected = true
}
