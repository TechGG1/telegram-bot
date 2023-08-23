package chain

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"telegram-bot/internal/chain/filter"
	"telegram-bot/internal/models"
)

type Poll struct {
	BaseAdviser
}

func (a *Poll) Execute(chatID int64, filt *models.FilterPoll, update tgbotapi.Update) {
	ok := filter.IsFilterExists(filt, chatID)
	if !ok {
		filter.AddFilterForChat(filt, chatID)
	}
	filter := filt.Poll[chatID]
	fmt.Println("000000000000000", filt.Poll[chatID])

	if filt.Poll[chatID].IsSpec {
		a.Next.Execute(chatID, filt, update)
		return
	}
	filter.Attr = make([]string, 0, 3)
	if err := a.sendPoll(chatID, "Choose your mood :)", models.PollQuestionsMood); err != nil {
		a.SomethingWentWrong(chatID)
	}
	filter.IsSpec = true
	filt.Poll[chatID] = filter
}
