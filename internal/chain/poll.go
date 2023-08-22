package chain

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"telegram-bot/internal/models"
)

type Poll struct {
	BaseAdviser
}

func (a *Poll) Execute(chatID int64, filter *models.Filter, update tgbotapi.Update) {
	if filter.IsSpec {
		a.Next.Execute(chatID, filter, update)
		fmt.Println("---poll", a.BaseAdviser)
		return
	}
	filter.Attr = make([]string, 0, 3)
	if err := a.sendPoll(chatID, "Choose your mood :)", models.PollQuestionsMood); err != nil {
		a.SomethingWentWrong(chatID)
	}
	filter.IsSpec = true
}
