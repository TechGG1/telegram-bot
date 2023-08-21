package chain

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"telegram-bot/internal/models"
)

type Mood struct {
	BaseAdviser
}

func (r *Mood) Execute(chatID int64, filter *models.Filter, update tgbotapi.Update) {
	if update.PollAnswer.OptionIDs == nil {
		r.Next.Execute(chatID, filter, update)
		return
	}
	if r.isMoodSelected {
		r.Next.Execute(chatID, filter, update)
		return
	}
	opts := update.PollAnswer.OptionIDs
	for i := 0; i < len(opts); i++ {
		switch opts[i] {
		case 0:
			filter.Attr = append(filter.Attr, "fine")
		case 1:
			filter.Attr = append(filter.Attr, "sad")
		case 2:
			filter.Attr = append(filter.Attr, "party")
		}
	}

	r.isMoodSelected = true
	if err := r.sendPoll(chatID, "Taste", models.PollQuestionsTaste); err != nil {
		r.SomethingWentWrong(chatID)
	}

	//msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Add %s value", filter.Attr))
	//r.SendMsg(msg)
	fmt.Println(fmt.Sprintf("Add %s value", filter.Attr))
	r.Next.Execute(chatID, filter, update)
}
