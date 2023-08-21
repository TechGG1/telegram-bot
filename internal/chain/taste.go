package chain

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"telegram-bot/internal/handler"
	"telegram-bot/internal/models"
)

type Taste struct {
	BaseAdviser
	H *handler.Handler
}

func (d *Taste) Execute(chatID int64, filter *models.Filter, update tgbotapi.Update) {
	if d.isTasteSelected {
		params, err := IdentifyParams(filter)
		if err != nil {
			d.H.Logger.Log.Error("error in converting params")
			return
		}
		d.H.FindBeerByParams(d.Bot, chatID, params)
		return
	}

	opts := update.PollAnswer.OptionIDs
	for i := 0; i < len(opts); i++ {
		switch opts[i] {
		case 0:
			filter.Attr = append(filter.Attr, "bitter")
		case 1:
			filter.Attr = append(filter.Attr, "sweet")
		case 2:
			filter.Attr = append(filter.Attr, "neutral")
		}
	}
	d.isTasteSelected = true
	msg := tgbotapi.NewMessage(chatID, "This is my advice!")
	d.SendMsg(msg)
	//d.Next.Execute(chatID, filter, update)
}

func IdentifyParams(filter *models.Filter) (map[string]string, error) {
	return map[string]string{
		"beer_name": "IPA",
		"abv_gt":    "6.0",
	}, nil
}
