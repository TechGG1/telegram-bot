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
	if filter.IsMood {
		msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Add from mood %s", filter.Attr))
		r.SendMsg(msg)
		r.Next.Execute(chatID, filter, update)
		fmt.Println("---mood", r.BaseAdviser)
		return
	}
	opts := update.PollAnswer.OptionIDs
	attrs := make([]string, 0, 3)
	for i := 0; i < len(opts); i++ {
		switch opts[i] {
		case 0:
			attrs = append(attrs, "fine")
		case 1:
			attrs = append(attrs, "sad")
		case 2:
			attrs = append(attrs, "party")
		}
	}
	filter.Attr = append(filter.Attr, attrs...)
	filter.IsMood = true
	if err := r.sendPoll(chatID, "Choose a taste of pairing food", models.PollQuestionsTaste); err != nil {
		r.SomethingWentWrong(chatID)
	}
}
