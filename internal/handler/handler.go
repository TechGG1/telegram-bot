package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"telegram-bot/internal/chain"
	"telegram-bot/logging"
)

type Handler struct {
	FileForUnknown tgbotapi.FileURL
	Logger         *logging.Logger
}

func NewHandler(memURL tgbotapi.FileURL, log *logging.Logger) *Handler {
	return &Handler{
		FileForUnknown: memURL,
		Logger:         log,
	}
}

func HandleCommand(handler *Handler, bot *tgbotapi.BotAPI, msg *tgbotapi.Message, ch chain.MessageHandler, update tgbotapi.Update) {
	switch msg.Command() {
	case "start":
		handler.Start(bot, msg.Chat.ID)
	case "help":
		handler.Help(bot, msg.Chat.ID)
	case "random":
		handler.RandomBeer(bot, msg.Chat.ID)
	case "name":
		handler.BeerName(bot, msg.Chat.ID, []byte(msg.Text))
	case "advice":
		ch.Execute(msg.Chat.ID, &chain.Filter{}, update)
	default:
		handler.UnknownReq(bot, msg.Chat.ID)
	}
}
