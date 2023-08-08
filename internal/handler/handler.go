package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"telegram-bot/logging"
)

func HandleCommand(handler *Handler, bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	switch msg.Command() {
	case "start":
		handler.sendMessage(bot, msg.Chat.ID, "Hello, I am bot")
	case "help":
		handler.sendMessage(bot, msg.Chat.ID, "Доступные команды:\n"+
			"/start - Начать общение с ботом\n"+
			"/help - Получить список доступных команд")
	default:
		handler.sendHa(bot, msg.Chat.ID)
	}
}

type Handler struct {
	UnknownCommand tgbotapi.FileURL
	Logger         *logging.Logger
}

func NewHandler(memURL tgbotapi.FileURL, log *logging.Logger) *Handler {
	return &Handler{
		UnknownCommand: memURL,
		Logger:         log,
	}
}

func (h *Handler) sendMessage(bot *tgbotapi.BotAPI, chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	bot.Send(msg)
}

func (h *Handler) sendHa(bot *tgbotapi.BotAPI, chatID int64) {
	h.sendMessage(bot, chatID, "ha-ha")
}
