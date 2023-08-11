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
	attrs := make([]string, 0, 5)
	for i := 0; i < len(opts); i++ {
		switch opts[i] {
		case 0: //resolution
			attrs = append(attrs, "fine")
		case 1: // ratio
			attrs = append(attrs, "sad")
		case 2: // bitrate
			attrs = append(attrs, "party")
		}
	}

	filter.Attr = attrs
	fmt.Println("Filter Attribute", filter.Attr)
	r.Next.Execute(chatID, filter, update)
}
