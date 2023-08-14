package chain

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Mood struct {
	BaseAdviser
}

func (r *Mood) Execute(chatID int64, filter *Filter, update tgbotapi.Update) {
	if update.PollAnswer == nil {
		r.Next.Execute(chatID, filter, update)
		return
	}
	opts := update.PollAnswer.OptionIDs
	fmt.Println("-----------", opts)
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
	filter.Attr = attrs
	fmt.Println("--------------  Filter Attribute  ------------", filter.Attr)
	r.isMoodSelected = true
	//if err := r.sendPoll(chatID, "Taste", models.PollQuestionsTaste); err != nil {
	//	r.SomethingWentWrong(chatID)
	//}

	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Add %s value", filter.Attr))
	r.SendMsg(msg)
	r.Next.Execute(chatID, filter, update)
}
