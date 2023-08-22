package chain

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"telegram-bot/internal/handler"
	"telegram-bot/internal/models"
)

type Taste struct {
	BaseAdviser
	H *handler.Handler
}

func (d *Taste) Execute(chatID int64, fil *models.FilterPoll, update tgbotapi.Update) {
	filter := fil.Poll[chatID]
	opts := update.PollAnswer.OptionIDs
	attrs := make([]string, 0, 4)
	for i := 0; i < len(opts); i++ {
		switch opts[i] {
		case 0:
			attrs = append(attrs, "bitter")
		case 1:
			attrs = append(attrs, "sweet")
		case 2:
			attrs = append(attrs, "neutral")
		case 3:
			attrs = append(attrs, "spicy")
		}
	}
	filter.Attr = append(filter.Attr, attrs...)
	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Add from taste %s", filter.Attr))
	d.SendMsg(msg)
	params, err := d.H.IdentifyParams(&filter)
	if err != nil {
		d.H.Logger.Log.Error("error in converting params")
		return
	}
	d.H.FindBeerByParams(d.Bot, chatID, params)

	handler.DeleteFromPoll(fil, chatID)
}
