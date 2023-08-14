package chain

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Taste struct {
	BaseAdviser
}

func (d *Taste) Execute(chatID int64, filter *Filter, update tgbotapi.Update) {
	if d.isSpecSelected {
		d.Next.Execute(chatID, filter, update)
		return
	}
	//
	//opts := update.PollAnswer.OptionIDs
	//attrs := make([]string, 0, 3)
	//for i := 0; i < len(opts); i++ {
	//	switch opts[i] {
	//	case 0: //resolution
	//		attrs = append(attrs, "bitter")
	//	case 1: // ratio
	//		attrs = append(attrs, "sweet")
	//	case 2: // bitrate
	//		attrs = append(attrs, "neutral")
	//	}
	//}
	d.isSpecSelected = true
	msg := tgbotapi.NewMessage(chatID, "end")
	d.SendMsg(msg)
	//d.Next.Execute(chatID, filter, update)
}
